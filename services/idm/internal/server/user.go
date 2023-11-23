package server

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/mihnea1711/POS_Project/services/idm/idm/proto_files"
	"github.com/mihnea1711/POS_Project/services/idm/internal/models"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
)

// GetUsers implements the GetUsers RPC method
func (s *MyIDMServer) GetUsers(ctx context.Context, req *proto_files.EmptyRequest) (*proto_files.UsersResponse, error) {
	// Extract pagination parameters from the request
	limit, page := utils.ExtractPaginationParams(req)

	// Ensure a database operation doesn't take longer than 5 seconds
	childCtx, cancel := context.WithTimeout(ctx, utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	log.Printf("[IDM] Received request to get users with pagination - Page: %d, Limit: %d", page, limit)

	// Call the database method to retrieve all users
	users, err := s.DbConn.GetAllUsers(childCtx, int(page), int(limit))
	if err != nil {
		log.Printf("[IDM] Error getting all users: %v", err)
		return nil, fmt.Errorf("error getting all users. %v", err)
	}

	log.Printf("[IDM] Retrieved %d users from the database", len(users))

	// Transform the database user models into proto user models
	var protoUsers []*proto_files.UserData
	for _, user := range users {
		protoUser := &proto_files.UserData{
			IDUser:   &proto_files.UserID{ID: int64(user.IDUser)},
			Username: user.Username,
		}
		protoUsers = append(protoUsers, protoUser)
	}

	log.Println("[IDM] Transformed database users into proto users")

	return &proto_files.UsersResponse{
		Users: protoUsers,
		Info: &proto_files.Info{
			Status:  http.StatusOK,
			Message: fmt.Sprintf("%d Users retrieved successfully", len(protoUsers)),
		},
	}, nil
}

// GetUserByID implements the GetUserByID RPC method
func (s *MyIDMServer) GetUserByID(ctx context.Context, req *proto_files.UserIDRequest) (*proto_files.UserResponse, error) {
	// Extract user ID from the request
	userID := req.UserID.ID

	// Ensure a database operation doesn't take longer than 5 seconds
	childCtx, cancel := context.WithTimeout(ctx, utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	log.Printf("[IDM] Received request to get user by ID: %d", userID)

	// Call the database method to retrieve the user by ID
	user, err := s.DbConn.GetUserByID(childCtx, int(userID))
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return &proto_files.UserResponse{
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
		// User not fetched
		log.Printf("[IDM] Did not fetch user with ID %d", userID)
		return &proto_files.UserResponse{
			Info: &proto_files.Info{
				Message: "User not fetched",
				Status:  http.StatusInternalServerError,
			},
		}, nil
	}

	// Populate the response with user data
	response := &proto_files.UserResponse{
		User: &proto_files.UserData{
			IDUser:   &proto_files.UserID{ID: int64(user.IDUser)},
			Username: user.Username,
		},
		Info: &proto_files.Info{
			Message: "User retrieved successfully",
			Status:  http.StatusOK,
		},
	}

	log.Printf("[IDM] Retrieved user with ID %d successfully", userID)
	return response, nil
}

// UpdateUserByID implements the UpdateUserByID RPC method
func (s *MyIDMServer) UpdateUserByID(ctx context.Context, req *proto_files.UpdateUserRequest) (*proto_files.EnhancedInfoResponse, error) {
	// Extract user ID and credentials from the request
	userID := req.UserData.IDUser.ID
	userCredentials := models.CredentialsRequest{
		Username: req.UserData.Username,
	}

	// Ensure a database operation doesn't take longer than 5 seconds
	childCtx, cancel := context.WithTimeout(ctx, utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	log.Printf("[IDM] Received request to update user by ID: %d", userID)

	// Call the database method to update the user by ID
	rowsAffected, err := s.DbConn.UpdateUserByID(childCtx, userCredentials, int(userID))
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return &proto_files.EnhancedInfoResponse{
				RowsAffected: int64(rowsAffected),
				Info: &proto_files.Info{
					Message: "User not found",
					Status:  http.StatusNotFound,
				},
			}, nil
		}
		// Check if the error is due to a duplicate entry violation
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok && mysqlErr.Number == utils.MySQLDuplicateEntryErrorCode {
			// Return a conflict error status
			return &proto_files.EnhancedInfoResponse{
				Info: &proto_files.Info{
					Message: "No changes made due to conflict",
					Status:  http.StatusConflict,
				},
			}, nil
		}
		log.Printf("[IDM] Error updating user by ID: %v", err)
		return nil, fmt.Errorf("error updating user. %v", err)
	}

	if rowsAffected == 0 {
		// No rows were affected, which means the user was not found
		log.Printf("[IDM] User with ID %d not found", userID)
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Message: "User not found. No changes made due to an unexpected error",
				Status:  http.StatusInternalServerError,
			},
		}, nil
	}

	// Return a success response
	log.Printf("[IDM] Updated user with ID %d successfully. Rows affected: %d", userID, rowsAffected)
	return &proto_files.EnhancedInfoResponse{
		RowsAffected: int64(rowsAffected),
		Info: &proto_files.Info{
			Message: "User updated successfully",
			Status:  http.StatusOK,
		},
	}, nil
}

