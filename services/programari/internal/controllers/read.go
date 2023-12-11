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

// GetAppointmentsByFilter retrieves appointments by filter.
func (aController *AppointmentController) GetAppointments(w http.ResponseWriter, r *http.Request) {
	log.Printf("[APPOINTMENT] Attempting to retrieve appointments.")

	// get query params for appointment
	filters, err := utils.ExtractFiltersFromRequest(r)
	if err != nil {
		errMsg := fmt.Sprintf("bad request: %s", err)
		log.Printf("[APPOINTMENT] GetAppointmentsByFilter: Failed to extract filters: %s", errMsg)

		// Respond with a bad request error
		response := models.ResponseData{
			Error:   errMsg,
			Message: "Failed to extract filters",
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	log.Printf("[APPOINTMENT] Fetching appointments with limit: %d, page: %d", limit, page)

	// Call the database method to get appointments
	appointments, err := aController.DbConn.FetchAppointments(ctx, filters, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[APPOINTMENT] GetAppointmentsByFilter: Failed to fetch appointments: %s", errMsg)

		// Respond with an internal server error
		response := models.ResponseData{
			Error:   errMsg,
			Message: "Failed to fetch appointments",
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	log.Printf("[APPOINTMENT] Successfully fetched %d appointments", len(appointments))

	// Respond with success using the ResponseData struct
	response := models.ResponseData{
		Payload: appointments,
		Message: fmt.Sprintf("Successfully fetched %d appointments", len(appointments)),
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// GetAppointmentByID retrieves an appointment by ID
func (aController *AppointmentController) GetAppointmentByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[APPOINTMENT] Attempting to retrieve an appointment by ID.")

	// Extract appointment ID from request parameters
	vars := mux.Vars(r)
	appointmentID, err := strconv.Atoi(vars[utils.FETCH_APPOINTMENT_BY_ID_PARAMETER])
	if err != nil {
		errMsg := "Invalid appointment ID"
		log.Printf("[APPOINTMENT] %s: %s\n", errMsg, err)

		// Respond with an error using the ResponseData struct
		response := models.ResponseData{
			Message: "Bad Request",
			Error:   errMsg,
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// Use aController.DbConn to fetch the appointment by ID from the database
	appointment, err := aController.DbConn.FetchAppointmentByID(ctx, appointmentID)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Error getting appointment by ID: %s", err.Error())
			log.Printf("[APPOINTMENT] %s", errMsg)
			// Create a conflict response using ResponseData
			utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to get appointment by ID. Appointment not found"})
			return
		}

		errMsg := fmt.Sprintf("Internal server error: %s", err)
		log.Printf("[APPOINTMENT] Failed to fetch appointment by ID: %s\n", errMsg)

		// Respond with an error using the ResponseData struct
		response := models.ResponseData{
			Message: "Failed to fetch appointment by id",
			Error:   errMsg,
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Check if the appointment exists
	if appointment == nil {
		errMsg := fmt.Sprintf("No appointment found with ID: %d", appointmentID)
		log.Printf("[APPOINTMENT] %s\n", errMsg)

		// Respond with an error using the ResponseData struct
		response := models.ResponseData{
			Message: "Appointment not found or an unexpected error happened.",
			Error:   errMsg,
		}
		utils.RespondWithJSON(w, http.StatusNotFound, response)
		return
	}

	log.Printf("[APPOINTMENT] Successfully fetched appointment with ID %d", appointment.IDProgramare)

	// Respond with success using the ResponseData struct
	response := models.ResponseData{
		Message: fmt.Sprintf("Successfully fetched appointment with ID: %d", appointment.IDProgramare),
		Payload: appointment,
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}
