package db

import "github.com/rs/zerolog/log"

type CoflPlayer struct {
	MinecraftUuid string
}

func PlayersFromDb(startId, endId int) (<-chan CoflPlayer, error) {
	log.Info().Msgf("getting players from db, from %v to %v", startId, endId)
	rows, err := db.Query("SELECT value FROM UuId WHERE Id >= ? AND Id <= ?", startId, endId)
	if err != nil {
		return nil, err
	}

	channel := make(chan CoflPlayer, 100)
	go func() {
		defer close(channel)
		for rows.Next() {
			var uuid []byte
			err := rows.Scan(&uuid)
			if err != nil {
				continue
			}

			if len(uuid) == 0 {
				continue
			}

			player := CoflPlayer{
				MinecraftUuid: string(uuid),
			}
			channel <- player
		}
	}()
	return channel, nil
}
