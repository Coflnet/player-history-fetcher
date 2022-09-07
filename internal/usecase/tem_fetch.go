package usecase

import (
	"github.com/Coflnet/player-name-fetcher/internal/mongo"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
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

		slowDown, _ := strconv.Atoi(os.Getenv("SLOW_DOWN_MS"))
		time.Sleep(time.Millisecond * time.Duration(slowDown))
	}
}

func processPlayer(player *mongo.TemPlayer) error {
	return QueuePlayer(player.ID.PlayerUUID)
}
