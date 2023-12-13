package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/database"
	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
)

type ConsultationController struct {
	DbConn database.Database
}

func (cc *ConsultationController) handleContextTimeout(ctx context.Context, w http.ResponseWriter) {
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
