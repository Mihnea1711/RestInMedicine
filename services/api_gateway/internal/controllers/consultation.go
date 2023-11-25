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
	decodedResponseDoctor, statusDoctor, errDoctor := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.DOCTOR_FETCH_DOCTOR_BY_ID_ENDPOINT, consultationRequest.IDDoctor), utils.DOCTOR_PORT, nil)
	if errDoctor != nil {
		log.Printf("[GATEWAY] Error redirecting doctor ID request: %v", errDoctor)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to validate doctor ID", errDoctor.Error())
		return
	}
	if statusDoctor != http.StatusOK {
		log.Printf("[GATEWAY] Doctor ID doesn't exist or an unexpected error occured with status: %d", statusDoctor)
		utils.SendErrorResponse(w, statusDoctor, decodedResponseDoctor.Message, decodedResponseDoctor.Error)
		return
	}

	// Check if appointmentRequest.IDPatient exists
	decodedResponsePatient, statusPatient, errPatient := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.PATIENT_FETCH_PATIENT_BY_ID_ENDPOINT, consultationRequest.IDPatient), utils.PATIENT_PORT, nil)
	if errPatient != nil {
		log.Printf("[GATEWAY] Error redirecting patient ID request: %v", errPatient)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to validate patient ID", errPatient.Error())
		return
	}
	if statusPatient != http.StatusOK {
		log.Printf("[GATEWAY] Patient ID doesn't exist or an unexpected error occured with status: %d", statusPatient)
		utils.SendErrorResponse(w, statusPatient, decodedResponsePatient.Message, decodedResponsePatient.Error)
		return
	}

	// Redirect the request body to programare module to create the programare
	decodedResponse, status, err := gc.redirectRequestBody(ctx, utils.POST, utils.CONSULTATION_CREATE_CONSULTATIE_ENDPOINT, utils.CONSULTATION_PORT, consultationRequest)
	if err != nil {
		log.Printf("[GATEWAY] Failed to redirect request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	// Check the gRPC response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] CreateConsultation: Request successful with status %d", status)
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

	// Redirect the request to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, utils.GET, utils.CONSULTATION_FETCH_ALL_CONSULTATII_ENDPOINT, utils.CONSULTATION_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting consultation request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	// Check the gRPC response status and handle accordingly
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

// GetConsultationsByDoctorID handles the retrieval of consultations by doctor ID.
func (gc *GatewayController) GetConsultationsByDoctorID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to get consultations by doctorID.")

	// Get DoctorID from request params
	doctorIDString := mux.Vars(r)[utils.GET_CONSULTATION_BY_DOCTOR_ID_PARAMETER]
	// Convert doctorIDString to int64
	doctorID, err := strconv.ParseInt(doctorIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid consultation doctor ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid consultation doctor ID", err.Error())
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Check if appointmentRequest.IDDoctor exists
	decodedResponseDoctor, statusDoctor, errDoctor := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.DOCTOR_FETCH_DOCTOR_BY_ID_ENDPOINT, doctorID), utils.DOCTOR_PORT, nil)
	if errDoctor != nil {
		log.Printf("[GATEWAY] Error redirecting doctor ID request: %v", errDoctor)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to validate doctor ID", errDoctor.Error())
		return
	}
	if statusDoctor != http.StatusOK {
		log.Printf("[GATEWAY] Doctor ID doesn't exist or an unexpected error occured with status: %d", statusDoctor)
		utils.SendErrorResponse(w, statusDoctor, decodedResponseDoctor.Message, decodedResponseDoctor.Error)
		return
	}

	// Redirect the request body to appointment module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.CONSULTATION_FETCH_CONSULTATIE_BY_DOCTOR_ID_ENDPOINT, doctorID), utils.CONSULTATION_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting consultation request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	// Check the gRPC response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] GetConsultationsByDoctorID: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] GetConsultationsByDoctorID: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Consultations not found: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] GetConsultationsByDoctorID: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetConsultationsByPatientID handles the retrieval of consultations by patient ID.
func (gc *GatewayController) GetConsultationsByPatientID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to get consultations by patientID.")

	// Get PatientID from request params
	patientIDString := mux.Vars(r)[utils.GET_CONSULTATION_BY_PATIENT_ID_PARAMETER]
	// Convert patientIDString to int64
	patientID, err := strconv.ParseInt(patientIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid consultation patient ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid consultation patient ID", err.Error())
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Check if appointmentRequest.IDPatient exists
	decodedResponsePatient, statusPatient, errPatient := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.PATIENT_FETCH_PATIENT_BY_ID_ENDPOINT, patientID), utils.PATIENT_PORT, nil)
	if errPatient != nil {
		log.Printf("[GATEWAY] Error redirecting patient ID request: %v", errPatient)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to validate patient ID", errPatient.Error())
		return
	}
	if statusPatient != http.StatusOK {
		log.Printf("[GATEWAY] Patient ID doesn't exist or an unexpected error occured with status: %d", statusPatient)
		utils.SendErrorResponse(w, statusPatient, decodedResponsePatient.Message, decodedResponsePatient.Error)
		return
	}

	// Redirect the request body to appointment module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.CONSULTATION_FETCH_CONSULTATIE_BY_PATIENT_ID_ENDPOINT, patientID), utils.CONSULTATION_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting consultation request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	// Check the gRPC response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] GetConsultationsByPatientID: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] GetConsultationsByPatientID: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Consultation PatientID not found: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] GetConsultationsByPatientID: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetConsultationsByDate handles the retrieval of consultations by date.
