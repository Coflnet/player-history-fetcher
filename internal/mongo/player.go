package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MojangPlayer struct {
	Name        string `json:"name"`
	ChangedToAt int64  `json:"changedToAt,omitempty"`
}

type Player struct {
	MojangPlayer *[]MojangPlayer `bson:"mojangPlayer"`
	UUID         string          `bson:"uuid"`
	LastQueue    time.Time       `bson:"lastQueue"`
	LastWrite    time.Time       `bson:"lastWrite"`
}

func PlayerByUUID(uuid string) (*Player, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"uuid", uuid}}

	response := playerCollection.FindOne(ctx, filter)

	if response.Err() != nil {
		if response.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, response.Err()
	}

	var player Player
	err := response.Decode(&player)
	if err != nil {
		return nil, err
	}

	return &player, nil
}

func InsertEmptyPlayer(player *Player) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := playerCollection.InsertOne(ctx, player)

	return err
}

func SetMojangPlayer(uuid string, player []MojangPlayer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.D{{"uuid", uuid}}
	update := bson.D{{"$set", bson.D{{"mojangPlayer", player}}}}

	_, err := playerCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	filter = bson.D{{"uuid", uuid}}
	update = bson.D{{"$set", bson.D{{"lastWrite", time.Now()}}}}
	_, err = playerCollection.UpdateOne(ctx, filter, update)

	return err
}
