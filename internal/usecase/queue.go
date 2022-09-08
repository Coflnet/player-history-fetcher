package usecase

import (
	"github.com/Coflnet/player-name-fetcher/internal/db"
	"github.com/Coflnet/player-name-fetcher/internal/kafka"
	"github.com/Coflnet/player-name-fetcher/internal/mongo"
	"github.com/rs/zerolog/log"
	"time"
)

var startTime = time.Now()

func QueuePlayers(players []db.CoflPlayer) int {
	playersToQueue := make(chan db.CoflPlayer, len(players))
	sum := 0

	for _, player := range players {
		go func(p db.CoflPlayer) {
			defer close(playersToQueue)
			if queueCanBeSkipped(p.MinecraftUuid) {
				return
			}
			playersToQueue <- p
		}(player)
	}

	payloads := make([]kafka.PlayerKafkaPayload, 0)
	for player := range playersToQueue {
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

		sum++
	}

	err := kafka.WritePlayerPayloads(payloads)
	if err != nil {
		log.Error().Err(err).Msgf("can not write player payloads")
		return sum
	}
	return sum
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
		return true
	}

	return false
}
