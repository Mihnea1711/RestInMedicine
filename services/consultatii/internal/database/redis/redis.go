package redis

import (
	"context"
	"fmt"

	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/config"
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
		return nil, fmt.Errorf("failed to connect to Redis: %w", err)
	}

	return &RedisClient{client: redisClient}, nil
}

func (rc *RedisClient) Close() error {
	if err := rc.client.Close(); err != nil {
		return fmt.Errorf("failed to close Redis client: %w", err)
	}
	return nil
}

func (rc *RedisClient) GetClient() *redis.Client {
	return rc.client
}
