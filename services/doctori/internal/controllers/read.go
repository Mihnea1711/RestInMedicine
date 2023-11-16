package controllers

import (
	"context"
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

	doctors, err := dController.DbConn.FetchDoctors(ctx, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg})
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{Payload: doctors})
}

func (dController *DoctorController) GetDoctorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars[utils.FETCH_DOCTOR_BY_ID_PARAMETER]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid doctor ID: %s", idStr)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg})
		return
	}

	log.Printf("[DOCTOR] Fetching doctor with ID: %d...", id)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	doctor, err := dController.DbConn.FetchDoctorByID(ctx, id)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to fetch doctor with ID %d: %s", id, err)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg})
		return
	}

	if doctor == nil {
		errMsg := fmt.Sprintf("No doctor found with ID: %d", id)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg})
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{Payload: doctor})
}

func (dController *DoctorController) GetDoctorByEmail(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars[utils.FETCH_DOCTOR_BY_EMAIL_PARAMETER]

	log.Printf("[DOCTOR] Fetching doctor with email: %s...", email)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	doctor, err := dController.DbConn.FetchDoctorByEmail(ctx, email)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to fetch doctor with email %s: %s", email, err)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg})
		return
	}

	if doctor == nil {
		errMsg := fmt.Sprintf("No doctor found with email: %s", email)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg})
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{Payload: doctor})
}

func (dController *DoctorController) GetDoctorByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userIDString := vars[utils.FETCH_DOCTOR_BY_USER_ID_PARAMETER]

	userID, err := strconv.Atoi(userIDString)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid User ID: %s", userIDString)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg})
		return
	}

	log.Printf("[DOCTOR] Fetching doctor with user ID: %d...", userID)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	doctor, err := dController.DbConn.FetchDoctorByUserID(ctx, userID)
	if err != nil {
		errMsg := fmt.Sprintf("Failed to fetch doctor with user ID %d: %s", userID, err)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg})
		return
	}

	if doctor == nil {
		errMsg := fmt.Sprintf("No doctor found with user ID: %d", userID)
		log.Println("[DOCTOR] " + errMsg)
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg})
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{Payload: doctor})
}
