package database

import (
	"context"
	"database/sql"

	"github.com/mihnea1711/POS_Project/services/idm/internal/models"
)

type Database interface {
	AddUser(ctx context.Context, newUser models.UserRegistration) (int, error)

	GetAllUsers(ctx context.Context, page, limit int) ([]models.User, error)
	GetUserByID(ctx context.Context, userID int) (*models.User, error)
	GetUserByUsername(ctx context.Context, username string) (*models.User, error)

	UpdateUserByID(ctx context.Context, userCredentials models.CredentialsRequest, userID int) (int, error)

	DeleteUserByID(ctx context.Context, userID int) (int, error)

	GetUserPasswordByUsername(ctx context.Context, username string) (string, error)

	GetUserRoleByUserID(ctx context.Context, userID int) (string, error)
	GetUserRoleByUsername(ctx context.Context, username string) (string, error)
	UpdateUserRoleByUserID(ctx context.Context, userID int, newRole string) (int, error)

	UpdateUserPasswordByUserID(ctx context.Context, userID int, newPassword string) (int, error)

	AddUserToTrash(ctx context.Context, userData models.TrashData) error
	GetDataFromTrashByUserID(ctx context.Context, userID int) (*models.TrashData, error)
	RemoveUserFromTrash(ctx context.Context, userID int) (int, error)

	GetDB() *sql.DB
	Close() error
}
