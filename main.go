package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/dot-run-code/emsgrpc/proto"

	"github.com/dot-run-code/emsgrpc/services"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var setting = LoadSettings()
	authService := services.NewAuthoizationService(setting.ClientId,
		setting.ClientSecret,
		setting.ValidAudiences,
		setting.ValidIssuer,
		nil)
	hostname := os.Getenv("SVC_HOST_NAME")
	if len(hostname) <= 0 {
		hostname = "0.0.0.0"
	}

	port := setting.Port

	if len(port) <= 0 {
		port = "5003"
	}
	//publisherService := services.NewKafkaPublisherService(setting.KafkaBrokers)
	var publisherService services.Publisher = &services.MockPublisher{}
	topicService := services.NewTopicService(publisherService, setting.UploadPath)

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	lis, err := net.Listen("tcp", hostname+":"+port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	//opts := []grpc.ServerOption{}
	//grpcServer := grpc.NewServer(opts...)
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(authService.Unary()),
		grpc.StreamInterceptor(authService.Stream()),
	)

	proto.RegisterTopicServer(grpcServer, topicService)

	// reflection service on gRPC server.
	reflection.Register(grpcServer)

	go func() {
		fmt.Println("Server running on ", (hostname + ":" + port))
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	// Wait for Control C to exit
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	// Block until a signal is received
	<-ch
	fmt.Println("Stopping the server")
	grpcServer.Stop()
	fmt.Println("Closing the listener")
	lis.Close()
	fmt.Println("Server Shutdown")

}
