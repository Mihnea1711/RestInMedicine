package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (pController *PatientController) TogglePatientActivity(w http.ResponseWriter, r *http.Request) {
	log.Println("[PATIENT] Setting patient activity...")

	// Decode the patient details from the context
	reqData := r.Context().Value(utils.DECODED_PATIENT_ACTIVITY).(*models.ActivityData)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Use pController.DbConn to update the patient in the database
	rowsAffected, err := pController.DbConn.SetPatientActivityByUserID(ctx, reqData.IsActive, reqData.IDUser)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Error setting patient activity: %s", err.Error())
			log.Printf("[PATIENT] %s", errMsg)
			// Create a conflict response using ResponseData
			utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to update patient. Patient not found"})
			return
		}

		errMsg := fmt.Sprintf("Internal server error: %s", err)
		log.Printf("[PATIENT] Failed to update patient in the database: %s\n", errMsg)

		// Create an error response using ResponseData
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Internal database server error"})
		return
	}

	// Check if any rows were updated
	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("Patient with user ID: %d not modified", reqData.IDUser)
		log.Printf("[PATIENT] %s", errMsg)

		// Create an error response using ResponseData
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Patient data did not change."})
		return
	}

	log.Printf("[PATIENT] Successfully updated patient with user ID %d", reqData.IDUser)
	// Create a success response using ResponseData
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: fmt.Sprintf("Patient with user ID %d updated successfully", reqData.IDUser),
		Payload: models.RowsAffected{
			RowsAffected: rowsAffected,
		},
	})
}
