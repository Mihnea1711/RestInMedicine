package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/idm/internal/models"
	"github.com/redis/go-redis/v9"
)

// AddUserToBlacklistInRedis adds a token to the Redis blacklist.
func (rc *RedisClient) AddTokenToBlacklistInRedis(ctx context.Context, blacklistUserModel models.BlacklistToken) error {
	// Generate a key to uniquely identify the token in the blacklist
	blacklistKey := fmt.Sprintf("blacklist:%s", blacklistUserModel.Token)

	// Use a transaction to set the token's JWT token as the value of the key
	err := rc.client.Watch(ctx, func(tx *redis.Tx) error {
		// Check if the token is already in the blacklist
		_, err := tx.Get(ctx, blacklistKey).Result()
		if err == redis.Nil {
			// If the token is not in the blacklist, set the key and value
			if err := tx.Set(ctx, blacklistKey, blacklistUserModel.Token, 0).Err(); err != nil {
				log.Printf("[IDM] Failed to set Redis key for token (Token %s): %v", blacklistUserModel.Token, err)
				return err
			}
			log.Printf("[IDM] Added (Token %s) to Redis blacklist", blacklistUserModel.Token)
		} else if err != nil {
			log.Printf("[IDM] Error checking Redis blacklist for (Token %s): %v", blacklistUserModel.Token, err)
			return err
		}

		return nil
	})

	if err != nil {
		log.Printf("[IDM] Error adding (Token %s) to Redis blacklist: %v", blacklistUserModel.Token, err)
		return err
	}

	return nil
}

// RemoveUserFromBlacklistInRedis removes a token from the Redis blacklist and returns the number of rows affected.
func (rc *RedisClient) RemoveTokenFromBlacklistInRedis(ctx context.Context, blacklistUserModel models.BlacklistToken) (int64, error) {
	// Generate a key to uniquely identify the token in the blacklist
	blacklistKey := fmt.Sprintf("blacklist:%s", blacklistUserModel.Token)

	// Use Redis command to delete the token from the blacklist and count the number of removed items
	rowsAffected, err := rc.client.Del(ctx, blacklistKey).Result()
	if err != nil {
		log.Printf("[IDM] Error removing (Token %s) from Redis blacklist: %v", blacklistUserModel.Token, err)
		return 0, err
	}

	if rowsAffected > 0 {
		log.Printf("[IDM] Removed token (Token %s) from Redis blacklist, rows affected: %d", blacklistUserModel.Token, rowsAffected)
	} else {
		log.Printf("[IDM] Token %s not found in Redis blacklist, no rows affected", blacklistUserModel.Token)
	}

	return rowsAffected, nil
}

// IsTokenBlacklisted checks if a token is in the Redis blacklist.
func (rc *RedisClient) IsTokenBlacklisted(ctx context.Context, token string) (bool, error) {
	// Use Redis command to check if the token exists in the blacklist
	exists, err := rc.client.SIsMember(ctx, "blacklist:%s", token).Result()
	if err != nil {
		log.Printf("[IDM] Error checking if token is in Redis blacklist: %v", err)
		return false, err
	}

	if exists {
		log.Printf("[IDM] Token is in Redis blacklist")
	} else {
		log.Printf("[IDM] Token is not in Redis blacklist")
	}

	return exists, nil
}
