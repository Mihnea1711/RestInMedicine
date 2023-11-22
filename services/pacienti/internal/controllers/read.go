package controllers

import (
	"context"
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

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	patients, err := pController.DbConn.FetchPatients(ctx, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("Internal server error: %s", err)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Failed to fetch patients"})
		return
	}

	log.Printf("[PATIENT] Successfully fetched %d patients", len(patients))
	// Use utils.RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{Payload: patients, Message: "Patients fetched successfully"})
}

func (pController *PatientController) GetPatientByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientIDStr := vars[utils.FETCH_PATIENT_BY_ID_PARAMETER]

	patientID, err := strconv.Atoi(patientIDStr)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid patient ID: %s", patientIDStr)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Invalid request"})
		return
	}

	log.Printf("[PATIENT] Fetching patient with ID: %d...", patientID)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	patient, err := pController.DbConn.FetchPatientByID(ctx, patientID)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to fetch patient with ID %d: %s", patientID, err)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Failed to fetch patient by id"})
		return
	}

	if patient == nil {
		errMsg := fmt.Sprintf("No patient found with ID: %d", patientID)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Patient not found or an unexpected error happened."})
		return
	}

	log.Printf("[PATIENT] Successfully fetched patient with ID %d", patientID)
	// Use utils.RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{Payload: patient, Message: "Patient fetched successfully"})
}

func (pController *PatientController) GetPatientByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	patientEmail := vars[utils.FETCH_PATIENT_BY_EMAIL_PARAMETER]

	log.Printf("[PATIENT] Fetching patient with email: %s...", patientEmail)

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	patient, err := pController.DbConn.FetchPatientByEmail(ctx, patientEmail, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to fetch patient with email %s: %s", patientEmail, err)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Failed to fetch patient by email"})
		return
	}

	if patient == nil {
		errMsg := fmt.Sprintf("No patient found with email: %s", patientEmail)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Patient not found or an unexpected error happened."})
		return
	}

	log.Printf("[PATIENT] Successfully fetched patient with email %s", patientEmail)
	// Use utils.RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{Payload: patient, Message: "Patient fetched successfully"})
}

func (pController *PatientController) GetPatientByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDString := vars[utils.FETCH_PATIENT_BY_USER_ID_PARAMETER]

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid User ID: %s", userIDString)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Invalid request"})
		return
	}

	log.Printf("[PATIENT] Fetching patient with user ID: %d...", userID)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	patient, err := pController.DbConn.FetchPatientByUserID(ctx, userID)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to fetch patient with user ID %d: %s", userID, err)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Failed to fetch patient by userID"})
		return
	}

	if patient == nil {
		errMsg := fmt.Sprintf("No patient found with user ID: %d", userID)
		log.Printf("[PATIENT] %s", errMsg)

		// Use utils.RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Patient not found or an unexpected error happened."})
		return
	}

	log.Printf("[PATIENT] Successfully fetched patient with user ID %d", userID)
	// Use utils.RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{Payload: patient, Message: "Patient fetched successfully"})
}
