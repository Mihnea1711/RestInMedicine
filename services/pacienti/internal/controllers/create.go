package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (dController *PacientController) CreatePacient(w http.ResponseWriter, r *http.Request) {
	log.Printf("[PATIENT] Attempting to create a new PATIENT.")
	doctor := r.Context().Value(utils.DECODED_PATIENT).(*models.Pacient)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Use dc.DB to save the doctor to the database
	err := dController.DbConn.SavePacient(ctx, doctor)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[PATIENT] Failed to save doctor to the database: %s\n", errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
			Error: errMsg,
		})
		return
	}

	// Use RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: "Patient created",
	})
}
