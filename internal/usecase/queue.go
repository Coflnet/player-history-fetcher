package usecase

import (
	"github.com/Coflnet/player-name-fetcher/internal/db"
	"github.com/Coflnet/player-name-fetcher/internal/kafka"
	"github.com/Coflnet/player-name-fetcher/internal/mongo"
	"github.com/rs/zerolog/log"
	"time"
)

var startTime = time.Now()

func QueuePlayers(players []db.CoflPlayer) {
	playersToQueue := make([]db.CoflPlayer, 0)

	for _, player := range players {
		if queueCanBeSkipped(player.MinecraftUuid) {
			continue
		}

		playersToQueue = append(playersToQueue, player)
	}

	payloads := make([]kafka.PlayerKafkaPayload, 0)
	for _, player := range playersToQueue {
		log.Info().Msgf("queueing player %s", player.MinecraftUuid)
		err := mongo.InsertEmptyPlayer(&mongo.Player{
			UUID:      player.MinecraftUuid,
			LastQueue: time.Now(),
		})
		if err != nil {
			log.Error().Err(err).Msgf("can not insert player %s", player.MinecraftUuid)
			continue
		}

		payloads = append(payloads, kafka.PlayerKafkaPayload{
			UUID: player.MinecraftUuid,
		})
	}

	kafka.WritePlayerPayloads(payloads)
}

func queueCanBeSkipped(uuid string) bool {

	existingPlayer, err := mongo.PlayerByUUID(uuid)
	if err != nil {
		return false
	}

	if existingPlayer == nil {
		return false
	}

	oneYearAgo := time.Now().AddDate(-1, 0, 0)

	if existingPlayer.LastQueue.After(oneYearAgo) {
		log.Warn().Msgf("player %s was already queued at %s", existingPlayer.UUID, existingPlayer.LastQueue)
		return true
	}

	return false
}
