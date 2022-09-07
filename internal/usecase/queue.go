package usecase

import (
	"github.com/Coflnet/player-name-fetcher/internal/kafka"
	"github.com/Coflnet/player-name-fetcher/internal/mongo"
	"github.com/rs/zerolog/log"
	"time"
)

var startTime = time.Now()

func QueuePlayer(uuid string) error {
	existingPlayer, err := mongo.PlayerByUUID(uuid)
	if err != nil {
		return err
	}

	if queueCanBeSkipped(existingPlayer) {
		log.Warn().Msgf("skipping player %s", uuid)
		return nil
	}

	log.Info().Msgf("queueing player %s", uuid)
	err = mongo.InsertEmptyPlayer(&mongo.Player{
		UUID:      uuid,
		LastQueue: time.Now(),
	})

	if err != nil {
		return err
	}

	return kafka.WritePlayerPayload(kafka.PlayerKafkaPayload{
		UUID: uuid,
	})
}

func queueCanBeSkipped(player *mongo.Player) bool {
	if player == nil {
		return false
	}

	log.Info().Msgf("check if %s is after %s: %t", player.LastQueue, startTime, player.LastQueue.After(startTime))
	if player.LastQueue.After(startTime) {
		log.Warn().Msgf("player %s was already queued at %s", player.UUID, player.LastQueue)
		return true
	}

	return false
}
