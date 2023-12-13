package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/doctori/internal/database"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

type DoctorController struct {
	DbConn database.Database
}

func (pc *DoctorController) handleContextTimeout(ctx context.Context, w http.ResponseWriter) {
	select {
	case <-ctx.Done():
		errMsg := "Request canceled or timed out"
		log.Printf("[PATIENT] %s", errMsg)

		// Use RespondWithJSON for conflict response
		utils.RespondWithJSON(w, http.StatusRequestTimeout, models.ResponseData{
			Message: "Failed to create patient",
			Error:   errMsg,
		})
		return
	default:
		// No context timeout, do nothing
	}
}
