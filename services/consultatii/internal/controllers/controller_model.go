package controllers

import "github.com/mihnea1711/POS_Project/services/consultatii/internal/database"

type ConsultationController struct {
	DbConn database.Database
}
