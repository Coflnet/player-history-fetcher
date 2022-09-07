package mongo

import (
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/bson"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	temPlayerCollection *mongo.Collection
	playerCollection    *mongo.Collection
)

func Init() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_URL")))

	temPlayerCollection = client.Database("inventories").Collection("players")
	playerCollection = client.Database("players").Collection("players")

	return err
}

func TemPlayersChannel() (<-chan TemPlayer, error) {
	ctx := context.Background()

	cur, err := temPlayerCollection.Find(ctx, bson.D{})

	if err != nil {
		return nil, err
	}

	channel := make(chan TemPlayer, 100)

	go func() {
		defer close(channel)

		for cur.Next(ctx) {
			var player TemPlayer
			err := cur.Decode(&player)

			if err != nil {
				log.Error().Err(err).Msgf("can not decode player, %v", player)
			}

			channel <- player
		}
	}()

	return channel, nil
}
