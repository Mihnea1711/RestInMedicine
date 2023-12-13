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

	aController.handleContextTimeout(ctx, w)

	// Use aController.DbConn to save the appointment to the database
	lastInsertID, err := aController.DbConn.SaveAppointment(ctx, appointment)
	if err != nil {
		handleDatabaseCreateError(w, err)
		return
	}

	// Check if the appointment was created
	if lastInsertID == 0 {
		errorMsg := "Patient has not been saved to the database due to an unexpected error."
		log.Printf("[APPOINTMENT] %s", errorMsg)

		// Use RespondWithJSON for conflict response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
			Message: "Failed to create appointment",
			Error:   errorMsg,
		})
		return
	}

	// Construct the URI of the newly created patient
	baseURI := utils.FETCH_ALL_APPOINTMENTS_ENDPOINT
	newAppointmentURI := fmt.Sprintf("%s/%d", baseURI, lastInsertID)

	log.Printf("[APPOINTMENT] Successfully created appointment %d", lastInsertID)

	// the origin server SHOULD send a 201 (Created) response containing a Location header field that provides an identifier for the primary resource created
	w.Header().Set("Location", newAppointmentURI)
	// Use RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusCreated, models.ResponseData{
		Message: "Appointment created successfully",
		Payload: models.LastInsertedID{
			LastInsertedID: lastInsertID,
		},
	})
}

func handleDatabaseCreateError(w http.ResponseWriter, err error) {
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
}
