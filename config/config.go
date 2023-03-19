package config

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Config struct {
	MongoUri     string
	DatabaseName string
}

func Load() *Config {
	return &Config{
		MongoUri:     os.Getenv("MONGODB_URI"),
		DatabaseName: os.Getenv("DB_NAME"),
	}
}

func ConnectDB(cfg *Config) (*mongo.Database, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}

	return client.Database(cfg.DatabaseName), nil
}

func CreateIndexes(db *mongo.Database, db_name string, unique_key string) (*mongo.Database, error) {
	_, err := db.Collection(db_name).Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: unique_key, Value: 1}},
			Options: options.Index().SetUnique(true),
		},
	)
	return db, err
}
