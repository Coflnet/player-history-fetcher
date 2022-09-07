package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"time"
)

func WritePlayerPayloads(payload []PlayerKafkaPayload) error {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	messages := []kafka.Message{}
	for _, p := range payload {
		serialized, err := json.Marshal(p)
		if err != nil {
			return err
		}
		messages = append(messages, kafka.Message{
			Key:   []byte(p.UUID),
			Value: serialized,
		})
	}

	err := playerWriter.WriteMessages(ctx, messages...)
	return err
}
