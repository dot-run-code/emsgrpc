package services

type Publisher interface {
	PublishToKafka(string, map[string][]byte, []byte)
}
