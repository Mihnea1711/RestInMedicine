package redis

import (
	"context"
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/idm/internal/models"
	"github.com/redis/go-redis/v9"
)

// AddUserToBlacklistInRedis adds a user to the Redis blacklist.
func (rc *RedisClient) AddUserToBlacklistInRedis(ctx context.Context, blacklistUserModel models.BlacklistToken) error {
	// Generate a key to uniquely identify the user in the blacklist
	blacklistKey := fmt.Sprintf("blacklist:%d", blacklistUserModel.IDUser)

	// Use a transaction to set the user's JWT token as the value of the key
	err := rc.client.Watch(ctx, func(tx *redis.Tx) error {
		// Check if the user is already in the blacklist
		_, err := tx.Get(ctx, blacklistKey).Result()
		if err == redis.Nil {
			// If the user is not in the blacklist, set the key and value
			if err := tx.Set(ctx, blacklistKey, blacklistUserModel.Token, 0).Err(); err != nil {
				return err
			}
			log.Printf("[IDM] Added user (ID %d) to Redis blacklist", blacklistUserModel.IDUser)
		} else if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		log.Printf("[IDM] Error adding user (ID %d) to Redis blacklist: %v", blacklistUserModel.IDUser, err)
		return err
	}

	return nil
}

// need further testing (not necessary)
/*
func (rc *RedisClient) AddUserToBlacklistInRedis(ctx context.Context, blacklistUserModel models.BlacklistToken) error {
    // Generate a key to uniquely identify the user in the blacklist
    blacklistKey := fmt.Sprintf("blacklist:%d", blacklistUserModel.IDUser)

    // Use a transaction to set the user's data in the Redis Hash
    err := rc.client.Watch(ctx, func(tx *redis.Tx) error {
        // Check if the user is already in the blacklist
        _, err := tx.HGetAll(ctx, blacklistKey).Result()
        if err == redis.Nil {
            // If the user is not in the blacklist, set the key and value
            values := map[string]interface{}{
                "token":  blacklistUserModel.Token,
                "reason": blacklistUserModel.Reason,
            }
            if err := tx.HMSet(ctx, blacklistKey, values).Err(); err != nil {
                return err
            }
            log.Printf("[IDM] Added user (ID %d) to Redis blacklist", blacklistUserModel.IDUser)
        } else if err != nil {
            return err
        }

        return nil
    })

    if err != nil {
        log.Printf("[IDM] Error adding user (ID %d) to Redis blacklist: %v", blacklistUserModel.IDUser, err)
        return err
    }

    return nil
}
*/

// RemoveUserFromBlacklistInRedis removes a user from the Redis blacklist and returns the number of rows affected.
func (rc *RedisClient) RemoveUserFromBlacklistInRedis(ctx context.Context, blacklistUserModel models.BlacklistToken) (int64, error) {
	// Generate a key to uniquely identify the user in the blacklist
	blacklistKey := fmt.Sprintf("blacklist:%d", blacklistUserModel.IDUser)

	// Use Redis command to delete the user from the blacklist and count the number of removed items
	rowsAffected, err := rc.client.Del(ctx, blacklistKey).Result()
	if err != nil {
		log.Printf("[IDM] Error removing user (ID %d) from Redis blacklist: %v", blacklistUserModel.IDUser, err)
		return 0, err
	}

	log.Printf("[IDM] Removed user (ID %d) from Redis blacklist, rows affected: %d", blacklistUserModel.IDUser, rowsAffected)
	return rowsAffected, nil
}

// IsUserInBlacklist checks if a user is in the Redis blacklist.
func (rc *RedisClient) IsUserInBlacklist(ctx context.Context, userID int) (bool, error) {
	// Generate a key to uniquely identify the user in the blacklist
	blacklistKey := fmt.Sprintf("blacklist:%d", userID)

	// Use Redis command to check if the key exists
	exists, err := rc.client.Exists(ctx, blacklistKey).Result()
	if err != nil {
		log.Printf("[IDM] Error checking if user (ID %d) is in Redis blacklist: %v", userID, err)
		return false, err
	}

	return exists == 1, nil
}

// IsTokenBlacklisted checks if a token is in the Redis blacklist.
func (rc *RedisClient) IsTokenBlacklisted(ctx context.Context, token string) (bool, error) {
	// Use Redis command to check if the token exists in the blacklist
	exists, err := rc.client.SIsMember(ctx, "blacklist:tokens", token).Result()
	if err != nil {
		log.Printf("[IDM] Error checking if token is in Redis blacklist: %v", err)
		return false, err
	}

	return exists, nil
}
