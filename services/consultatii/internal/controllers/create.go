package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
)

func (cController *ConsultatieController) CreateConsultatie(w http.ResponseWriter, r *http.Request) {
	log.Printf("[CONSULTATION] Attempting to create a new consultatie.")
	consultatie := r.Context().Value(utils.DECODED_CONSULTATION).(*models.Consultatie)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	existingConsultatie, err := cController.DbConn.FetchConsultatieByID(ctx, consultatie.IDConsultatie)
	if err == nil && existingConsultatie != nil {
		utils.RespondWithJSON(w, http.StatusConflict, "Consultatie with the same ID already exists")
		return
	}

	// Use cController.DbConn to save the programare to the database
	err = cController.DbConn.SaveConsultatie(ctx, consultatie)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[CONSULTATION] Failed to save consultatie to the database: %s\n", errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, errMsg)
		return
	}

	log.Printf("[CONSULTATION] Successfully created consultatie %d", consultatie.IDConsultatie)
	utils.RespondWithJSON(w, http.StatusOK, "Consultatie created")
}
