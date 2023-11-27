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

func (dController *DoctorController) DeleteDoctorByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	doctorIDStr := vars[utils.DELETE_DOCTOR_BY_ID_PARAMETER]

	doctorID, err := strconv.Atoi(doctorIDStr)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid doctor ID: %s", doctorIDStr)
		log.Println("[DOCTOR] " + errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Invalid patient delete request"})
		return
	}

	log.Printf("[DOCTOR] Attempting to delete doctor with ID: %d...", doctorID)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	rowsAffected, err := dController.DbConn.DeleteDoctorByID(ctx, doctorID)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Error deleting doctor: %s", err.Error())
			log.Printf("[DOCTOR] %s", errMsg)
			// Create a conflict response using ResponseData
			utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to delete doctor. Doctor not found"})
			return
		}

		errMsg := fmt.Sprintf("Failed to delete doctor with ID %d: %s", doctorID, err)
		log.Println("[DOCTOR] " + errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Failed to delete patient"})
		return
	}

	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("No doctor found with ID: %d", doctorID)
		log.Println("[DOCTOR] " + errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Patient not found or an unexpected error happened."})
		return
	}

	log.Printf("[DOCTOR] Successfully deleted doctor with ID: %d", doctorID)

	// Use RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: fmt.Sprintf("Doctor with ID: %d deleted successfully", doctorID),
		Payload: models.RowsAffected{
			RowsAffected: rowsAffected,
		},
	})
}

func (dController *DoctorController) DeleteDoctorByUserID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	doctorUserIDStr := vars[utils.DELETE_DOCTOR_BY_USER_ID_PARAMETER]
	doctorUserID, err := strconv.Atoi(doctorUserIDStr)
	if err != nil {
		errMsg := fmt.Sprintf("Invalid doctor ID: %s", doctorUserIDStr)
		log.Println("[DOCTOR] " + errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusBadRequest, models.ResponseData{Error: errMsg, Message: "Invalid patient delete request"})
		return
	}

	log.Printf("[DOCTOR] Attempting to delete doctor with user ID: %d...", doctorUserID)

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.DB_REQ_TIMEOUT_SEC_MULTIPLIER*time.Second)
	defer cancel()

	// get doctor by user id
	doctor, err := dController.DbConn.FetchDoctorByUserID(ctx, doctorUserID)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Error getting doctor by user ID: %s", err.Error())
			log.Printf("[DOCTOR] %s", errMsg)
			// Create a conflict response using ResponseData
			utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to get doctor by user ID. Doctor not found"})
			return
		}

		errMsg := fmt.Sprintf("Failed to fetch doctor with user ID %d: %s", doctorUserID, err)
		log.Printf("[DOCTOR] %s", errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Message: "Failed to fetch doctor by doctorUserID", Error: errMsg})
		return
	}

	if doctor == nil {
		errMsg := fmt.Sprintf("No doctor found with user ID: %d", doctorUserID)
		log.Printf("[DOCTOR] %s", errMsg)
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Message: "Doctor not found or an unexpected error happened.", Error: errMsg})
		return
	}

	rowsAffected, err := dController.DbConn.DeleteDoctorByUserID(ctx, doctorUserID)
	if err != nil {
		// Check if the error is due to no rows found
		if err == sql.ErrNoRows {
			errMsg := fmt.Sprintf("Error deleting doctor: %s", err.Error())
			log.Printf("[DOCTOR] %s", errMsg)
			// Create a conflict response using ResponseData
			utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Failed to delete doctor. Doctor not found"})
			return
		}

		errMsg := fmt.Sprintf("Failed to delete doctor with user ID %d: %s", doctorUserID, err)
		log.Println("[DOCTOR] " + errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{Error: errMsg, Message: "Failed to delete doctor"})
		return
	}

	if rowsAffected == 0 {
		errMsg := fmt.Sprintf("No doctor found with user ID: %d", doctorUserID)
		log.Println("[DOCTOR] " + errMsg)

		// Use RespondWithJSON for error response
		utils.RespondWithJSON(w, http.StatusNotFound, models.ResponseData{Error: errMsg, Message: "Patient not found or an unexpected error happened."})
		return
	}

	log.Printf("[DOCTOR] Successfully deleted doctor with user ID: %d", doctorUserID)

	// Use RespondWithJSON for success response
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: fmt.Sprintf("Doctor with user ID: %d deleted successfully", doctorUserID),
		Payload: models.ComplexResponse{
			RowsAffected: rowsAffected,
			DeletedID:    doctor.IDDoctor,
		},
	})
}
