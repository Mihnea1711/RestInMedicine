package server

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/mihnea1711/POS_Project/services/idm/idm/proto_files"
	"github.com/mihnea1711/POS_Project/services/idm/internal/models"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
)

// Register implements the Register RPC method
func (s *MyIDMServer) Register(ctx context.Context, req *proto_files.RegisterRequest) (*proto_files.InfoResponse, error) {
	if req == nil || req.UserCredentials == nil {
		log.Println("[IDM] Registration request or user credentials is nil")
		return nil, errors.New("request or user credentials is nil")
	}

	// Ensure a database operation doesn't take longer than 5 seconds
	childCtx, cancel := context.WithTimeout(ctx, utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	userCredentials := models.UserRegistration{
		Username: req.UserCredentials.Username,
		Password: req.UserCredentials.Password,
		Role:     req.Role,
	}

	// Hash the user's password before adding to the database
	hashedPassword, err := utils.HashPassword(userCredentials.Password)
	if err != nil {
		log.Printf("[IDM] Error hashing password: %v", err)
		return nil, fmt.Errorf("error hashing password. %v", err)
	}

	// Set the hashed password back to the user registration
	userCredentials.Password = hashedPassword

	// Call the database method to add the user to the database
	lastUserID, err := s.DbConn.AddUser(childCtx, userCredentials)
	if err != nil {
		// Check if the error is due to a duplicate entry violation
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok && mysqlErr.Number == utils.MySQLDuplicateEntryErrorCode {
			log.Printf("[IDM] Registration unsuccessful. Conflict occurred for username: %s", userCredentials.Username)
			// Return a conflict error status
			return &proto_files.InfoResponse{
				Info: &proto_files.Info{
					Message: fmt.Sprintf("Registration unsuccessful. Conflict occurred: %v", err),
					Status:  http.StatusConflict,
				},
			}, nil
		}

		log.Printf("[IDM] Error adding user to the database: %v", err)
		return nil, fmt.Errorf("error adding user to the database. %v", err)
	}

	if lastUserID == 0 {
		log.Printf("[IDM] User not added to the db for username: %s", userCredentials.Username)
		// No rows were affected, which means the user was not added
		return &proto_files.InfoResponse{
			Info: &proto_files.Info{
				Message: "Failed to register user",
				Status:  http.StatusInternalServerError,
			},
		}, nil
	}

	log.Printf("[IDM] Registration successful for username: %s. Proceed to login.", userCredentials.Username)
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

	// Ensure a database operation doesn't take longer than 5 seconds
	childCtx, cancel := context.WithTimeout(ctx, utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Start a database transaction
	tx, err := s.DbConn.GetDB().BeginTx(childCtx, &sql.TxOptions{Isolation: sql.LevelSerializable})
	if err != nil {
		log.Printf("[IDM] Error starting database transaction: %v", err)
		return nil, fmt.Errorf("error starting database transaction. %v", err)
	}
	defer tx.Rollback()

	// Retrieve the hashed password for the user from the database
	hashedPassword, err := s.DbConn.GetUserPasswordByUsername(childCtx, userCredentials.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in the database: %s", userCredentials.Username)
			// Roll back the transaction
			tx.Rollback()
			return &proto_files.LoginResponse{
				Token: "",
				Info: &proto_files.Info{
					Status:  http.StatusNotFound,
					Message: "Username or Password are incorrect",
				},
			}, nil
		} else {
			log.Printf("[IDM] Error retrieving user's hashed password: %v", err)
			// Roll back the transaction
			tx.Rollback()
			return nil, fmt.Errorf("error retrieving user's hashed password. %v", err)
		}
	}

	// Verify the password
	err = utils.VerifyPassword(hashedPassword, userCredentials.Password)
	if err != nil {
		log.Printf("[IDM] Invalid password for user: %s", userCredentials.Username)
		// Roll back the transaction
		tx.Rollback()
		return &proto_files.LoginResponse{
			Token: "",
			Info: &proto_files.Info{
				Status:  http.StatusNotFound,
				Message: "Username or Password are incorrect",
			},
		}, nil
	}

	// Retrieve the user's role
	userRole, err := s.DbConn.GetUserRoleByUsername(childCtx, userCredentials.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in the database: %s", userCredentials.Username)
			// Roll back the transaction
			tx.Rollback()
			return &proto_files.LoginResponse{
				Token: "",
				Info: &proto_files.Info{
					Status:  http.StatusNotFound,
					Message: "Username does not exist",
				},
			}, nil
		} else {
			log.Printf("[IDM] Error retrieving user's role: %v", err)
			// Roll back the transaction
			tx.Rollback()
			return nil, fmt.Errorf("error retrieving user's role. %v", err)
		}
	}

	userComplete, err := s.DbConn.GetUserByUsername(childCtx, userCredentials.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("[IDM] User not found in the database: %s", userCredentials.Username)
			// Roll back the transaction
			tx.Rollback()
			return &proto_files.LoginResponse{
				Token: "",
				Info: &proto_files.Info{
					Status:  http.StatusNotFound,
					Message: "Username or Password are incorrect",
				},
			}, nil
		} else {
			log.Printf("[IDM] Error retrieving user info: %v", err)
			// Roll back the transaction
			tx.Rollback()
			return nil, fmt.Errorf("error retrieving user info. %v", err)
		}
	}

	// Generate a JWT token
	token, err := utils.CreateJWT(userComplete.IDUser, userRole, s.JwtConfig)
	if err != nil || token == "" {
		tx.Rollback()
		log.Printf("[IDM] Error generating JWT: %v", err)
		return nil, fmt.Errorf("error generating JWT. %v", err)
	}

	// Commit the transaction if everything is successful
	if err := tx.Commit(); err != nil {
		log.Printf("[IDM] Error committing database transaction: %v", err)
		return nil, fmt.Errorf("error committing database transaction. %v", err)
	}

	log.Printf("[IDM] Login successful for user: %s", userCredentials.Username)
	return &proto_files.LoginResponse{
		Token: token,
		Info: &proto_files.Info{
			Status:  http.StatusOK,
			Message: "Login successful",
		},
	}, nil
}
