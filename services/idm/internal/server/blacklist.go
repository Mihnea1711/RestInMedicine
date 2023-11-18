package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/idm/idm/proto_files"
	"github.com/mihnea1711/POS_Project/services/idm/internal/models"
)

// AddUserToBlacklist implements the AddUserToBlacklist RPC method
func (s *MyIDMServer) AddUserToBlacklist(ctx context.Context, req *proto_files.BlacklistRequest) (*proto_files.InfoResponse, error) {
	blacklistUserModel := models.BlacklistToken{
		IDUser: int(req.UserID.ID),
		Token:  req.Token,
	}

	// Use the Redis client to add the user to the blacklist.
	err := s.RedisConn.AddUserToBlacklistInRedis(ctx, blacklistUserModel)
	if err != nil {
		log.Printf("[IDM] Error adding user to blacklist: %v", err)
		return nil, fmt.Errorf("error adding user to blacklist. %v", err)
	}

	// Handle a successful addition to the blacklist.
	return &proto_files.InfoResponse{
		Info: &proto_files.Info{
			Message: "User added to the blacklist successfully",
			Status:  int64(http.StatusOK),
		},
	}, nil
}

// CheckUserInBlacklist implements the CheckUserInBlacklist RPC method
func (s *MyIDMServer) CheckUserInBlacklist(ctx context.Context, req *proto_files.UserIDRequest) (*proto_files.InfoResponse, error) {
	userID := req.UserID.ID

	// Check if the user is in the Redis blacklist
	isInBlacklist, err := s.RedisConn.IsUserInBlacklist(ctx, int(userID))
	if err != nil {
		// Handle the error and return an error response
		return nil, fmt.Errorf("error checking if user is in blacklist. %v", err)
	}

	if isInBlacklist {
		// Handle the case where the user is in the blacklist
		return &proto_files.InfoResponse{
			Info: &proto_files.Info{
				Message: "User is in the blacklist",
				Status:  http.StatusForbidden,
			},
		}, nil
	}

	// Handle the case where the user is not in the blacklist
	return &proto_files.InfoResponse{
		Info: &proto_files.Info{
			Message: "User is not in the blacklist",
			Status:  http.StatusOK,
		},
	}, nil
}

// RemoveUserFromBlacklist implements the RemoveUserFromBlacklist RPC method
func (s *MyIDMServer) RemoveUserFromBlacklist(ctx context.Context, req *proto_files.UserIDRequest) (*proto_files.EnhancedInfoResponse, error) {
	blacklistUserModel := models.BlacklistToken{
		IDUser: int(req.UserID.ID),
	}

	// Use the Redis client to remove the user from the blacklist and get the number of rows affected.
	rowsAffected, err := s.RedisConn.RemoveUserFromBlacklistInRedis(ctx, blacklistUserModel)
	if err != nil {
		log.Printf("[IDM] Error removing user from blacklist: %v", err)
		// Handle the error and return an appropriate gRPC response
		return nil, fmt.Errorf("failed to remove user from blacklist. %v", err)
	}

	if rowsAffected == 0 {
		// No rows were affected, which means the user was not in the blacklist.
		log.Printf("[IDM] User with ID %d is not in the blacklist", blacklistUserModel.IDUser)
		// Handle the case where the user was not found in the blacklist
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Message: "User not found in the blacklist",
				Status:  int64(http.StatusNotFound),
			},
		}, nil
	}

	// Handle a successful removal from the blacklist.
	return &proto_files.EnhancedInfoResponse{
		RowsAffected: rowsAffected,
		Info: &proto_files.Info{
			Message: "User removed from the blacklist successfully",
			Status:  int64(http.StatusOK),
		},
	}, nil
}
