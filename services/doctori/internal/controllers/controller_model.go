package controllers

import "github.com/mihnea1711/POS_Project/services/doctori/internal/database"

type DoctorController struct {
	DbConn *database.MySQLDatabase
}
