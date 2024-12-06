package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ModeInfo struct {
	Mode  string `bson:"mode"`
	Usage int    `bson:"usage"`
}

func GetPopularModes(client *mongo.Client, areaCode string) ([]ModeInfo, error) {
	collection := client.Database("game_data").Collection("mode_usage")

	var results []ModeInfo
	filter := bson.M{"area_code": areaCode}
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var mode ModeInfo
		if err := cursor.Decode(&mode); err != nil {
			return nil, err
		}
		results = append(results, mode)
	}

	return results, nil
}

func NewMongoClient(uri string) (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return client, err
}
