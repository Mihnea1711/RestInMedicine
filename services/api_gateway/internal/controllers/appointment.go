package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// CreateAppointment handles the creation of a new appointment.
func (gc *GatewayController) CreateAppointment(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to create an appointment.")

	// Take appointment data from the context after validation
	appointmentRequest := r.Context().Value(utils.DECODED_APPOINTMENT_DATA).(*models.AppointmentData)

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Check if appointmentRequest.IDDoctor exists
	decodedResponseDoctor, statusDoctor, errDoctor := gc.redirectRequestBody(ctx, http.MethodGet, utils.DOCTOR_HOST, fmt.Sprintf("%s/%d", utils.DOCTOR_FETCH_DOCTOR_BY_ID_ENDPOINT, appointmentRequest.IDDoctor), utils.DOCTOR_PORT, nil)
	if errDoctor != nil {
		log.Printf("[GATEWAY] Error redirecting doctor ID request: %v", errDoctor)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to validate doctor ID", errDoctor.Error())
		return
	}
	if statusDoctor != http.StatusOK {
		log.Printf("[GATEWAY] Doctor ID doesn't exist or an unexpected error occured with status: %d", statusDoctor)
		utils.SendErrorResponse(w, statusDoctor, decodedResponseDoctor.Message, decodedResponseDoctor.Error)
		return
	}

	// Check if appointmentRequest.IDPatient exists
	decodedResponsePatient, statusPatient, errPatient := gc.redirectRequestBody(ctx, http.MethodGet, utils.PATIENT_HOST, fmt.Sprintf("%s/%d", utils.PATIENT_FETCH_PATIENT_BY_ID_ENDPOINT, appointmentRequest.IDPatient), utils.PATIENT_PORT, nil)
	if errPatient != nil {
		log.Printf("[GATEWAY] Error redirecting patient ID request: %v", errPatient)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to validate patient ID", errPatient.Error())
		return
	}
	if statusPatient != http.StatusOK {
		log.Printf("[GATEWAY] Patient ID doesn't exist or an unexpected error occured with status: %d", statusPatient)
		utils.SendErrorResponse(w, statusPatient, decodedResponsePatient.Message, decodedResponsePatient.Error)
		return
	}

	// Redirect the request body to appointment module to create the appointment
	decodedResponse, status, err := gc.redirectRequestBody(ctx, utils.POST, utils.APPOINTMENT_HOST, utils.APPOINTMENT_CREATE_APPOINTMENT_ENDPOINT, utils.APPOINTMENT_PORT, appointmentRequest)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting appointment request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", err.Error())
		return
	}

	// Check the response status and handle accordingly
	switch status {
	case http.StatusCreated:
		log.Printf("[GATEWAY] CreateAppointment: Request successful with status %d", status)
		locationHeader := decodedResponse.Header.Get(utils.HEADER_LOCATION_KEY)
		w.Header().Set(utils.HEADER_LOCATION_KEY, fmt.Sprintf("/api%s", locationHeader))
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusConflict:
		log.Printf("[GATEWAY] CreateAppointment: Request failed with conflict status %d", status)
		utils.SendErrorResponse(w, http.StatusConflict, decodedResponse.Message, "Appointment Create Conflict: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] CreateAppointment: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetAppointments handles the retrieval of all appointments.
func (gc *GatewayController) GetAppointments(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to get all appointments.")

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Create target url based on jwt subject
	targetURL, err := gc.GenerateTargetURL(r, utils.APPOINTMENT_FETCH_ALL_APPOINTMENTS_ENDPOINT)
	if err != nil {
		log.Printf("[GATEWAY] Error generating target URL: %v", err)
		utils.SendErrorResponse(w, http.StatusUnauthorized, "Error generating target URL", err.Error())
		return
	}

	// Redirect the request to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, utils.GET, utils.APPOINTMENT_HOST, targetURL, utils.APPOINTMENT_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting appointment request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", err.Error())
		return
	}

	// Check the response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] GetAppointments: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	default:
		log.Printf("[GATEWAY] GetAppointments: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetAppointmentByID handles the retrieval of a appointment by ID.
func (gc *GatewayController) GetAppointmentByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to get appointment by ID.")

	// Get appointmentIDString from request params
	appointmentIDString := mux.Vars(r)[utils.GET_APPOINTMENT_ID_PARAMETER]
	// Convert appointmentIDString to int64
	appointmentID, err := strconv.ParseInt(appointmentIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid appointment ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid appointment ID", err.Error())
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, http.MethodGet, utils.APPOINTMENT_HOST, fmt.Sprintf("%s/%d", utils.APPOINTMENT_FETCH_APPOINTMENT_BY_ID_ENDPOINT, appointmentID), utils.APPOINTMENT_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting appointment request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", err.Error())
		return
	}

	// Check the response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] GetAppointmentByID: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] GetAppointmentByID: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Appointment not found: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] GetAppointmentByID: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// UpdateAppointmentByID handles the update of a specific appointment by ID.
func (gc *GatewayController) UpdateAppointmentByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to update a appointment.")

	// Get appointmentID from request params
	appointmentIDString := mux.Vars(r)[utils.GET_APPOINTMENT_ID_PARAMETER]
	// Convert appointmentIDString to int64
	appointmentID, err := strconv.ParseInt(appointmentIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid appointment ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid appointment ID", err.Error())
		return
	}

	// Take appointment data from the context after validation
	appointmentData := r.Context().Value(utils.DECODED_APPOINTMENT_DATA).(*models.AppointmentData)

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Check if appointmentRequest.IDDoctor exists
	decodedResponseDoctor, statusDoctor, errDoctor := gc.redirectRequestBody(ctx, http.MethodGet, utils.APPOINTMENT_HOST, fmt.Sprintf("%s/%d", utils.DOCTOR_FETCH_DOCTOR_BY_ID_ENDPOINT, appointmentData.IDDoctor), utils.DOCTOR_PORT, nil)
	if errDoctor != nil {
		log.Printf("[GATEWAY] Error redirecting doctor ID request: %v", errDoctor)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to validate doctor ID", errDoctor.Error())
		return
	}
	if statusDoctor != http.StatusOK {
		log.Printf("[GATEWAY] Doctor ID doesn't exist or an unexpected error occured with status: %d", statusDoctor)
		utils.SendErrorResponse(w, statusDoctor, decodedResponseDoctor.Message, decodedResponseDoctor.Error)
		return
	}

	// Check if appointmentRequest.IDPatient exists
	decodedResponsePatient, statusPatient, errPatient := gc.redirectRequestBody(ctx, http.MethodGet, utils.APPOINTMENT_HOST, fmt.Sprintf("%s/%d", utils.PATIENT_FETCH_PATIENT_BY_ID_ENDPOINT, appointmentData.IDPatient), utils.PATIENT_PORT, nil)
	if errPatient != nil {
		log.Printf("[GATEWAY] Error redirecting patient ID request: %v", errPatient)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to validate patient ID", errPatient.Error())
		return
	}
	if statusPatient != http.StatusOK {
		log.Printf("[GATEWAY] Patient ID doesn't exist or an unexpected error occured with status: %d", statusPatient)
		utils.SendErrorResponse(w, statusPatient, decodedResponsePatient.Message, decodedResponsePatient.Error)
		return
	}

	// Redirect the request body to appointment module to update the appointment
	decodedResponse, status, err := gc.redirectRequestBody(ctx, utils.PUT, utils.APPOINTMENT_HOST, fmt.Sprintf("%s/%d", utils.APPOINTMENT_UPDATE_APPOINTMENT_BY_ID_ENDPOINT, appointmentID), utils.APPOINTMENT_PORT, appointmentData)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting appointment request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", err.Error())
		return
	}

	// Check the response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] UpdateAppointmentByID: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] UpdateAppointmentByID: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Appointment not found: "+decodedResponse.Error)
		return
	case http.StatusConflict:
		log.Printf("[GATEWAY] UpdateAppointmentByID: Request failed with conflict status %d", status)
		utils.SendErrorResponse(w, http.StatusConflict, decodedResponse.Message, "Appointment Update Conflict: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] UpdateAppointmentByID: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// DeleteAppointmentByID handles the deletion of a appointment by ID.
func (gc *GatewayController) DeleteAppointmentByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to delete appointment by ID.")

	// Get appointmentID from request params
	appointmentIDString := mux.Vars(r)[utils.DELETE_APPOINTMENT_ID_PARAMETER]
	// Convert appointmentIDString to int64
	appointmentID, err := strconv.ParseInt(appointmentIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid appointment ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid appointment ID", err.Error())
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, utils.DELETE, utils.APPOINTMENT_HOST, fmt.Sprintf("%s/%d", utils.APPOINTMENT_DELETE_APPOINTMENT_BY_ID_ENDPOINT, appointmentID), utils.APPOINTMENT_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting appointment request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", err.Error())
		return
	}

	// Check the response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] DeleteAppointmentByID: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] DeleteAppointmentByID: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Appointment not found: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] DeleteAppointmentByID: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}
