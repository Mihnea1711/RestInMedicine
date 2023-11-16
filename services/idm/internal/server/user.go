package server

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/idm/idm/proto_files"
	"github.com/mihnea1711/POS_Project/services/idm/internal/models"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
)

// GetUsers implements the GetUsers RPC method
func (s *MyIDMServer) GetUsers(ctx context.Context, req *proto_files.EmptyRequest) (*proto_files.UsersResponse, error) {
	page := req.Page
	limit := req.Limit

	// Call the database method to retrieve all users
	users, err := s.DbConn.GetAllUsers(int(page), int(limit))
	if err != nil {
		log.Printf("[IDM] Error getting all users: %v", err)
		return &proto_files.UsersResponse{
			Info: &proto_files.Info{
				Status:  http.StatusInternalServerError,
				Message: "Internal Server Error",
				Error:   "Error getting all users",
			},
		}, err
	}

	// Transform the database user models into proto user models
	var protoUsers []*proto_files.UserData
	for _, user := range users {
		protoUser := &proto_files.UserData{
			IDUser:   &proto_files.UserID{ID: int64(user.IDUser)},
			Username: user.Username,
		}
		protoUsers = append(protoUsers, protoUser)
	}

	return &proto_files.UsersResponse{
		Users: protoUsers,
		Info: &proto_files.Info{
			Status:  http.StatusOK,
			Message: "Users retrieved successfully",
		},
	}, nil
}

// GetUserByID implements the GetUserByID RPC method
func (s *MyIDMServer) GetUserByID(ctx context.Context, req *proto_files.UserIDRequest) (*proto_files.UserResponse, error) {
	userID := req.UserID.ID

	// Call the database method to retrieve the user by ID
	user, err := s.DbConn.GetUserByID(int(userID))
	if err != nil {
		log.Printf("[IDM] Error getting user by ID: %v", err)
		// Handle the error, maybe return an error response
		return &proto_files.UserResponse{
			Info: &proto_files.Info{
				Message: "Error getting user by ID",
				Status:  http.StatusInternalServerError,
				Error:   "Internal Server Error",
			},
		}, err
	}

	if user.IDUser == 0 {
		// User not found
		// Handle the not found case, maybe return a not found response
		return &proto_files.UserResponse{
			Info: &proto_files.Info{
				Message: "User not found",
				Status:  http.StatusNotFound,
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

	return response, nil
}

// UpdateUserByID implements the UpdateUserByID RPC method
func (s *MyIDMServer) UpdateUserByID(ctx context.Context, req *proto_files.UpdateUserRequest) (*proto_files.EnhancedInfoResponse, error) {
	userID := req.UserData.IDUser.ID

	userCredentials := models.CredentialsRequest{
		Username: req.UserData.Username,
	}

	// Call the database method to update the user by ID
	rowsAffected, err := s.DbConn.UpdateUserByID(userCredentials, int(userID))
	if err != nil {
		log.Printf("[IDM] Error updating user by ID: %v", err)
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Message: "Error updating user",
				Status:  http.StatusInternalServerError,
				Error:   "Internal Server Error",
			},
		}, err
	}

	if rowsAffected == 0 {
		// No rows were affected, which means the user was not found
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Message: "User not found or no changes made",
				Status:  http.StatusNotFound,
			},
		}, nil
	}

	// Return a success response
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
	userID := req.UserID.ID

	// Call the database method to delete the user by ID
	rowsAffected, err := s.DbConn.DeleteUserByID(int(userID))
	if err != nil {
		log.Printf("[IDM] Error deleting user by ID: %v", err)
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Message: "Error deleting user",
				Status:  http.StatusInternalServerError,
				Error:   "Internal Server Error",
			},
		}, err
	}

	if rowsAffected == 0 {
		// No rows were affected, which means the user was not found
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Message: "User not found",
				Status:  http.StatusNotFound,
			},
		}, nil
	}

	// Return a success response
	return &proto_files.EnhancedInfoResponse{
		RowsAffected: int64(rowsAffected),
		Info: &proto_files.Info{
			Message: "User deleted successfully",
			Status:  http.StatusOK,
		},
	}, nil
}

// GetUserRole implements the GetUserRole RPC method
func (s *MyIDMServer) GetUserRole(ctx context.Context, req *proto_files.UserIDRequest) (*proto_files.RoleResponse, error) {
	userID := req.UserID.ID

	// Call the database method to retrieve the user's role by ID
	userRole, err := s.DbConn.GetUserRoleByUserID(int(userID))
	if err != nil {
		log.Printf("[IDM] Error getting user's role: %v", err)
		return &proto_files.RoleResponse{
			Info: &proto_files.Info{
				Error:  "Error getting user's role",
				Status: http.StatusInternalServerError,
			},
		}, nil
	}

	if userRole == "" {
		// Role not found
		return &proto_files.RoleResponse{
			Info: &proto_files.Info{
				Message: "User role not assigned",
				Status:  http.StatusNotFound,
			},
		}, nil
	}

	// Return the user's role in the response
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
	userID := req.UserID.ID
	newRole := req.Role

	// Call the database method to change the user's role
	rowsAffected, err := s.DbConn.UpdateUserRoleByUserID(int(userID), newRole)
	if err != nil {
		log.Printf("[IDM] Error changing user's role: %v", err)
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Error:  fmt.Sprintf("Error changing user's role: %v", err),
				Status: http.StatusInternalServerError,
			},
		}, err
	}

	if rowsAffected == 0 {
		// No rows were affected, which means the user was not found
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Message: "User not found or no changes made",
				Status:  http.StatusNotFound,
			},
		}, nil
	}

	// Return a success response
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
	username := req.Username
	password, err := s.DbConn.GetUserPasswordByUsername(username)
	if err != nil {
		log.Printf("[IDM] Error getting user's password: %v", err)
		return &proto_files.PasswordResponse{
			Info: &proto_files.Info{
				Error:  fmt.Sprintf("Error getting user's password: %v", err),
				Status: http.StatusInternalServerError,
			},
		}, err
	}

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
	userID := req.UserID.ID
	newPassword := req.Password

	// Hash the new password before updating it
	hashedPassword, err := utils.HashPassword(newPassword)
	if err != nil {
		log.Printf("[IDM] Error hashing password: %v", err)
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Error:  fmt.Sprintf("Error hashing password: %v", err),
				Status: http.StatusInternalServerError,
			},
		}, err
	}

	// Call the database method to change the user's password
	rowsAffected, err := s.DbConn.UpdateUserPasswordByUserID(int(userID), hashedPassword)
	if err != nil {
		log.Printf("[IDM] Error changing user's password: %v", err)
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Error:  fmt.Sprintf("Error changing user's password: %v", err),
				Status: http.StatusInternalServerError,
			},
		}, err
	}

	if rowsAffected == 0 {
		// No rows were affected, which means the user was not found
		return &proto_files.EnhancedInfoResponse{
			Info: &proto_files.Info{
				Message: "User not found or no changes made",
				Status:  http.StatusNotFound,
			},
		}, nil
	}

	// Return a success response
	return &proto_files.EnhancedInfoResponse{
		RowsAffected: int64(rowsAffected),
		Info: &proto_files.Info{
			Message: "User password updated successfully",
			Status:  http.StatusOK,
		},
	}, nil
}
