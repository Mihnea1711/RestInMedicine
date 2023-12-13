package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

// Update a appointment by ID
func (aController *AppointmentController) UpdateAppointmentByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[APPOINTMENT] Attempting to update an appointment by ID.")

	// Decode the appointment details from the context
	appointment := r.Context().Value(utils.DECODED_APPOINTMENT).(*models.Appointment)

	// Get the appointment ID from the request path
	vars := mux.Vars(r)
	appointmentID, err := strconv.Atoi(vars[utils.UPDATE_APPOINTMENT_BY_ID_PARAMETER])
	if err != nil {
		errMsg := fmt.Sprintf("Invalid appointment ID: %d", appointmentID)
		log.Printf("[APPOINTMENT] %s", errMsg)
		// Create an error response using ResponseData
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Invalid appointment update request"})
		return
	}

	// Assign the ID to the appointment object
	appointment.IDProgramare = appointmentID

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Use aController.DbConn to update the appointment by ID in the database
	rowsAffected, err := aController.DbConn.UpdateAppointmentByID(ctx, appointment)
	if err != nil {
		handleDatabaseUpdateError(w, err)
		return
	}

	// Check if the appointment exists and was updated
	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("No appointment found with ID: %d", appointment.IDProgramare)
		log.Println("[APPOINTMENT] " + errMsg)

		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: "Appointment not found"})
		return
	}

	log.Printf("[APPOINTMENT] Successfully updated appointment %d", appointment.IDProgramare)
	// Create a success response using ResponseData
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: fmt.Sprintf("Appointment with ID %d updated successfully", appointmentID),
		Payload: models.RowsAffected{
			RowsAffected: rowsAffected,
		},
	})
}

func handleDatabaseUpdateError(w http.ResponseWriter, err error) {
	// Check if the error is due to no rows found
	if err == sql.ErrNoRows {
		errMsg := fmt.Sprintf("Error updating appointment: %s", err.Error())
		log.Printf("[APPOINTMENT] %s", errMsg)
		// Create a conflict response using ResponseData
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to update appointment. Appointment not found"})
		return
	}
	// Check if the error is a MySQL duplicate entry error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok && mysqlErr.Number == utils.MySQLDuplicateEntryErrorCode {
		errMsg := fmt.Sprintf("Conflict error: %s", mysqlErr.Message)
		log.Printf("[APPOINTMENT] %s", errMsg)

		// Create a conflict response using ResponseData
		utils.RespondWithJSON(w, http.StatusConflict, models.ResponseData{Error: errMsg, Message: "Failed to update appointment. Duplicate entry violation"})
		return
	}

	errMsg := fmt.Sprintf("internal server error: %s", err)
	log.Printf("[APPOINTMENT] Failed to update appointment in the database: %s\n", errMsg)

	// Create an error response using ResponseData
	utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Internal database server error"})
}
