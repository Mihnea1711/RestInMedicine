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
	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func (dController *DoctorController) GetDoctors(w http.ResponseWriter, r *http.Request) {
	log.Println("[DOCTOR] Fetching all doctors...")

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	log.Printf("[DOCTOR] Fetching doctors with limit: %d, page: %d", limit, page)

	doctors, err := dController.DbConn.FetchDoctors(ctx, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[DOCTOR] Error fetching doctors: %s", errMsg)

		// Use RespondWithJSON for error response with a message
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
			Message: "Failed to fetch doctors",
			Error:   errMsg,
		})
		return
	}

	log.Printf("[DOCTOR] Successfully fetched %d doctors", len(doctors))
	// Use RespondWithJSON for success response with a message
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: fmt.Sprintf("Successfully fetched %d doctors", len(doctors)),
		Payload: doctors,
	})
}

func (dController *DoctorController) GetDoctorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	doctorIDStr := vars[utils.FETCH_DOCTOR_BY_ID_PARAMETER]

	doctorID, err := strconv.Atoi(doctorIDStr)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid doctor ID: %s", doctorIDStr)
		log.Printf("[DOCTOR] %s", errMsg)
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Message: "Bad Request", Error: errMsg})
		return
	}

	log.Printf("[DOCTOR] Fetching doctor with ID: %d...", doctorID)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	doctor, err := dController.DbConn.FetchDoctorByID(ctx, doctorID)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Error getting doctor by ID: %s", err.Error())
			log.Printf("[DOCTOR] %s", errMsg)
			// Create a conflict response using ResponseData
			utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to get doctor by ID. Doctor not found"})
			return
		}

		errMsg := fmt.Sprintf("Failed to fetch doctor with ID %d: %s", doctorID, err)
		log.Printf("[DOCTOR] %s", errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Message: "Failed to fetch doctor by id", Error: errMsg})
		return
	}

	if doctor == nil {
		errMsg := fmt.Sprintf("No doctor found with ID: %d", doctorID)
		log.Printf("[DOCTOR] %s", errMsg)
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Message: "Doctor not found or an unexpected error happened.", Error: errMsg})
		return
	}

	log.Printf("[DOCTOR] Successfully fetched doctor with ID: %d", doctorID)
	// Use RespondWithJSON for success response with a message
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: fmt.Sprintf("Successfully fetched doctor with ID: %d", doctorID),
		Payload: doctor,
	})
}

func (dController *DoctorController) GetDoctorByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	doctorEmail := vars[utils.FETCH_DOCTOR_BY_EMAIL_PARAMETER]

	log.Printf("[DOCTOR] Fetching doctor with email: %s...", doctorEmail)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	doctor, err := dController.DbConn.FetchDoctorByEmail(ctx, doctorEmail)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Error getting doctor by email: %s", err.Error())
			log.Printf("[DOCTOR] %s", errMsg)
			// Create a conflict response using ResponseData
			utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to get doctor by email. Doctor not found"})
			return
		}

		errMsg := fmt.Sprintf("Failed to fetch doctor with email %s: %s", doctorEmail, err)
		log.Printf("[DOCTOR] %s", errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Message: "Failed to fetch doctor by email", Error: errMsg})
		return
	}

	if doctor == nil {
		errMsg := fmt.Sprintf("No doctor found with email: %s", doctorEmail)
		log.Printf("[DOCTOR] %s", errMsg)
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Message: "Doctor not found or an unexpected error happened.", Error: errMsg})
		return
	}

	log.Printf("[DOCTOR] Successfully fetched doctor with email: %s", doctorEmail)
	// Use RespondWithJSON for success response with a message
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: fmt.Sprintf("Successfully fetched doctor with email: %s", doctorEmail),
		Payload: doctor,
	})
}

func (dController *DoctorController) GetDoctorByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDString := vars[utils.FETCH_DOCTOR_BY_USER_ID_PARAMETER]

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid User ID: %s", userIDString)
		log.Printf("[DOCTOR] %s", errMsg)
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Message: "Bad Request", Error: errMsg})
		return
	}

	log.Printf("[DOCTOR] Fetching doctor with user ID: %d...", userID)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	doctor, err := dController.DbConn.FetchDoctorByUserID(ctx, userID)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Error getting doctor by user ID: %s", err.Error())
			log.Printf("[DOCTOR] %s", errMsg)
			// Create a conflict response using ResponseData
			utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to get doctor by user ID. Doctor not found"})
			return
		}

		errMsg := fmt.Sprintf("Failed to fetch doctor with user ID %d: %s", userID, err)
		log.Printf("[DOCTOR] %s", errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Message: "Failed to fetch doctor by userID", Error: errMsg})
		return
	}

	if doctor == nil {
		errMsg := fmt.Sprintf("No doctor found with user ID: %d", userID)
		log.Printf("[DOCTOR] %s", errMsg)
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Message: "Doctor not found or an unexpected error happened.", Error: errMsg})
		return
	}

	log.Printf("[DOCTOR] Successfully fetched doctor with user ID: %d", userID)
	// Use RespondWithJSON for success response with a message
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: fmt.Sprintf("Successfully fetched doctor with user ID: %d", userID),
		Payload: doctor,
	})
}
