package models

import "github.com/mihnea1711/POS_Project/services/idm/pkg/utils"

type User struct {
	ID       int            `db:"id" json:"id"`
	Username string         `db:"username" json:"username"`
	Password string         `db:"password" json:"-"`
	Role     utils.UserRole `db:"role" json:"role"`
}

type JWSToken struct {
	ID     int    `db:"id" json:"id"`
	UserID int    `db:"user_id" json:"userId"`
	Token  string `db:"token" json:"token"`
	Valid  bool   `db:"valid" json:"valid"`
}

type BlacklistToken struct {
	ID     int    `db:"id" json:"id"`
	UserID int    `db:"user_id" json:"userId"`
	Token  string `db:"token" json:"token"`
}
