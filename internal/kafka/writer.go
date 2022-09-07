package kafka

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"time"
)

func WritePlayerPayload(payload PlayerKafkaPayload) error {
	serialized, err := json.Marshal(payload)

	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = playerWriter.WriteMessages(ctx, kafka.Message{
		Key:   []byte(payload.UUID),
		Value: serialized,
	})

	return err
}