// DeleteUserByID implements the DeleteUserByID RPC method
func (s *MyIDMServer) DeleteUserByID(ctx context.Context, req *proto_files.UserIDRequest) (*proto_files.EnhancedInfoResponse, error) {
	// Extract user ID from the request
	userID := req.UserID.ID

	// Ensure a database operation doesn't take longer than 5 seconds
	childCtx, cancel := context.WithTimeout(ctx, utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	log.Printf("[IDM] Received request to delete user by ID: %d", userID)

	// Call the database method to delete the user by ID
	rowsAffected, err := s.DbConn.DeleteUserByID(childCtx, int(userID))
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return &proto_files.EnhancedInfoResponse{
				RowsAffected: int64(rowsAffected),
				Info: &proto_files.Info{
					Message: "User not found",
					Status:  http.StatusNotFound,
				},
			}, nil
		}
		log.Printf("[IDM] Error deleting user by ID: %v", err)
		return nil, fmt.Errorf("error deleting user. %v", err)
	}

	if rowsAffected == 0 {
		// No rows were affected, which means the user was not found
		log.Printf("[IDM] User with ID %d not found or an unexpected error occured.", userID)
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Message: "User not found. No changes made due to an unexpected error",
				Status:  http.StatusNotFound,
			},
		}, nil
	}

	// Return a success response
	log.Printf("[IDM] Deleted user with ID %d successfully. Rows affected: %d", userID, rowsAffected)
	return &proto_files.EnhancedInfoResponse{
		RowsAffected: int64(rowsAffected),
		Info: &proto_files.Info{
			Message: "User deleted successfully",
			Status:  http.StatusOK,
		},
	}, nil
}

// / GetUserRole implements the GetUserRole RPC method
func (s *MyIDMServer) GetUserRole(ctx context.Context, req *proto_files.UserIDRequest) (*proto_files.RoleResponse, error) {
	// Extract user ID from the request
	userID := req.UserID.ID

	// Ensure a database operation doesn't take longer than 5 seconds
	childCtx, cancel := context.WithTimeout(ctx, utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	log.Printf("[IDM] Received request to get user's role by ID: %d", userID)

	// Call the database method to retrieve the user's role by ID
	userRole, err := s.DbConn.GetUserRoleByUserID(childCtx, int(userID))
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return &proto_files.RoleResponse{
				Role: userRole,
				Info: &proto_files.Info{
					Message: "User not found",
					Status:  http.StatusNotFound,
				},
			}, nil
		}
		log.Printf("[IDM] Error getting user's role: %v", err)
		return nil, err
	}

	if userRole == "" {
		// Role not found
		log.Printf("[IDM] User role not assigned for user with ID: %d", userID)
		return &proto_files.RoleResponse{
			Info: &proto_files.Info{
				Message: "User role not assigned or user has not been found in db",
				Status:  http.StatusNotFound,
			},
		}, nil
	}

	// Return the user's role in the response
	log.Printf("[IDM] User role retrieved successfully for user with ID: %d", userID)
	return &proto_files.RoleResponse{
		Role: userRole,
		Info: &proto_files.Info{
			Message: "User role retrieved successfully",
			Status:  http.StatusOK,
		},
	}, nil
}

