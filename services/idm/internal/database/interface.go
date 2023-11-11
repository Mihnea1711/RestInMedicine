package database

import "github.com/mihnea1711/POS_Project/services/idm/internal/models"

type Database interface {
	AddUser(newUser models.UserRegistration) (int, error)

	GetAllUsers() ([]models.User, error)
	GetUserByID(userID int) (models.User, error)
	GetUserByUsername(username string) (models.User, error)

	UpdateUserByID(userCredentials models.CredentialsRequest, userId int) (int, error)

	DeleteUserByID(userID int) (int, error)

	GetUserPasswordByUsername(username string) (string, error)

	GetUserRoleByUserID(userID int) (string, error)
	GetUserRoleByUsername(username string) (string, error)
	UpdateUserRoleByUserID(userID int, newRole string) (int, error)

	UpdateUserPasswordByUserID(userID int, newPassword string) (int, error)
	UpdateUserTokenByID(userID int, newToken string) (int, error)

	Close() error
}
