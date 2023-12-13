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
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/models"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func (pController *PatientController) GetPatients(w http.ResponseWriter, r *http.Request) {
	log.Println("[PATIENT] Fetching all patients...")

	// Extract query params from request
	filters, err := utils.ExtractFiltersFromRequest(r)
	if err != nil {
		errMsg := fmt.Sprintf("bad request: %s", err)
		log.Printf("[DOCTOR] GetPatients: Failed to extract filters: %s", errMsg)

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

	pController.handleContextTimeout(ctx, w)

	log.Printf("[PATIENT] Fetching patients with limit: %d, page: %d", limit, page)

	patients, err := pController.DbConn.FetchPatients(ctx, filters, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("failed to fetch patients: %v", err)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Failed to fetch patients"})
		return
	}

	log.Printf("[PATIENT] Successfully fetched %d patients", len(patients))
	// Use utils.RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{Payload: patients, Message: fmt.Sprintf("Successfully fetched %d patients", len(patients))})
}

func (pController *PatientController) GetPatientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientIDStr := vars[utils.FETCH_PATIENT_BY_ID_PARAMETER]

	patientID, err := strconv.Atoi(patientIDStr)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid patient ID: %s", patientIDStr)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Bad request"})
		return
	}

	log.Printf("[PATIENT] Fetching patient with ID: %d...", patientID)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	pController.handleContextTimeout(ctx, w)

	patient, err := pController.DbConn.FetchPatientByID(ctx, patientID)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Failed to fetch patient with ID %d: %s", patientID, err)
			log.Printf("[PATIENT] %s", errMsg)

			// Use utils.RespondWithJSON for error response
			utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to fetch patient by ID. Patient not found"})
			return
		}

		errMsg := fmt.Sprintf("Failed to fetch patient with ID %d: %s", patientID, err)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Failed to fetch patient by id"})
		return
	}

	if patient == nil || !patient.IsActive {
		errMsg := fmt.Sprintf("No patient found with ID: %d", patientID)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Patient not found or an unexpected error happened."})
		return
	}

	log.Printf("[PATIENT] Successfully fetched patient with ID %d", patientID)
	// Use utils.RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{Payload: patient, Message: fmt.Sprintf("Successfully fetched patient with ID: %d", patientID)})
}

func (pController *PatientController) GetPatientByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientEmail := vars[utils.FETCH_PATIENT_BY_EMAIL_PARAMETER]

	log.Printf("[PATIENT] Fetching patient with email: %s...", patientEmail)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	pController.handleContextTimeout(ctx, w)

	patient, err := pController.DbConn.FetchPatientByEmail(ctx, patientEmail)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Failed to fetch patient with email %s: %s", patientEmail, err)
			log.Printf("[PATIENT] %s", errMsg)

			// Use utils.RespondWithJSON for error response
			utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to fetch patient by email. Patient not found"})
			return
		}

		errMsg := fmt.Sprintf("Failed to fetch patient with email %s: %s", patientEmail, err)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Failed to fetch patient by email"})
		return
	}

	if patient == nil || !patient.IsActive {
		errMsg := fmt.Sprintf("No patient found with email: %s", patientEmail)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Patient not found or an unexpected error happened."})
		return
	}

	log.Printf("[PATIENT] Successfully fetched patient with email %s", patientEmail)
	// Use utils.RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{Payload: patient, Message: fmt.Sprintf("Successfully fetched patient with email: %s", patientEmail)})
}

func (pController *PatientController) GetPatientByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDString := vars[utils.FETCH_PATIENT_BY_USER_ID_PARAMETER]

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid User ID: %s", userIDString)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Bad request"})
		return
	}

	log.Printf("[PATIENT] Fetching patient with user ID: %d...", userID)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	pController.handleContextTimeout(ctx, w)

	patient, err := pController.DbConn.FetchPatientByUserID(ctx, userID)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Failed to fetch patient with user ID %s: %s", userIDString, err)
			log.Printf("[PATIENT] %s", errMsg)

			// Use utils.RespondWithJSON for error response
			utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to fetch patient by user ID. Patient not found"})
			return
		}

		errMsg := fmt.Sprintf("Failed to fetch patient with user ID %d: %s", userID, err)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Failed to fetch patient by userID"})
		return
	}

	if patient == nil || !patient.IsActive {
		errMsg := fmt.Sprintf("No patient found with user ID: %d", userID)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Patient not found or an unexpected error happened."})
		return
	}

	log.Printf("[PATIENT] Successfully fetched patient with user ID %d", userID)
	// Use utils.RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{Payload: patient, Message: fmt.Sprintf("Successfully fetched patient with user ID: %d", userID)})
}
