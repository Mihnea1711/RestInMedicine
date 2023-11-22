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

func (pController *PatientController) CreatePatient(w http.ResponseWriter, r *http.Request) {
	log.Printf("[PATIENT] Attempting to create a new PATIENT.")
	patient := r.Context().Value(utils.DECODED_PATIENT).(*models.Pacient)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Use dController.DbConn to save the patient to the database
	lastInsertID, err := pController.DbConn.SavePatient(ctx, patient)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[PATIENT] Failed to save patient to the database: %s\n", errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
			Message: "Failed to create patient",
			Error:   errMsg,
		})
		return
	}

	if lastInsertID == 0 {
		errorMsg := "Patient has not been saved to the database."
		log.Printf("[PATIENT] %s", errorMsg)

		// Use RespondWithJSON for conflict response
		utils.RespondWithJSON(w, http.StatusConflict, models.ResponseData{
			Message: "Failed to create patient",
			Error:   errorMsg,
		})
		return
	}

	log.Printf("[PATIENT] Successfully created patient %d", lastInsertID)
	// Use RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: fmt.Sprintf("Patient created with ID %d", lastInsertID),
		Payload: lastInsertID,
	})
}
