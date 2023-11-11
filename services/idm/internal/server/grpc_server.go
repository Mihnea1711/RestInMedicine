package server

import (
	"github.com/mihnea1711/POS_Project/services/idm/idm"
	"github.com/mihnea1711/POS_Project/services/idm/internal/database"
	"github.com/mihnea1711/POS_Project/services/idm/internal/database/redis"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/config"
)

type MyIDMServer struct {
	DbConn    database.Database
	RedisConn *redis.RedisClient
	JwtConfig config.JWTConfig
	idm.UnimplementedIDMServer
}
