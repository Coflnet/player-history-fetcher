package usecase

import (
	"github.com/Coflnet/player-name-fetcher/internal/db"
	"github.com/rs/zerolog/log"
	"os"
	"strconv"
	"time"
)

var start = 367_512_453

func StartCoflFetch() {

	batchSize := 10_000
	for {

		players, err := db.PlayersFromDb(start, start+batchSize)
		if err != nil {
			log.Error().Err(err).Msgf("can not get players from db")
		}

		batch := make([]db.CoflPlayer, 0)
		for player := range players {
			batch = append(batch, player)
		}
		playersQueued := QueuePlayers(batch)

		log.Info().Msgf("queued %d players, %d players skipped", playersQueued, len(batch)-playersQueued)
		slowDown, _ := strconv.Atoi(os.Getenv("SLOW_DOWN_MS"))
		time.Sleep(time.Millisecond * time.Duration(slowDown))

		if start > 440_252_059 {
			return
		}
		start += batchSize
	}
}
