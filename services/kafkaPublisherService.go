package services

import (
	"log"
	"strings"

	"github.com/Shopify/sarama"
)

type KafkaPublisherService struct {
	syncProducer sarama.SyncProducer
}

func NewKafkaPublisherService(brokerUrl string) *KafkaPublisherService {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true

	brokerList := strings.Split(brokerUrl, ",")
	if len(brokerList) == 0 {
		log.Fatalf("Kafka broker is not valid")
		panic("Kafka broker is not valid")
	}
	//client, err := sarama.NewClient(brokerList, config)
	// if err != nil {
	// 	log.Println("Can not create kafka client:%v\n", err)
	// }
	producer, err := sarama.NewSyncProducer(brokerList, config)

	if err != nil {
		log.Println("Can not create kafka broker:%v\n", err)
		//temporary for debugging
		//panic("Can not create kafka producer")
		return nil
	}
	return &KafkaPublisherService{syncProducer: producer}
}

func (p *KafkaPublisherService) PublishToKafka(topic string, headers map[string][]byte, message []byte) {
	var recordHeaders = make([]sarama.RecordHeader, 0)
	for k, v := range headers {
		recordHeader := sarama.RecordHeader{Key: []byte(k), Value: v}
		recordHeaders = append(recordHeaders, recordHeader)
	}
	msg := &sarama.ProducerMessage{
		Topic:   topic,
		Value:   sarama.StringEncoder(message),
		Headers: recordHeaders,
	}

	p.syncProducer.SendMessage(msg)

}
