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
	// log.Printf("[MONGO] MongoDB Configuration: %+v", cfg)

	clientOptions := options.Client().
		ApplyURI(fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.User, cfg.Password, cfg.Host, cfg.Port))

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Printf("[MONGO] Error connecting to MongoDB: %v", err)
		return nil, err
	}

	// Ping the MongoDB server to ensure connectivity
	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		log.Printf("[MONGO] Error pinging MongoDB server: %v", err)
		return nil, err
	}

	db := client.Database(cfg.Database)

	log.Printf("[MONGO] Connected to MongoDB: %s", cfg.Database)
	return &MongoDB{client: client, db: db}, nil
}

func (db *MongoDB) Close(ctx context.Context) error {
	err := db.client.Disconnect(ctx)
	if err != nil {
		log.Printf("[MONGO] Error disconnecting from MongoDB: %v", err)
	}
	return err
}