func (gc *GatewayController) GetConsultationsByDate(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to get appointments by date.")

	// Get date from request params
	dateStr := mux.Vars(r)[utils.GET_CONSULTATION_BY_DATE_PARAMETER]

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to appointment module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%s", utils.CONSULTATION_FETCH_CONSULTATIE_BY_DATE_ENDPOINT, dateStr), utils.CONSULTATION_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY]Error redirecting consultation request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	// Check the gRPC response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] GetConsultationsByDate: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] GetConsultationsByDate: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Consultations not found: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] GetConsultationsByDate: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetConsultationByID handles the retrieval of a consultation by ID.
func (gc *GatewayController) GetConsultationByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to get consultation by ID.")

	// Get PatientID from request params
	consultationIDString := mux.Vars(r)[utils.GET_CONSULTATION_BY_ID_PARAMETER]
	// Convert patientIDString to int64
	consultationID, err := strconv.ParseInt(consultationIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid consultation ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid consultation ID", err.Error())
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.CONSULTATION_FETCH_CONSULTATIE_BY_ID_ENDPOINT, consultationID), utils.CONSULTATION_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting consultation request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	// Check the gRPC response status and handle accordingly
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
	// Convert consultationIDString to int64
	consultationID, err := strconv.ParseInt(consultationIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid consultation ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid consultation ID", err.Error())
		return
	}

	// Take consultation data from the context after validation
	consultationData := r.Context().Value(utils.DECODED_CONSULTATION_DATA).(*models.ConsultationData)

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Check if appointmentRequest.IDDoctor exists
	decodedResponseDoctor, statusDoctor, errDoctor := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.DOCTOR_FETCH_DOCTOR_BY_ID_ENDPOINT, consultationData.IDDoctor), utils.DOCTOR_PORT, nil)
	if errDoctor != nil {
		log.Printf("[GATEWAY] Error redirecting doctor ID request: %v", errDoctor)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to validate doctor ID", errDoctor.Error())
		return
	}
	if statusDoctor != http.StatusOK {
		log.Printf("[GATEWAY] Doctor ID doesn't exist or an unexpected error occured with status: %d", statusDoctor)
		utils.SendErrorResponse(w, statusDoctor, decodedResponseDoctor.Message, decodedResponseDoctor.Error)
		return
	}

	// Check if appointmentRequest.IDPatient exists
	decodedResponsePatient, statusPatient, errPatient := gc.redirectRequestBody(ctx, http.MethodGet, fmt.Sprintf("%s/%d", utils.PATIENT_FETCH_PATIENT_BY_ID_ENDPOINT, consultationData.IDPatient), utils.PATIENT_PORT, nil)
	if errPatient != nil {
		log.Printf("[GATEWAY] Error redirecting patient ID request: %v", errPatient)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to validate patient ID", errPatient.Error())
		return
	}
	if statusPatient != http.StatusOK {
		log.Printf("[GATEWAY] Patient ID doesn't exist or an unexpected error occured with status: %d", statusPatient)
		utils.SendErrorResponse(w, statusPatient, decodedResponsePatient.Message, decodedResponsePatient.Error)
		return
	}

	// Redirect the request body to programare module to create the consultation
	decodedResponse, status, err := gc.redirectRequestBody(ctx, utils.PUT, fmt.Sprintf("%s/%d", utils.CONSULTATION_UPDATE_CONSULTATIE_BY_ID_ENDPOINT, consultationID), utils.CONSULTATION_PORT, consultationData)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting consultation request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	// Check the gRPC response status and handle accordingly
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
	// Convert consultationIDString to int64
	consultationID, err := strconv.ParseInt(consultationIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid consultation ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid consultation ID", err.Error())
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, utils.DELETE, fmt.Sprintf("%s/%d", utils.CONSULTATION_DELETE_CONSULTATIE_BY_ID_ENDPOINT, consultationID), utils.CONSULTATION_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting consultation request: %v", err)
		utils.SendErrorResponse(w, http.StatusInternalServerError, "Failed to redirect request", err.Error())
		return
	}

	// Check the gRPC response status and handle accordingly
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
