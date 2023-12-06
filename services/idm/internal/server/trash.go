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

func (s *MyIDMServer) RestoreUserByID(ctx context.Context, req *proto_files.UserIDRequest) (*proto_files.EnhancedInfoResponse, error) {
	// Extract user ID from the request
	userID := req.UserID.ID

	// Ensure a database operation doesn't take longer than 5 seconds
	childCtx, cancel := context.WithTimeout(ctx, utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	log.Println("[IDM] Received request to restore user from trash")

	// Get user data from the trash
	user, err := s.DbConn.GetDataFromTrashByUserID(childCtx, int(userID))
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in trash")
			return &proto_files.EnhancedInfoResponse{
				Info: &proto_files.Info{
					Message: "User not found in trash",
					Status:  http.StatusNotFound,
				},
			}, nil
		}

		log.Printf("[IDM] Error getting user from trash by ID: %v", err)
		return nil, fmt.Errorf("error getting user from trash by ID: %v", err)
	}
	if user == nil {
		// User not fetched from trash
		log.Printf("[IDM] Did not fetch user with ID %d from trash", userID)
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Message: "User not fetched from trash",
				Status:  http.StatusInternalServerError,
			},
		}, nil
	}

	// Remove user from trash
	removeResponse, err := s.DbConn.RemoveUserFromTrash(ctx, user.IDUser)
	if err != nil {
		log.Printf("[IDM] Error removing user from trash: %v", err)
		return nil, fmt.Errorf("error removing user from trash: %v", err)
	}
	// Check if no rows were affected, indicating that the user was not found in trash
	if removeResponse == 0 {
		log.Printf("[IDM] User with ID %d not found in trash", userID)
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Message: "User not found in trash",
				Status:  http.StatusNotFound,
			},
		}, nil
	}

	// Add user to the User table and Role table
	addResponse, err := s.DbConn.AddUser(ctx, models.UserRegistration{
		Username: user.Username,
		Password: user.Password,
		Role:     user.Role,
	})
	if err != nil {
		log.Printf("[IDM] Error adding user to User and Role tables: %v", err)
		return nil, fmt.Errorf("error adding user to User and Role tables: %v", err)
	}
	// Check if no rows were affected, indicating that the user was not found in trash
	if addResponse == 0 {
		log.Printf("[IDM] User with ID %d has not been added back", userID)
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Message: "User not added back",
				Status:  http.StatusInternalServerError,
			},
		}, nil
	}

	// Return the response
	return &proto_files.EnhancedInfoResponse{
		RowsAffected: int64(removeResponse),
		Info: &proto_files.Info{
			Message: "User successfully removed from trash and restored to its original form",
			Status:  http.StatusOK,
		},
	}, nil
}
