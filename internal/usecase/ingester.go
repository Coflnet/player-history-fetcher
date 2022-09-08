package usecase

import (
	"encoding/json"
	"github.com/Coflnet/player-name-fetcher/internal/kafka"
	"github.com/rs/zerolog/log"
	kafkago "github.com/segmentio/kafka-go"
	"time"
)

func StartIngester() error {
	for {
		msg, err := kafka.ReadPlayerPayload()
		if err != nil {
			log.Error().Err(err).Msgf("can not read player payload")
			continue
		}

		err = processPayloadMessage(msg)
		if err != nil {
			log.Error().Err(err).Msgf("can not process player payload")
			continue
		}

		err = kafka.CommitPlayerPayload(msg)
		if err != nil {
			log.Error().Err(err).Msgf("can not commit message")
		}

		time.Sleep(1500 * time.Millisecond)
	}
}

func processPayloadMessage(msg kafkago.Message) error {
	var player kafka.PlayerKafkaPayload
	err := json.Unmarshal(msg.Value, &player)
	if err != nil {
		return err
	}

	return FetchUUID(player.UUID)
}
