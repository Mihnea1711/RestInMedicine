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

// Retrieve all appointments
func (aController *AppointmentController) GetAppointments(w http.ResponseWriter, r *http.Request) {
	log.Printf("[APPOINTMENT] Attempting to retrieve appointments.")

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	log.Printf("[APPOINTMENT] Fetching appointments with limit: %d, page: %d", limit, page)

	// Use aController.DbConn to fetch all appointments from the database
	appointments, err := aController.DbConn.FetchAppointments(ctx, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[APPOINTMENT] Failed to fetch appointments: %s\n", errMsg)

		// Respond with an error using the ResponseData struct
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

// GetAppointmentsByDoctorID retrieves appointments by Doctor ID
func (aController *AppointmentController) GetAppointmentsByDoctorID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[APPOINTMENT] Attempting to retrieve appointments by Doctor ID.")

	// Extract Doctor ID from request parameters
	vars := mux.Vars(r)
	doctorID, err := strconv.Atoi(vars[utils.FETCH_APPOINTMENTS_BY_DOCTOR_ID_PARAMETER])
	if err != nil {
		errMsg := "Invalid Doctor ID"
		log.Printf("[APPOINTMENT] %s: %s\n", errMsg, err)

		// Respond with an error using the ResponseData struct
		response := models.ResponseData{
			Message: "Bad Request",
			Error:   errMsg,
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	log.Printf("[APPOINTMENT] Fetching appointments by doctor ID %d with limit: %d, page: %d", doctorID, limit, page)

	// Use aController.DbConn to fetch appointments by Doctor ID from the database
	appointments, err := aController.DbConn.FetchAppointmentsByDoctorID(ctx, doctorID, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("Internal server error: %s", err)
		log.Printf("[APPOINTMENT] Failed to fetch appointments by Doctor ID: %s\n", errMsg)

		// Respond with an error using the ResponseData struct
		response := models.ResponseData{
			Message: "Failed to fetch appointments by doctor ID",
			Error:   errMsg,
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	log.Printf("[APPOINTMENT] Successfully fetched appointments of doctor %d", doctorID)

	// Respond with success using the ResponseData struct
	response := models.ResponseData{
		Message: fmt.Sprintf("Successfully fetched appointments for Doctor ID: %d", doctorID),
		Payload: appointments,
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// GetAppointmentsByPatientID retrieves appointments by Patient ID
func (aController *AppointmentController) GetAppointmentsByPatientID(w http.ResponseWriter, r *http.Request) {
	// Log: Start attempting to retrieve appointments by Patient ID
	log.Printf("[APPOINTMENT] Attempting to retrieve appointments by Patient ID.")

	// Extract Patient ID from request parameters
	vars := mux.Vars(r)
	patientID, err := strconv.Atoi(vars[utils.FETCH_APPOINTMENTS_BY_PACIENT_ID_PARAMETER])
	if err != nil {
		errMsg := "Invalid Patient ID"
		log.Printf("[APPOINTMENT] %s: %s\n", errMsg, err)

		// Respond with an error using the ResponseData struct
		response := models.ResponseData{
			Message: "Bad Request",
			Error:   errMsg,
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	log.Printf("[APPOINTMENT] Fetching appointments by patient ID %d with limit: %d, page: %d", patientID, limit, page)

	// Use aController.DbConn to fetch appointments by Patient ID from the database
	appointments, err := aController.DbConn.FetchAppointmentsByPatientID(ctx, patientID, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("Internal server error: %s", err)
		log.Printf("[APPOINTMENT] Failed to fetch appointments by Patient ID: %s\n", errMsg)

		// Respond with an error using the ResponseData struct
		response := models.ResponseData{
			Message: "Failed to fetch appointments by Patient ID",
			Error:   errMsg,
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Log: Successfully fetched appointments
	log.Printf("[APPOINTMENT] Successfully fetched appointments of patient %d", patientID)

	// Respond with success using the ResponseData struct
	response := models.ResponseData{
		Message: fmt.Sprintf("Successfully fetched appointments for Patient ID: %d", patientID),
		Payload: appointments,
	}

	// Serialize the appointments to JSON and send the response
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// GetAppointmentsByDate retrieves appointments by date
func (aController *AppointmentController) GetAppointmentsByDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("[APPOINTMENT] Attempting to retrieve appointments by date.")

	// Use mux.Vars to get the date parameter from the route
	vars := mux.Vars(r)
	dateStr := vars[utils.FETCH_APPOINTMENTS_BY_DATE_PARAMETER]

	// Parse the date string into a time.Time object
	date, err := time.Parse(utils.TIME_PARSE_SYNTAX, dateStr)
	if err != nil {
		errMsg := "Invalid date format"
		log.Printf("[APPOINTMENT] %s: %s\n", errMsg, err)

		// Respond with an error using the ResponseData struct
		response := models.ResponseData{
			Message: "Bad Request",
			Error:   errMsg,
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	log.Printf("[APPOINTMENT] Fetching appointments by date %s with limit: %d, page: %d", date, limit, page)

	// Use aController.DbConn to fetch appointments by date from the database
	appointments, err := aController.DbConn.FetchAppointmentsByDate(ctx, date, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("Internal server error: %s", err)
		log.Printf("[APPOINTMENT] Failed to fetch appointments by date: %s\n", errMsg)

		// Respond with an error using the ResponseData struct
		response := models.ResponseData{
			Message: "Failed to fetch appointments by date",
			Error:   errMsg,
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	log.Printf("[APPOINTMENT] Successfully fetched appointments from %s", date)

	// Respond with success using the ResponseData struct
	response := models.ResponseData{
		Message: fmt.Sprintf("Successfully fetched appointments for date: %s", date),
		Payload: appointments,
	}

	// Serialize the appointments to JSON and send the response
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// GetAppointmentsByStatus retrieves appointments by status
func (aController *AppointmentController) GetAppointmentsByStatus(w http.ResponseWriter, r *http.Request) {
	log.Printf("[APPOINTMENT] Attempting to retrieve appointments by status.")

	// Use mux.Vars to get the status parameter from the route
	vars := mux.Vars(r)
	status := vars[utils.FETCH_APPOINTMENTS_BY_STATUS_PARAMETER]

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	log.Printf("[APPOINTMENT] Fetching appointments by status %s with limit: %d, page: %d", status, limit, page)

	// Use aController.DbConn to fetch appointments by status from the database
	appointments, err := aController.DbConn.FetchAppointmentsByStatus(ctx, status, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("Internal server error: %s", err)
		log.Printf("[APPOINTMENT] Failed to fetch appointments by status: %s\n", errMsg)

		// Respond with an error using the ResponseData struct
		response := models.ResponseData{
			Message: "Failed to fetch appointments by status",
			Error:   errMsg,
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	log.Printf("[APPOINTMENT] Successfully fetched appointments with status of %s", status)

	// Respond with success using the ResponseData struct
	response := models.ResponseData{
		Message: fmt.Sprintf("Successfully fetched appointments with status: %s", status),
		Payload: appointments,
	}

	// Serialize the appointments to JSON and send the response
	utils.RespondWithJSON(w, http.StatusOK, response)
}
