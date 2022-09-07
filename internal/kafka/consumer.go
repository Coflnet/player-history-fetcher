package kafka

import (
	"context"
	"github.com/segmentio/kafka-go"
	"time"
)

func ReadPlayerPayload() (kafka.Message, error) {
	ctx := context.Background()

	return playerReader.FetchMessage(ctx)
}

func CommitPlayerPayload(msg kafka.Message) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return playerReader.CommitMessages(ctx, msg)
}
