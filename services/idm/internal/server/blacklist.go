package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/idm/idm/proto_files"
	"github.com/mihnea1711/POS_Project/services/idm/internal/models"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
)

// AddUserToBlacklist implements the AddUserToBlacklist RPC method
func (s *MyIDMServer) AddUserToBlacklist(ctx context.Context, req *proto_files.BlacklistRequest) (*proto_files.InfoResponse, error) {
	blacklistUserModel := models.BlacklistToken{
		IDUser: int(req.UserID.ID),
		Token:  req.Token,
	}

	// Ensure a database operation doesn't take longer than 5 seconds
	childCtx, cancel := context.WithTimeout(ctx, utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Call the database method to retrieve the user by ID
	user, err := s.DbConn.GetUserByID(childCtx, int(blacklistUserModel.IDUser))
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return &proto_files.InfoResponse{
				Info: &proto_files.Info{
					Message: "User not found",
					Status:  http.StatusNotFound,
				},
			}, nil
		}
		log.Printf("[IDM] Error getting user by ID: %v", err)
		return nil, fmt.Errorf("error getting user by ID. %v", err)
	}

	if user == nil {
		// User not found
		log.Printf("[IDM] User ID does not exist. User not added: %v", err)
		return &proto_files.InfoResponse{
			Info: &proto_files.Info{
				Message: "User ID does not exist. User not added to blacklist",
				Status:  int64(http.StatusNotFound),
			},
		}, nil
	}

	// Check if the user is in the Redis blacklist
	isInBlacklist, err := s.RedisConn.IsUserInBlacklist(ctx, int(user.IDUser))
	if err != nil {
		// Handle the error and return an error response
		log.Printf("[IDM] Error checking if user is in blacklist: %v", err)
		return nil, fmt.Errorf("error checking if user is in blacklist. %v", err)
	}

	if isInBlacklist {
		// Handle the case where the user is in the blacklist
		log.Printf("[IDM] User %d is already in the blacklist", int(user.IDUser))
		return &proto_files.InfoResponse{
			Info: &proto_files.Info{
				Message: "User is already in the blacklist",
				Status:  http.StatusConflict,
			},
		}, nil
	}

	// Use the Redis client to add the user to the blacklist.
	err = s.RedisConn.AddUserToBlacklistInRedis(ctx, blacklistUserModel)
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

	// Ensure a database operation doesn't take longer than 5 seconds
	childCtx, cancel := context.WithTimeout(ctx, utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Call the database method to retrieve the user by ID
	user, err := s.DbConn.GetUserByID(childCtx, int(userID))
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return &proto_files.InfoResponse{
				Info: &proto_files.Info{
					Message: "User not found",
					Status:  http.StatusNotFound,
				},
			}, nil
		}
		log.Printf("[IDM] Error getting user by ID: %v", err)
		return nil, fmt.Errorf("error getting user by ID. %v", err)
	}

	if user == nil {
		// User not found
		log.Printf("[IDM] User ID does not exist or an unexpected error occured.: %v", err)
		return &proto_files.InfoResponse{
			Info: &proto_files.Info{
				Message: "User ID does not exist or an unexpected error occured.",
				Status:  int64(http.StatusInternalServerError),
			},
		}, nil
	}

	// Check if the user is in the Redis blacklist
	isInBlacklist, err := s.RedisConn.IsUserInBlacklist(ctx, int(userID))
	if err != nil {
		// Handle the error and return an error response
		log.Printf("[IDM] Error checking if user is in blacklist: %v", err)
		return nil, fmt.Errorf("error checking if user is in blacklist. %v", err)
	}

	if isInBlacklist {
		// Handle the case where the user is in the blacklist
		log.Printf("[IDM] User %d is in the blacklist", int(userID))
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

	// Ensure a database operation doesn't take longer than 5 seconds
	childCtx, cancel := context.WithTimeout(ctx, utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Call the database method to retrieve the user by ID
	user, err := s.DbConn.GetUserByID(childCtx, int(blacklistUserModel.IDUser))
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return &proto_files.EnhancedInfoResponse{
				RowsAffected: 0,
				Info: &proto_files.Info{
					Message: "User not found",
					Status:  http.StatusNotFound,
				},
			}, nil
		}
		log.Printf("[IDM] Error getting user by ID: %v", err)
		return nil, fmt.Errorf("error getting user by ID. %v", err)
	}

	if user == nil {
		// User not found
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Message: "User ID does not exist. User not removed",
				Status:  int64(http.StatusNotFound),
			},
		}, nil
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
