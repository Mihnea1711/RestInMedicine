package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/programari/internal/models"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

// Delete a appointment by ID
func (aController *AppointmentController) DeleteAppointmentByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[APPOINTMENT] Attempting to delete a appointment by ID.")
	vars := mux.Vars(r)
	appointmentID, err := strconv.Atoi(vars[utils.DELETE_APPOINTMENT_BY_ID_PARAMETER])
	if err != nil {
		errMsg := fmt.Sprintf("Invalid appointment ID: %d", appointmentID)
		log.Printf("[APPOINTMENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Invalid appointment delete request"})
		return
	}

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Use aController.DbConn to delete the appointment by ID from the database
	rowsAffected, err := aController.DbConn.DeleteAppointmentByID(ctx, appointmentID)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Error deleting appointment: %s", err.Error())
			log.Printf("[APPOINTMENT] %s", errMsg)
			// Create a conflict response using ResponseData
			utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to delete appointment. Appointment not found"})
			return
		}

		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[APPOINTMENT] Failed to delete appointment by ID: %s\n", errMsg)
		response := models.ResponseData{Error: errMsg}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Check if the appointment exists and was deleted
	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("No appointment found with ID: %d", appointmentID)
		log.Println("[APPOINTMENT] " + errMsg)

		response := models.ResponseData{Error: "Appointment not found"}
		utils.RespondWithJSON(w, http.StatusNotFound, response)
		return
	}

	log.Printf("[APPOINTMENT] Successfully deleted appointment %d", appointmentID)
	// Use utils.RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: fmt.Sprintf("Appointment with ID: %d deleted successfully", appointmentID),
		Payload: models.RowsAffected{
			RowsAffected: rowsAffected,
		},
	})
}
