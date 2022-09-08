package main

import (
	"github.com/Coflnet/player-name-fetcher/internal/db"
	"github.com/Coflnet/player-name-fetcher/internal/kafka"
	"github.com/Coflnet/player-name-fetcher/internal/mongo"
	"github.com/Coflnet/player-name-fetcher/internal/usecase"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msgf("starting coordinator..")

	err := mongo.Init()

	if err != nil {
		log.Panic().Err(err).Msgf("can not connect to mongo")
	}

	err = kafka.InitWriter()
	if err != nil {
		log.Panic().Err(err).Msgf("can not connect to kafka")
	}

	err = db.Init()
	if err != nil {
		log.Panic().Err(err).Msgf("can not connect to db")
	}

	usecase.StartCoflFetch()
	usecase.StartTemFetch()
}
