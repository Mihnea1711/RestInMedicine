package controllers

import "github.com/mihnea1711/POS_Project/services/doctori/internal/database/mysql"

type DoctorController struct {
	DbConn *mysql.MySQLDatabase
}
