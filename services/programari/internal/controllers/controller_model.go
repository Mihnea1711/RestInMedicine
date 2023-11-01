package controllers

import "github.com/mihnea1711/POS_Project/services/programari/internal/database"

type ProgramareController struct {
	DbConn database.Database
}