// UpdateUserRole implements the UpdateUserRole RPC method
func (s *MyIDMServer) UpdateUserRole(ctx context.Context, req *proto_files.UpdateRoleRequest) (*proto_files.EnhancedInfoResponse, error) {
	// Extract user ID and new role from the request
	userID := req.UserID.ID
	newRole := req.Role

	// Ensure a database operation doesn't take longer than 5 seconds
	childCtx, cancel := context.WithTimeout(ctx, utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	log.Printf("[IDM] Received request to update user's role. UserID: %d, NewRole: %s", userID, newRole)

	// Call the database method to change the user's role
	rowsAffected, err := s.DbConn.UpdateUserRoleByUserID(childCtx, int(userID), newRole)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return &proto_files.EnhancedInfoResponse{
				RowsAffected: int64(rowsAffected),
				Info: &proto_files.Info{
					Message: "User not found",
					Status:  http.StatusNotFound,
				},
			}, nil
		}
		log.Printf("[IDM] Error changing user's role: %v", err)
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Message: "Error changing user's role",
				Status:  http.StatusInternalServerError,
			},
		}, err
	}

	if rowsAffected == 0 {
		// No rows were affected, which means the user was not found
		log.Printf("[IDM] User not found or no changes made. UserID: %d, NewRole: %s", userID, newRole)
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Message: "User not found or no changes made",
				Status:  http.StatusNotFound,
			},
		}, nil
	}

	// Return a success response
	log.Printf("[IDM] User role updated successfully. UserID: %d, NewRole: %s", userID, newRole)
	return &proto_files.EnhancedInfoResponse{
		RowsAffected: int64(rowsAffected),
		Info: &proto_files.Info{
			Message: "User role updated successfully",
			Status:  http.StatusOK,
		},
	}, nil
}

// GetUserPassword implements the GetUserPassword RPC method
func (s *MyIDMServer) GetUserPassword(ctx context.Context, req *proto_files.UsernameRequest) (*proto_files.PasswordResponse, error) {
	// Extract username from the request
	username := req.Username

	// Ensure a database operation doesn't take longer than 5 seconds
	childCtx, cancel := context.WithTimeout(ctx, utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	log.Printf("[IDM] Received request to get user's password. Username: %s", username)

	// Call the database method to retrieve the user's password by username
	password, err := s.DbConn.GetUserPasswordByUsername(childCtx, username)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return &proto_files.PasswordResponse{
				Password: password,
				Info: &proto_files.Info{
					Message: "User not found",
					Status:  http.StatusNotFound,
				},
			}, nil
		}
		log.Printf("[IDM] Error getting user's password: %v", err)
		return nil, err
	}

	if password == "" {
		log.Printf("[IDM] Error retrieving password for user with Username %s from DB", username)
		return nil, fmt.Errorf("user %s has not been found or an unexpected error happened", username)
	}

	// Return the user's password in the response
	log.Printf("[IDM] User password retrieved successfully. Username: %s", username)
	return &proto_files.PasswordResponse{
		Info: &proto_files.Info{
			Message: "User password retrieved successfully",
			Status:  http.StatusOK,
		},
		Password: password,
	}, nil
}

// UpdateUserPassword implements the UpdateUserPassword RPC method
func (s *MyIDMServer) UpdateUserPassword(ctx context.Context, req *proto_files.UpdatePasswordRequest) (*proto_files.EnhancedInfoResponse, error) {
	// Extract user ID and new password from the request
	userID := req.UserID.ID
	newPassword := req.Password

	// Ensure a database operation doesn't take longer than 5 seconds
	childCtx, cancel := context.WithTimeout(ctx, utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	log.Printf("[IDM] Received request to update user's password. UserID: %d", userID)

	// Hash the new password before updating it in the database
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		log.Printf("[IDM] Error hashing password: %v", err)
		return nil, fmt.Errorf("error hashing password. %v", err)
	}

	// Call the database method to change the user's password
	rowsAffected, err := s.DbConn.UpdateUserPasswordByUserID(childCtx, int(userID), hashedPassword)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in DB")
			return &proto_files.EnhancedInfoResponse{
				RowsAffected: int64(rowsAffected),
				Info: &proto_files.Info{
					Message: "User not found",
					Status:  http.StatusNotFound,
				},
			}, nil
		}
		log.Printf("[IDM] Error changing user's password: %v", err)
		return nil, fmt.Errorf("error changing user's password")
	}

	if rowsAffected == 0 {
		// No rows were affected, which means the user was not found
		log.Printf("[IDM] User not found or no changes made. UserID: %d", userID)
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Message: "User not found or no changes made",
				Status:  http.StatusNotFound,
			},
		}, nil
	}

	// Return a success response
	log.Printf("[IDM] User password updated successfully. UserID: %d", userID)
	return &proto_files.EnhancedInfoResponse{
		RowsAffected: int64(rowsAffected),
		Info: &proto_files.Info{
			Message: "User password updated successfully",
			Status:  http.StatusOK,
		},
	}, nil
}
