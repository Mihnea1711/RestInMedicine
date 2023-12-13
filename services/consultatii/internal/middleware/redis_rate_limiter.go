package middleware

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/database/redis"
	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
)

type RedisRateLimiter struct {
	rdb            *redis.RedisClient
	context        context.Context
	rate           int
	windowDuration time.Duration
}

func NewRedisRateLimiter(ctx context.Context, rdb *redis.RedisClient, rate int, windowDuration time.Duration) *RedisRateLimiter {
	return &RedisRateLimiter{
		rdb:            rdb,
		context:        ctx,
		rate:           rate,
		windowDuration: windowDuration,
	}
}

func (r *RedisRateLimiter) getKey(ip string) string {
	return fmt.Sprintf("rate_limit:%s:%d", ip, time.Now().Unix()/int64(r.windowDuration.Seconds()))
}

func (r *RedisRateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ip, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			log.Printf("[CONSULTATION_LIMITER] Error splitting remote addr %s", err)
			utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
				Error:   err.Error(),
				Message: "Internal Server Error",
			})
			return
		}
		key := r.getKey(ip)

		val, err := r.rdb.GetClient().Incr(r.context, key).Result()
		if err != nil {
			log.Printf("[CONSULTATION_LIMITER] Error incrementing rate limit key %s: %v", key, err)
			utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
				Error:   err.Error(),
				Message: "Internal Server Error",
			})
			return
		}
		if val == 1 {
			// The key is new, set its TTL
			expireCmd := r.rdb.GetClient().Expire(r.context, key, r.windowDuration)
			if expireCmd.Err() != nil {
				log.Printf("[CONSULTATION_LIMITER] Error setting TTL for rate limit key %s: %v", key, expireCmd.Err())
			} else if !expireCmd.Val() {
				log.Printf("[CONSULTATION_LIMITER] Key %s does not exist, could not set TTL.", key)
			} else {
				log.Printf("[CONSULTATION_LIMITER] New key %s created with TTL of %v", key, r.windowDuration)
			}
		}

		if val > int64(r.rate) {
			log.Printf("[CONSULTATION_LIMITER] Rate limit exceeded for IP %s", req.RemoteAddr)
			utils.RespondWithJSON(w, http.StatusTooManyRequests, models.ResponseData{
				Message: "Too Many Requests",
			})
			return
		}

		next.ServeHTTP(w, req)
	})
}
