package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (dController *DoctorController) ToggleDoctorActivity(w http.ResponseWriter, r *http.Request) {
	log.Println("[DOCTOR] Setting doctor activity...")

	// Decode the doctor details from the context
	reqData := r.Context().Value(utils.DECODED_DOCTOR_ACTIVITY).(*models.ActivityData)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Use dController.DbConn to update the doctor in the database
	rowsAffected, err := dController.DbConn.SetDoctorActivityByUserID(ctx, reqData.IsActive, reqData.IDUser)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Error setting doctor activity: %s", err.Error())
			log.Printf("[DOCTOR] %s", errMsg)
			// Create a conflict response using ResponseData
			utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to update doctor. Doctor not found"})
			return
		}

		errMsg := fmt.Sprintf("Internal server error: %s", err)
		log.Printf("[DOCTOR] Failed to update doctor in the database: %s\n", errMsg)

		// Create an error response using ResponseData
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Internal database server error"})
		return
	}

	// Check if any rows were updated
	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("Doctor with user ID: %d not modified", reqData.IDUser)
		log.Printf("[DOCTOR] %s", errMsg)

		// Create an error response using ResponseData
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Doctor data did not change."})
		return
	}

	log.Printf("[DOCTOR] Successfully updated doctor with user ID %d", reqData.IDUser)
	// Create a success response using ResponseData
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: fmt.Sprintf("Doctor with user ID %d updated successfully", reqData.IDUser),
		Payload: models.RowsAffected{
			RowsAffected: rowsAffected,
		},
	})
}
