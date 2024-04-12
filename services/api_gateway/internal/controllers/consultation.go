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

// CreateConsultation handles the creation of a new consultation.
func (gc *GatewayController) CreateConsultation(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to create a consultation.")

	// Take consultation data from the context after validation
	consultationRequest := r.Context().Value(utils.DECODED_CONSULTATION_DATA).(*models.ConsultationData)

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Check if appointmentRequest.IDDoctor exists
	targetURL := fmt.Sprintf("%s/%d", utils.DOCTOR_FETCH_DOCTOR_BY_ID_ENDPOINT, consultationRequest.IDDoctor)
	decodedResponseDoctor, statusDoctor, errDoctor := gc.redirectRequestBody(ctx, http.MethodGet, utils.DOCTOR_HOST, targetURL, utils.DOCTOR_PORT, nil)
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
	targetURL = fmt.Sprintf("%s/%d", utils.PATIENT_FETCH_PATIENT_BY_ID_ENDPOINT, consultationRequest.IDPatient)
	decodedResponsePatient, statusPatient, errPatient := gc.redirectRequestBody(ctx, http.MethodGet, utils.PATIENT_HOST, targetURL, utils.PATIENT_PORT, nil)
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

	// Check if there is an appointment with the same IDPatient, IDDoctor and date already created.
	targetQuery := fmt.Sprintf(
		"%s=%d&%s=%d&%s=%s",
		utils.QUERY_ID_PATIENT, consultationRequest.IDPatient,
		utils.QUERY_ID_DOCTOR, consultationRequest.IDDoctor,
		utils.QUERY_DATE, consultationRequest.Date.Format(utils.TIME_PARSE))
	targetURL = fmt.Sprintf("%s?%s", utils.APPOINTMENT_FETCH_ALL_APPOINTMENTS_ENDPOINT, targetQuery)
	log.Println(targetURL)
	log.Println(targetQuery)
	decodedResponseAppointment, statusAppointment, errAppointment := gc.redirectRequestBody(ctx, http.MethodGet, utils.APPOINTMENT_HOST, targetURL, utils.APPOINTMENT_PORT, nil)
	log.Println("aici --------------------------------------------------------------------------")
	log.Println(decodedResponseAppointment, statusAppointment, errAppointment)
	if errAppointment != nil {
		log.Printf("[GATEWAY] Error redirecting appointment ID request: %v", errAppointment)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to validate appointment ID", errAppointment.Error())
		return
	}
	if statusAppointment != http.StatusOK || decodedResponseAppointment.Payload == nil {
		log.Printf("[GATEWAY] Appointment doesn't exist for this consultation or an unexpected error occured with status: %d", statusAppointment)
		utils.SendErrorResponse(w, http.StatusFailedDependency, decodedResponseAppointment.Message, decodedResponseAppointment.Error)
		return
	}

	// Redirect the request body to appointment module to create the appointment
	decodedResponse, status, err := gc.redirectRequestBody(ctx, utils.POST, utils.CONSULTATION_HOST, utils.CONSULTATION_CREATE_CONSULTATIE_ENDPOINT, utils.CONSULTATION_PORT, consultationRequest)
	if err != nil {
		log.Printf("[GATEWAY] Failed to redirect request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", err.Error())
		return
	}

	// Check the response status and handle accordingly
	switch status {
	case http.StatusCreated:
		log.Printf("[GATEWAY] CreateConsultation: Request successful with status %d", status)
		locationHeader := decodedResponse.Header.Get(utils.HEADER_LOCATION_KEY)
		w.Header().Set(utils.HEADER_LOCATION_KEY, fmt.Sprintf("/api%s", locationHeader))
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusConflict:
		log.Printf("[GATEWAY] CreateConsultation: Request failed with conflict status %d", status)
		utils.SendErrorResponse(w, http.StatusConflict, decodedResponse.Message, "Consultation Create Conflict: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] CreateConsultation: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetConsultations handles the retrieval of all consultations.
func (gc *GatewayController) GetConsultations(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to get all consultations.")

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Create target url based on jwt subject
	targetURL, err := gc.GenerateTargetURL(r, utils.CONSULTATION_FETCH_ALL_CONSULTATII_ENDPOINT)
	if err != nil {
		log.Printf("[GATEWAY] Error generating target URL: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Error generating target URL", err.Error())
		return
	}

	// Redirect the request to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, utils.GET, utils.CONSULTATION_HOST, targetURL, utils.CONSULTATION_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting consultation request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", err.Error())
		return
	}

	// Check the response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] GetConsultations: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	default:
		log.Printf("[GATEWAY] GetConsultations: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetConsultationByID handles the retrieval of a consultation by ID.
func (gc *GatewayController) GetConsultationByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to get consultation by ID.")

	// Get PatientID from request params
	consultationIDString := mux.Vars(r)[utils.GET_CONSULTATION_BY_ID_PARAMETER]

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, http.MethodGet, utils.CONSULTATION_HOST, fmt.Sprintf("%s/%s", utils.CONSULTATION_FETCH_CONSULTATIE_BY_ID_ENDPOINT, consultationIDString), utils.CONSULTATION_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting consultation request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", err.Error())
		return
	}

	// Check the response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] GetConsultationByID: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] GetConsultationByID: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Consultation not found: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] GetConsultationByID: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// UpdateConsultationByID handles the update of a specific consultation by ID.
func (gc *GatewayController) UpdateConsultationByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to update a consultation.")

	// Get consultationIDString from request params
	consultationIDString := mux.Vars(r)[utils.UPDATE_CONSULTATION_BY_ID_PARAMETER]

	// Take consultation data from the context after validation
	consultationData := r.Context().Value(utils.DECODED_CONSULTATION_DATA).(*models.ConsultationData)

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Check if appointmentRequest.IDDoctor exists
	decodedResponseDoctor, statusDoctor, errDoctor := gc.redirectRequestBody(ctx, http.MethodGet, utils.DOCTOR_HOST, fmt.Sprintf("%s/%d", utils.DOCTOR_FETCH_DOCTOR_BY_ID_ENDPOINT, consultationData.IDDoctor), utils.DOCTOR_PORT, nil)
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
	decodedResponsePatient, statusPatient, errPatient := gc.redirectRequestBody(ctx, http.MethodGet, utils.PATIENT_HOST, fmt.Sprintf("%s/%d", utils.PATIENT_FETCH_PATIENT_BY_ID_ENDPOINT, consultationData.IDPatient), utils.PATIENT_PORT, nil)
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

	// Redirect the request body to appointment module to create the consultation
	decodedResponse, status, err := gc.redirectRequestBody(ctx, utils.PUT, utils.CONSULTATION_HOST, fmt.Sprintf("%s/%s", utils.CONSULTATION_UPDATE_CONSULTATIE_BY_ID_ENDPOINT, consultationIDString), utils.CONSULTATION_PORT, consultationData)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting consultation request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", err.Error())
		return
	}

	// Check the response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] UpdateConsultationByID: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] UpdateConsultationByID: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Consultation not found: "+decodedResponse.Error)
		return
	case http.StatusConflict:
		log.Printf("[GATEWAY] UpdateConsultationByID: Request failed with conflict status %d", status)
		utils.SendErrorResponse(w, http.StatusConflict, decodedResponse.Message, "Consultation Update Conflict: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] UpdateConsultationByID: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// DeleteConsultationByID handles the deletion of a consultation by ID.
func (gc *GatewayController) DeleteConsultationByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to delete consultation by ID.")

	// Get consultationIDString from request params
	consultationIDString := mux.Vars(r)[utils.DELETE_CONSULTATION_BY_ID_PARAMETER]

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, utils.DELETE, utils.CONSULTATION_HOST, fmt.Sprintf("%s/%s", utils.CONSULTATION_DELETE_CONSULTATIE_BY_ID_ENDPOINT, consultationIDString), utils.CONSULTATION_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting consultation request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", err.Error())
		return
	}

	// Check the response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] DeleteConsultationByID: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] DeleteConsultationByID: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Appointment not found: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] DeleteConsultationByID: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}
