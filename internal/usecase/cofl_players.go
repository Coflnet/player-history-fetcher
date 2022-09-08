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

	for {

		players, err := db.PlayersFromDb(start, start+2000)
		if err != nil {
			log.Error().Err(err).Msgf("can not get players from db")
		}

		batch := make([]db.CoflPlayer, 0)
		for player := range players {
			batch = append(batch, player)

			if len(batch) < 1000 {
				continue
			}

			playersQueued := QueuePlayers(batch)

			log.Info().Msgf("queued %d players", playersQueued)
			log.Info().Msgf("%d players skipped", len(batch)-playersQueued)

			batch = make([]db.CoflPlayer, 0)
			slowDown, _ := strconv.Atoi(os.Getenv("SLOW_DOWN_MS"))
			time.Sleep(time.Millisecond * time.Duration(slowDown))
		}

		start += 1000

		if start > 440_252_059 {
			return
		}

	}
}
