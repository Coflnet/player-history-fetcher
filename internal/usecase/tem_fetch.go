package usecase

import (
	"github.com/Coflnet/player-name-fetcher/internal/db"
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

	batch := make([]db.CoflPlayer, 0)
	for player := range ch {
		batch = append(batch, db.CoflPlayer{
			MinecraftUuid: player.ID.PlayerUUID,
		})

		if err != nil {
			log.Error().Err(err).Msgf("can not process player, %v", player)
		}

		if len(batch) < 100 {
			continue
		}

		QueuePlayers(batch)
		batch = make([]db.CoflPlayer, 0)

		slowDown, _ := strconv.Atoi(os.Getenv("SLOW_DOWN_MS"))
		time.Sleep(time.Millisecond * time.Duration(slowDown))
	}
}
