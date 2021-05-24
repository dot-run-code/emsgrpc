package services

import (
	"log"
)

type MockPublisher struct {
}

func (p *MockPublisher) PublishToKafka(topic string, headers map[string][]byte, message []byte) {
	log.Println("Publish to kafka successfully")
}
