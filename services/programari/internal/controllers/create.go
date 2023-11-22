package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

func (aController *AppointmentController) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	log.Printf("[APPOINTMENT] Attempting to create a new appointment.")
	// Decode the appointment details from the context
	appointment := r.Context().Value(utils.DECODED_APPOINTMENT).(*models.Appointment)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Use aController.DbConn to save the appointment to the database
	lastInsertID, err := aController.DbConn.SaveAppointment(ctx, appointment)
	if err != nil {
		// Check if the error is due to a duplicate entry violation
		mysqlErr, ok := err.(*mysql.MySQLError)
		if ok && mysqlErr.Number == utils.MySQLDuplicateEntryErrorCode {
			errMsg := fmt.Sprintf("Conflict error: %s", mysqlErr.Message)
			log.Printf("[APPOINTMENT] %s", errMsg)

			// Create a conflict response using ResponseData
			utils.RespondWithJSON(w, http.StatusConflict, models.ResponseData{
				Error:   errMsg,
				Message: "Failed to create appointment. Duplicate entry violation"})
			return
		}

		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[APPOINTMENT] Failed to save appointment to the database: %s\n", errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
			Message: "Failed to create appointment",
			Error:   errMsg,
		})
		return
	}

	// Check if the appointment was created
	if lastInsertID == 0 {
		errorMsg := "Patient has not been saved to the database due to an unexpected error."
		log.Printf("[PATIENT] %s", errorMsg)

		// Use RespondWithJSON for conflict response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
			Message: "Failed to create appointment",
			Error:   errorMsg,
		})
		return
	}

	log.Printf("[APPOINTMENT] Successfully created appointment %d", lastInsertID)
	// Use RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: "Appointment created successfully",
		Payload: models.LastInsertedID{
			LastInsertedID: lastInsertID,
		},
	})
}
