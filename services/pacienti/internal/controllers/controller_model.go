package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/database"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

type PatientController struct {
	DbConn database.Database
}

func (pc *PatientController) handleContextTimeout(ctx context.Context, w http.ResponseWriter) {
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
