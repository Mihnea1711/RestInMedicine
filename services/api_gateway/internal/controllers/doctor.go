package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/idm/proto_files"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils/wrappers"
)

// CreateDoctor handles the creation of a new doctor.
func (gc *GatewayController) CreateDoctor(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to create a doctor.")

	// Take credentials data from the context after validation
	doctorRequest := r.Context().Value(utils.DECODED_DOCTOR_DATA).(*models.DoctorData)

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Check if IDUser exists
	userResponse, err := gc.IDMClient.GetUserByID(ctx, &proto_files.UserIDRequest{UserID: &proto_files.UserID{ID: int64(doctorRequest.IDUser)}})
	if err != nil {
		log.Printf("[GATEWAY] Error fetching user by ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Internal Server Error", "Failed to fetch user by ID: "+err.Error())
		return
	}
	// Check response for nils
	userResponseWrapper := &wrappers.UserResponse{Response: userResponse}
	utils.CheckNilResponse(w, http.StatusInternalServerError, "Get User By ID response is nil", userResponseWrapper.IsResponseNil, "Received nil response while getting the user.")
	utils.CheckNilResponse(w, http.StatusInternalServerError, "Get User By ID response info is nil", userResponseWrapper.IsInfoNil, "Received nil response.Info while getting user by id.")
	utils.CheckNilResponse(w, http.StatusInternalServerError, userResponse.Info.Message, userResponseWrapper.IsUserNil, "Received nil response.User while getting the user.")

	// Check if there is already a patient associated with the user ID
	targetURL := fmt.Sprintf("%s/%d", utils.PATIENT_FETCH_PATIENT_BY_USER_ID_ENDPOINT, doctorRequest.IDUser)
	_, statusPatient, err := gc.redirectRequestBody(ctx, utils.GET, utils.PATIENT_HOST, targetURL, utils.PATIENT_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting doctor request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", "Failed to redirect request: "+err.Error())
		return
	}
	if statusPatient == http.StatusOK {
		log.Println("[GATEWAY] Error creating a new doctor. There is already a patient associated with this UserID")
		utils.SendErrorResponse(w, http.StatusConflict, "There is already a patient associated with this UserID", "Error creating a new doctor")
		return
	}

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, utils.POST, utils.DOCTOR_HOST, utils.DOCTOR_CREATE_DOCTOR_ENDPOINT, utils.DOCTOR_PORT, doctorRequest)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting doctor request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", err.Error())
		return
	}

	// Check the response status and handle accordingly
	switch status {
	case http.StatusCreated:
		log.Printf("[GATEWAY] CreateDoctor: Request successful with status %d", status)
		locationHeader := decodedResponse.Header.Get(utils.HEADER_LOCATION_KEY)
		w.Header().Set(utils.HEADER_LOCATION_KEY, fmt.Sprintf("/api%s", locationHeader))
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusConflict:
		log.Printf("[GATEWAY] CreateDoctor: Request failed with conflict status %d", status)
		utils.SendErrorResponse(w, http.StatusConflict, decodedResponse.Message, "Doctor Create Conflict: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] CreateDoctor: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetDoctors handles the retrieval of all doctors.
func (gc *GatewayController) GetDoctors(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to get all doctors.")

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request to another module
	targetURL := fmt.Sprintf("%s?%s", utils.DOCTOR_FETCH_ALL_DOCTORS_ENDPOINT, r.URL.RawQuery)
	decodedResponse, status, err := gc.redirectRequestBody(ctx, utils.GET, utils.DOCTOR_HOST, targetURL, utils.DOCTOR_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting doctor request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", err.Error())
		return
	}

	// Check the response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] GetDoctors: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	default:
		log.Printf("[GATEWAY] GetDoctors: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetDoctorByID handles the retrieval of a doctor by ID.
func (gc *GatewayController) GetDoctorByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to get doctor by ID.")

	// Get DoctorID from request params
	doctorIDString := mux.Vars(r)[utils.GET_DOCTOR_BY_ID_PARAMETER]
	// Convert doctorIDString to int64
	doctorID, err := strconv.ParseInt(doctorIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid doctor ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid doctor ID", err.Error())
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, http.MethodGet, utils.DOCTOR_HOST, fmt.Sprintf("%s/%d", utils.DOCTOR_FETCH_DOCTOR_BY_ID_ENDPOINT, doctorID), utils.DOCTOR_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting doctor request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", err.Error())
		return
	}

	// Check the response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] GetDoctorByID: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] GetDoctorByID: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Doctor not found: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] GetDoctorByID: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetDoctorByEmail handles the retrieval of a doctor by email.
func (gc *GatewayController) GetDoctorByEmail(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to get doctor by email.")

	// Get DoctorEmail from request params
	doctorEmail := mux.Vars(r)[utils.GET_DOCTOR_BY_EMAIL_PARAMETER]

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, http.MethodGet, utils.DOCTOR_HOST, fmt.Sprintf("%s/%s", utils.DOCTOR_FETCH_DOCTOR_BY_EMAIL_ENDPOINT, doctorEmail), utils.DOCTOR_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting doctor request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", err.Error())
		return
	}

	// Check the response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] GetDoctorByEmail: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] GetDoctorByEmail: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Doctor not found: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] GetDoctorByEmail: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// GetDoctorByUserID handles the retrieval of a doctor by user ID.
func (gc *GatewayController) GetDoctorByUserID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to get doctor by UserID.")

	// Get UserID from request params
	userIDString := mux.Vars(r)[utils.GET_DOCTOR_BY_USER_ID_PARAMETER]
	// Convert userIDString to int64
	userID, err := strconv.ParseInt(userIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid user ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid doctor user ID", err.Error())
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, http.MethodGet, utils.DOCTOR_HOST, fmt.Sprintf("%s/%d", utils.DOCTOR_FETCH_DOCTOR_BY_USER_ID_ENDPOINT, userID), utils.DOCTOR_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting doctor request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", err.Error())
		return
	}

	// Check the response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] GetDoctorByUserID: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] GetDoctorByUserID: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Doctor not found: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] GetDoctorByUserID: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// UpdateDoctorByID handles the update of a specific doctor by ID.
func (gc *GatewayController) UpdateDoctorByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to update a doctor.")

	// Get DoctorID from request params
	doctorIDString := mux.Vars(r)[utils.UPDATE_DOCTOR_BY_ID_PARAMETER]
	// Convert doctorIDString to int64
	doctorID, err := strconv.ParseInt(doctorIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid doctor ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid doctor ID", err.Error())
		return
	}

	// Take credentials data from the context after validation
	doctorData := r.Context().Value(utils.DECODED_DOCTOR_DATA).(*models.DoctorData)

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, http.MethodPut, utils.DOCTOR_HOST, fmt.Sprintf("%s/%d", utils.DOCTOR_UPDATE_DOCTOR_BY_ID_ENDPOINT, doctorID), utils.DOCTOR_PORT, doctorData)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting doctor request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", err.Error())
		return
	}

	// Check the response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] UpdateDoctorByID: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] UpdateDoctorByID: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Doctor not found: "+decodedResponse.Error)
		return
	case http.StatusConflict:
		log.Printf("[GATEWAY] UpdateDoctorByID: Request failed with conflict status %d", status)
		utils.SendErrorResponse(w, http.StatusConflict, decodedResponse.Message, "Doctor Update Conflict: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] UpdateDoctorByID: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}

// DeleteDoctorByID handles the deletion of a doctor by ID.
func (gc *GatewayController) DeleteDoctorByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[GATEWAY] Attempting to delete doctor by ID.")

	// Get DoctorID from request params
	doctorIDString := mux.Vars(r)[utils.DELETE_DOCTOR_BY_ID_PARAMETER]
	// Convert doctorIDString to int64
	doctorID, err := strconv.ParseInt(doctorIDString, 10, 64)
	if err != nil {
		log.Printf("[GATEWAY] Invalid doctor ID: %v", err)
		utils.SendErrorResponse(w, http.StatusBadRequest, "Invalid doctor ID", err.Error())
		return
	}

	// Create a context with a timeout (adjust the timeout as needed)
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_CONTEXT_TIMEOUT*time.Second)
	defer cancel()

	// Redirect the request body to another module
	decodedResponse, status, err := gc.redirectRequestBody(ctx, http.MethodDelete, utils.DOCTOR_HOST, fmt.Sprintf("%s/%d", utils.DOCTOR_DELETE_DOCTOR_BY_ID_ENDPOINT, doctorID), utils.DOCTOR_PORT, nil)
	if err != nil {
		log.Printf("[GATEWAY] Error redirecting doctor request: %v", err)
		utils.SendErrorResponse(w, http.StatusBadGateway, "Failed to redirect request", err.Error())
		return
	}

	// Check the response status and handle accordingly
	switch status {
	case http.StatusOK:
		log.Printf("[GATEWAY] DeleteDoctorByID: Request successful with status %d", status)
		utils.SendMessageResponse(w, http.StatusOK, decodedResponse.Message, decodedResponse.Payload)
		return
	case http.StatusNotFound:
		log.Printf("[GATEWAY] DeleteDoctorByID: Request failed with not found status %d", status)
		utils.SendErrorResponse(w, http.StatusNotFound, decodedResponse.Message, "Doctor not found: "+decodedResponse.Error)
		return
	default:
		log.Printf("[GATEWAY] DeleteDoctorByID: Request failed with unexpected status %d", status)
		utils.SendErrorResponse(w, http.StatusInternalServerError, decodedResponse.Message, "Unexpected status code: "+strconv.Itoa(status)+". Error: "+decodedResponse.Error)
		return
	}
}
