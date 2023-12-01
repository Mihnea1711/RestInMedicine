package mongo

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/database"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/config"
)

type MongoDB struct {
	client *mongo.Client
	db     *mongo.Database
}

func NewMongoDB(ctx context.Context, cfg *config.MongoDBConfig) (database.Database, error) {
	// auth := options.Client().SetAuth()
	clientOptions := options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d/?authSource=admin", cfg.User, cfg.Password, cfg.Host, cfg.Port))

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Printf("[CONSULTATION] Error connecting to MongoDB: %v", err)
		return nil, err
	}

	log.Printf("Connected client to Mongo: %s", fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.User, cfg.Password, cfg.Host, cfg.Port))

	// Ping the MongoDB server to ensure connectivity
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Printf("[CONSULTATION] Error pinging MongoDB server: %v", err)
		return nil, err
	}

	db := client.Database(cfg.Database)

	log.Printf("[CONSULTATION] Connected to MongoDB: %s", cfg.Database)
	return &MongoDB{client: client, db: db}, nil
}

func (db *MongoDB) Close(ctx context.Context) error {
	err := db.client.Disconnect(ctx)
	if err != nil {
		log.Printf("[CONSULTATION] Error disconnecting from MongoDB: %v", err)
	}
	return err
}
