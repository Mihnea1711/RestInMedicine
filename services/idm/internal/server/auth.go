package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/idm/idm/proto_files"
	"github.com/mihnea1711/POS_Project/services/idm/internal/models"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
)

// Register implements the Register RPC method
func (s *MyIDMServer) Register(ctx context.Context, req *proto_files.RegisterRequest) (*proto_files.InfoResponse, error) {
	if req == nil || req.UserCredentials == nil {
		return nil, errors.New("request or user credentials is nil")
	}

	userCredentials := models.UserRegistration{
		Username: req.UserCredentials.Username,
		Password: req.UserCredentials.Password,
		Role:     req.Role,
	}

	// Hash the user's password before adding to the database
	hashedPassword, err := utils.HashPassword(userCredentials.Password)
	if err != nil {
		log.Printf("[IDM] Error hashing password: %v", err)
		// Handle the error and return a meaningful InfoResponse
		return nil, fmt.Errorf("error hashing password. %v", err)
	}

	// Set the hashed password back to the user registration
	userCredentials.Password = hashedPassword

	// Call the database method to add the user to the database
	lastUserID, err := s.DbConn.AddUser(userCredentials)
	if err != nil {
		log.Printf("[IDM] Error adding user to the database: %v", err)
		// Handle the error and return a meaningful InfoResponse
		return nil, fmt.Errorf("error adding user to the database. %v", err)
	}

	if lastUserID == 0 {
		// No rows were affected, which means the user was not added
		// Handle the conflict situation and return a meaningful InfoResponse
		return &proto_files.InfoResponse{
			Info: &proto_files.Info{
				Message: "User not added",
				Status:  http.StatusConflict,
			},
		}, nil
	}

	// Return a successful response
	return &proto_files.InfoResponse{
		Info: &proto_files.Info{
			Message: "Registration successful. Proceed to login.",
			Status:  http.StatusOK,
		},
	}, nil
}

// Login implements the Login RPC method
func (s *MyIDMServer) Login(ctx context.Context, req *proto_files.LoginRequest) (*proto_files.LoginResponse, error) {
	userCredentials := models.UserRegistration{
		Username: req.UserCredentials.Username,
		Password: req.UserCredentials.Password,
	}

	// Retrieve the hashed password for the user from the database
	hashedPassword, err := s.DbConn.GetUserPasswordByUsername(userCredentials.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in the database: %s", userCredentials.Username)
			return nil, fmt.Errorf("user not found. %v", err)
		} else {
			log.Printf("[IDM] Error retrieving user's hashed password: %v", err)
			return nil, fmt.Errorf("error retrieving user's hashed password. %v", err)
		}
	}

	// Verify the password
	err = utils.VerifyPassword(hashedPassword, userCredentials.Password)
	if err != nil {
		log.Printf("[IDM] Invalid password for user: %s", userCredentials.Username)
		return nil, fmt.Errorf("invalid credentials. %v", err)
	}

	// Retrieve the user's role
	userRole, err := s.DbConn.GetUserRoleByUsername(userCredentials.Username)
	if err != nil {
		log.Printf("[IDM] Error retrieving user's role: %v", err)
		return nil, fmt.Errorf("error retrieving user's role. %v", err)
	}

	userComplete, err := s.DbConn.GetUserByUsername(userCredentials.Username)
	if err != nil {
		log.Printf("[IDM] Error retrieving user info: %v", err)
		return nil, fmt.Errorf("error retrieving user info. %v", err)
	}

	// Generate a JWT token
	token, err := utils.CreateJWT(userComplete.IDUser, userRole, s.JwtConfig)
	if err != nil {
		log.Printf("[IDM] Error generating JWT: %v", err)
		return nil, fmt.Errorf("error generating JWT. %v", err)
	}

	return &proto_files.LoginResponse{
		Token: token,
		Info: &proto_files.Info{
			Status:  http.StatusOK,
			Message: "Login successful",
		},
	}, nil
}
