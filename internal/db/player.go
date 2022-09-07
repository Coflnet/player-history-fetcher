package db

type CoflPlayer struct {
	MinecraftUuid string
}

func PlayersFromDb(startId, endId int) (<-chan CoflPlayer, error) {
	rows, err := db.Query("SELECT AccountUuid FROM McIds WHERE Id >= ? AND Id <= ?", startId, endId)
	if err != nil {
		return nil, err
	}

	channel := make(chan CoflPlayer, 100)
	defer close(channel)
	for rows.Next() {
		var uuid []byte
		err := rows.Scan(&uuid)
		if err != nil {
			return nil, err
		}

		if len(uuid) == 0 {
			continue
		}

		player := CoflPlayer{
			MinecraftUuid: string(uuid),
		}
		channel <- player
	}
	return channel, nil
}
