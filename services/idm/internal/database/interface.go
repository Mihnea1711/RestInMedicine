package database

import "github.com/mihnea1711/POS_Project/services/idm/internal/models"

type Database interface {
	AddUserToDB(newUser models.User) error

	GetAllUsersFromDB() ([]models.User, error)
	GetUserFromDBByID(userID int) (models.User, error)
	GetUserFromDBByUsername(username string) (models.User, error)

	UpdateUserInDB(updatedUser models.User) error

	DeleteUserFromDBByID(userID int) error

	GetUserPasswordByUsername(username string) (string, error)
	GetUserRoleByUserID(userID int) (string, error)
	GetUserRoleByUsername(username string) (string, error)

	ChangeUserRoleByUserID(userID int, newRole string) error
	ChangeUserRoleByUsername(username string, newRole string) error

	ChangeUserPasswordByUserID(userID int, newPassword string) error
	ChangeUserPasswordByUsername(username string, newPassword string) error

	Close() error
}
