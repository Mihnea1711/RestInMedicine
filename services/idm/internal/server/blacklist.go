package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/idm/idm/proto_files"
	"github.com/mihnea1711/POS_Project/services/idm/internal/models"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
)

// AddTokenToBlacklist implements the AddTokenToBlacklist RPC method
func (s *MyIDMServer) AddTokenToBlacklist(ctx context.Context, req *proto_files.BlacklistRequest) (*proto_files.InfoResponse, error) {
	blacklistModel := models.BlacklistToken{
		Token: req.Token,
	}

	// Ensure a database operation doesn't take longer than 5 seconds
	childCtx, cancel := context.WithTimeout(ctx, utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Check if the token is in the Redis blacklist
	isInBlacklist, err := s.RedisConn.IsTokenBlacklisted(childCtx, blacklistModel.Token)
	if err != nil {
		// Handle the error and return an error response
		log.Printf("[IDM] Error checking if token is in blacklist: %v", err)
		return nil, fmt.Errorf("error checking if token is in blacklist. %v", err)
	}

	if isInBlacklist {
		// Handle the case where the token is in the blacklist
		log.Printf("[IDM] Token %s is already in the blacklist", blacklistModel.Token)
		return &proto_files.InfoResponse{
			Info: &proto_files.Info{
				Message: "Token is already in the blacklist",
				Status:  http.StatusConflict,
			},
		}, nil
	}

	// Use the Redis client to add the token to the blacklist.
	err = s.RedisConn.AddTokenToBlacklistInRedis(ctx, blacklistModel)
	if err != nil {
		log.Printf("[IDM] Error adding token to blacklist: %v", err)
		return nil, fmt.Errorf("error adding token to blacklist. %v", err)
	}

	// Handle a successful addition to the blacklist.
	return &proto_files.InfoResponse{
		Info: &proto_files.Info{
			Message: "Token added to the blacklist successfully",
			Status:  int64(http.StatusOK),
		},
	}, nil
}

// CheckTokenInBlacklist implements the CheckTokenInBlacklist RPC method
func (s *MyIDMServer) CheckTokenInBlacklist(ctx context.Context, req *proto_files.BlacklistRequest) (*proto_files.InfoResponse, error) {
	token := req.Token

	// Ensure a database operation doesn't take longer than 5 seconds
	childCtx, cancel := context.WithTimeout(ctx, utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Check if the token is in the Redis blacklist
	isInBlacklist, err := s.RedisConn.IsTokenBlacklisted(childCtx, token)
	if err != nil {
		// Handle the error and return an error response
		log.Printf("[IDM] Error checking if token is in blacklist: %v", err)
		return nil, fmt.Errorf("error checking if token is in blacklist. %v", err)
	}

	if isInBlacklist {
		// Handle the case where the token is in the blacklist
		log.Printf("[IDM] Token %s is in the blacklist", token)
		return &proto_files.InfoResponse{
			Info: &proto_files.Info{
				Message: "Token is in the blacklist",
				Status:  http.StatusForbidden,
			},
		}, nil
	}

	// Handle the case where the token is not in the blacklist
	return &proto_files.InfoResponse{
		Info: &proto_files.Info{
			Message: "Token is not in the blacklist",
			Status:  http.StatusOK,
		},
	}, nil
}

// RemoveTokenFromBlacklist implements the RemoveTokenFromBlacklist RPC method
func (s *MyIDMServer) RemoveTokenFromBlacklist(ctx context.Context, req *proto_files.BlacklistRequest) (*proto_files.EnhancedInfoResponse, error) {
	blacklistModel := models.BlacklistToken{
		Token: req.Token,
	}

	// Ensure a database operation doesn't take longer than 5 seconds
	childCtx, cancel := context.WithTimeout(ctx, utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Use the Redis client to remove the token from the blacklist and get the number of rows affected.
	rowsAffected, err := s.RedisConn.RemoveTokenFromBlacklistInRedis(childCtx, blacklistModel)
	if err != nil {
		log.Printf("[IDM] Error removing token from blacklist: %v", err)
		// Handle the error and return an appropriate gRPC response
		return nil, fmt.Errorf("failed to remove token from blacklist. %v", err)
	}

	if rowsAffected == 0 {
		// No rows were affected, which means the token was not in the blacklist.
		log.Printf("[IDM] Token %s is not in the blacklist", blacklistModel.Token)
		// Handle the case where the token was not found in the blacklist
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Message: "Token not found in the blacklist",
				Status:  int64(http.StatusNotFound),
			},
		}, nil
	}

	// Handle a successful removal from the blacklist.
	return &proto_files.EnhancedInfoResponse{
		RowsAffected: rowsAffected,
		Info: &proto_files.Info{
			Message: "Token removed from the blacklist successfully",
			Status:  int64(http.StatusOK),
		},
	}, nil
}
