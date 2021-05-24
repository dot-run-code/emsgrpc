package services

import (
	"bytes"
	"context"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	"github.com/dot-run-code/emsgrpc/proto"
)

type TopicService struct {
	publisherService Publisher
	uploadFolder     string
}

func NewTopicService(p Publisher, folder string) *TopicService {
	return &TopicService{publisherService: p, uploadFolder: folder}
}

func (t *TopicService) Publish(ctx context.Context, req *proto.PublishRequest) (*proto.PublishResponse, error) {
	fmt.Printf("Incoming request content:%v\n", req.Content)
	return &proto.PublishResponse{Message: "successfully published"}, nil
}
func (t *TopicService) PublishStream(stream proto.Topic_PublishStreamServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		log.Println("Invalid request content")
		return errors.New("send proper metadata")
	}
	topic := md.Get("topic")[0]
	log.Printf("Topic:%v\n", topic)
	headers := make(map[string][]byte)
	for k, v := range md {
		if k != "authorization" && k != "topic" {
			buf := &bytes.Buffer{}
			gob.NewEncoder(buf).Encode(v)
			headers[k] = buf.Bytes()
		}
	}
	headers["x-header"] = []byte("grpc stream")
	ch := make(chan []byte)
	wg := &sync.WaitGroup{}
	go func(ch chan<- []byte, topic string, headers map[string][]byte) {
		defer close(ch)
		for {
			err := contextError(stream.Context())
			if err != nil {
				panic(err)
			}
			data, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				panic(err)
			}
			ch <- data.Content
		}
	}(ch, topic, headers)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go publish(t.publisherService, ch, topic, headers, wg)
	}
	wg.Wait()
	stream.SendAndClose(&proto.AcknowledgementResponse{Message: "successful"})
	return nil
}
func (t *TopicService) StreamFile(stream proto.Topic_StreamFileServer) error {
	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		log.Println("Invalid request content")
		return errors.New("send proper metadata")
	}
	topic := md.Get("topic")[0]
	log.Printf("Topic:%v\n", topic)
	headers := make(map[string][]byte)
	for k, v := range md {
		if k != "authorization" && k != "topic" {
			buf := &bytes.Buffer{}
			gob.NewEncoder(buf).Encode(v)
			headers[k] = buf.Bytes()
		}
	}
	dirLocation := filepath.Join(t.uploadFolder, topic)
	if _, err := os.Stat(dirLocation); os.IsNotExist(err) {
		os.Mkdir(dirLocation, os.ModePerm)
	}
	fileLocation := filepath.Join(t.uploadFolder, topic, strconv.Itoa(int(time.Now().Unix())))

	file, err := os.Create(fileLocation)
	defer file.Close()
	if err != nil {
		log.Println(err)
		stream.SendAndClose(&proto.AcknowledgementResponse{Message: err.Error(), Code: proto.PublishStatusCode_Failed})
		return err
	}
	for {
		err := contextError(stream.Context())
		if err != nil {
			log.Println(err)
			stream.SendAndClose(&proto.AcknowledgementResponse{Message: err.Error(), Code: proto.PublishStatusCode_Failed})
			return err
		}
		chumkStream, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
			stream.SendAndClose(&proto.AcknowledgementResponse{Message: err.Error(), Code: proto.PublishStatusCode_Failed})
			return err
		}
		_, err = file.Write(chumkStream.Content)

	}
	stream.SendAndClose(&proto.AcknowledgementResponse{Message: "successful"})

	return nil
}

func contextError(ctx context.Context) error {
	switch ctx.Err() {
	case context.Canceled:
		log.Println("request is canceled")
		return status.Error(codes.Canceled, "request is canceled")
	case context.DeadlineExceeded:
		log.Println("deadline is exceeded")
		return status.Error(codes.DeadlineExceeded, "deadline is exceeded")
	default:
		return nil
	}
}

func publish(publisherService Publisher, ch <-chan []byte, topic string, headers map[string][]byte, wg *sync.WaitGroup) {
	for content := range ch {
		publisherService.PublishToKafka(topic, headers, content)
	}
	wg.Done()
}
