package main

import (
	"github.com/Coflnet/player-name-fetcher/internal/kafka"
	"github.com/Coflnet/player-name-fetcher/internal/mongo"
	"github.com/Coflnet/player-name-fetcher/internal/usecase"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Info().Msgf("starting player name fetcher")
	err := kafka.InitReader()
	if err != nil {
		log.Panic().Err(err).Msgf("can not connect to kafka")
	}

	err = mongo.Init()
	if err != nil {
		log.Panic().Err(err).Msgf("can not connect to mongo")
	}

	log.Info().Msgf("start ingester")
	err = usecase.StartIngester()
	if err != nil {
		log.Panic().Err(err).Msgf("can not start ingester")
	}
}
