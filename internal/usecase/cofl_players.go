package usecase

import (
	"github.com/Coflnet/player-name-fetcher/internal/db"
	"github.com/rs/zerolog/log"
)

var start = 0

func StartCoflFetch() {

	for {

		players, err := db.PlayersFromDb(start, start+1000)
		if err != nil {
			log.Error().Err(err).Msgf("can not get players from db")
		}

		for player := range players {
			err := QueuePlayer(player.MinecraftUuid)
			if err != nil {
				log.Error().Err(err).Msgf("can not queue player: %v", player)
			}
		}

		start += 1000

		if start > 50_000 {
			return
		}
	}
}
