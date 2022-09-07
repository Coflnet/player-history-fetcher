package kafka

import (
	"github.com/segmentio/kafka-go"
	"os"
)

var (
	playerWriter *kafka.Writer
	playerReader *kafka.Reader
)

type PlayerKafkaPayload struct {
	UUID string `json:"uuid"`
}

func InitWriter() error {
	playerWriter = &kafka.Writer{
		Addr:                   kafka.TCP(os.Getenv("KAFKA_HOST")),
		Topic:                  "player-uuid-nicknames",
		AllowAutoTopicCreation: true,
	}

	return nil
}

func InitReader() error {
	playerReader = kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{os.Getenv("KAFKA_HOST")},
		Topic:   "player-uuid-nicknames",
		GroupID: "player-name-fetcher-cg",
	})

	return nil
}
