package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/config"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

// NewRedisClient creates a new Redis client based on the provided configuration.
func NewRedisClient(ctx context.Context, config *config.RedisConfig) (*RedisClient, error) {
	// Create a new Redis client using the configuration options
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	// Ping the Redis server to check the connection
	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		// Log an error if the connection attempt fails
		log.Printf("[CONSULTATION] Failed to connect to Redis: %v", err)
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	// Log a success message if the connection is successful
	log.Printf("[CONSULTATION] Connected to Redis at %s:%d, DB: %d", config.Host, config.Port, config.DB)

	// Return a RedisClient instance with the created Redis client
	return &RedisClient{client: redisClient}, nil
}

// Close closes the Redis client and returns an error if the operation fails.
func (rc *RedisClient) Close() error {
	// Close the Redis client
	if err := rc.client.Close(); err != nil {
		// Log an error if the client fails to close
		log.Printf("[CONSULTATION] Failed to close Redis client: %v", err)
		return fmt.Errorf("failed to close Redis client: %w", err)
	}

	// Log a success message if the client is closed successfully
	log.Println("[CONSULTATION] Redis client closed successfully")
	return nil
}

// GetClient returns the underlying Redis client.
func (rc *RedisClient) GetClient() *redis.Client {
	return rc.client
}
