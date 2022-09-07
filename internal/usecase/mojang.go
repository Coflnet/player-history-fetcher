package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/Coflnet/player-name-fetcher/internal/mongo"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

func FetchUUID(uuid string) error {
	url := fmt.Sprintf("https://api.mojang.com/user/profiles/%s/names", uuid)
	response, err := http.DefaultClient.Get(url)

	if err != nil {
		log.Err(err).Msgf("can not fetch uuid")
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error().Err(err).Msgf("can not close body")
		}
	}(response.Body)

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		log.Err(err).Msgf("can not read response body")
		return err
	}

	var player []mongo.MojangPlayer
	err = json.Unmarshal(bytes, &player)
	if err != nil {
		log.Err(err).Msgf("can not unmarshal json")
		return err
	}

	err = mongo.SetMojangPlayer(uuid, player)
	if err != nil {
		log.Err(err).Msgf("can not set mojang player")
		return err
	}

	log.Info().Msgf("player %s: %v", uuid, player)
	return nil
}
