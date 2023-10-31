package controllers

import "github.com/mihnea1711/POS_Project/services/pacienti/internal/database"

type PacientController struct {
	DbConn database.Database
}
