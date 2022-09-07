package usecase

import (
	"github.com/Coflnet/player-name-fetcher/internal/mongo"
	"github.com/rs/zerolog/log"
	"time"
)

func StartTemFetch() {
	ch, err := mongo.TemPlayersChannel()

	if err != nil {
		log.Error().Err(err).Msgf("can not get tem players")
		return
	}

	for player := range ch {
		err = processPlayer(&player)

		if err != nil {
			log.Error().Err(err).Msgf("can not process player, %v", player)
		}

		time.Sleep(time.Millisecond * 10)

	}
}

func processPlayer(player *mongo.TemPlayer) error {
	return QueuePlayer(player.ID.PlayerUUID)
}
