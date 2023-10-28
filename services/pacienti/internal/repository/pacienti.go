package repository

import "github.com/redis/go-redis"

type RedisRepo struct {
	Client *redis.Client
}
