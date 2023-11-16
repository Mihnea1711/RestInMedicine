package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/programari/pkg/config"
	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(ctx context.Context, config *config.RedisConfig) (*RedisClient, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", config.Host, config.Port),
		Password: config.Password,
		DB:       config.DB,
	})

	_, err := redisClient.Ping(ctx).Result()
	if err != nil {
		log.Printf("[APPOINTMENT] Error connecting to Redis: %v", err)
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	log.Printf("[APPOINTMENT] Connected to Redis on %s:%d", config.Host, config.Port)

	return &RedisClient{client: redisClient}, nil
}

func (rc *RedisClient) GetClient() *redis.Client {
	return rc.client
}

func (rc *RedisClient) Close() error {
	if err := rc.client.Close(); err != nil {
		log.Printf("[APPOINTMENT] Error closing Redis client: %v", err)
		return fmt.Errorf("failed to close Redis client: %w", err)
	}
	log.Printf("[APPOINTMENT] Closed Redis client")
	return nil
}
