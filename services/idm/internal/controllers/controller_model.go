package controllers

import (
	"github.com/mihnea1711/POS_Project/services/idm/internal/database"
	"github.com/mihnea1711/POS_Project/services/idm/internal/database/redis"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/config"
)

type IDMController struct {
	DbConn    database.Database
	RedisConn *redis.RedisClient
	jwtconfig config.JWTConfig
}
