package middleware

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/idm/internal/database/redis"
)

type RedisRateLimiter struct {
	rdb            *redis.RedisClient
	rate           int
	windowDuration time.Duration
}

func NewRedisRateLimiter(rdb *redis.RedisClient, rate int, windowDuration time.Duration) *RedisRateLimiter {
	return &RedisRateLimiter{
		rdb:            rdb,
		rate:           rate,
		windowDuration: windowDuration,
	}
}

func (r *RedisRateLimiter) getKey(ip string) string {
	return fmt.Sprintf("rate_limit:%s:%d", ip, time.Now().Unix()/int64(r.windowDuration.Seconds()))
}

func (r *RedisRateLimiter) Limit(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		ctx := context.Background()

		ip, _, err := net.SplitHostPort(req.RemoteAddr)
		if err != nil {
			log.Printf("[IDM] Error splitting remote addr %s", err) // Logging the error
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		key := r.getKey(ip)

		val, err := r.rdb.GetClient().Incr(ctx, key).Result()
		if err != nil {
			log.Printf("[IDM] Error incrementing rate limit key %s: %v", key, err) // Logging the error
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		if val == 1 {
			// The key is new, set its TTL
			expireCmd := r.rdb.GetClient().Expire(ctx, key, r.windowDuration)
			if expireCmd.Err() != nil {
				log.Printf("[IDM] Error setting TTL for rate limit key %s: %v", key, expireCmd.Err())
			} else if !expireCmd.Val() {
				log.Printf("[IDM] Key %s does not exist, could not set TTL.", key)
			} else {
				log.Printf("[IDM] New key %s created with TTL of %v", key, r.windowDuration)
			}
		}

		if val > int64(r.rate) {
			log.Printf("[IDM] Rate limit exceeded for IP %s", req.RemoteAddr)
			http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
			return
		}

		next.ServeHTTP(w, req)
	})
}
