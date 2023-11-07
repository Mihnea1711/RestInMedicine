package controllers

import "github.com/mihnea1711/POS_Project/services/idm/internal/database"

type IDMController struct {
	DbConn database.Database
}
