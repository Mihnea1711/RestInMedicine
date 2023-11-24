package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Retrieve all consultations
func (cController *ConsultationController) GetConsultations(w http.ResponseWriter, r *http.Request) {
	log.Printf("[CONSULTATION] Attempting to retrieve consultations.")

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than utils.REQUEST_TIMEOUT_DURATION seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	// Use cController.DbConn to fetch all consultations from the database
	consultations, err := cController.DbConn.FetchAllConsultations(ctx, page, limit)
	if err != nil {
		// Handle the case where an internal server error occurs during fetch
		errMsg := fmt.Sprintf("Internal server error: %s", err)
		log.Printf("[CONSULTATION] Failed to fetch consultations: %s\n", errMsg)

		// Respond with an error using the ResponseData struct
		response := models.ResponseData{
			Error:   errMsg,
			Message: "Failed to retrieve consultations. Internal server error.",
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	log.Printf("[CONSULTATION] Successfully fetched %d consultations", len(consultations))

	// Serialize the consultations to JSON and send the response
	response := models.ResponseData{
		Message: fmt.Sprintf("%d Consultations retrieved successfully.", len(consultations)), Payload: consultations,
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// Retrieve a consultation by ID
func (cController *ConsultationController) GetConsultationByID(w http.ResponseWriter, r *http.Request) {
	// Log the attempt to retrieve a consultation by ID
	log.Printf("[CONSULTATION] Attempting to retrieve a consultation by ID.")

	// Extract the consultation ID from the request URL parameters
	vars := mux.Vars(r)
	consultationID := vars[utils.FETCH_CONSULTATIE_BY_ID_PARAMETER]

	// Convert the ID to a primitive.ObjectID
	objectID, err := primitive.ObjectIDFromHex(consultationID)
	if err != nil {
		// Handle the case where an invalid consultation ID is provided
		response := models.ResponseData{
			Error:   "Invalid consultation ID",
			Message: "Invalid consultation ID. Please provide a valid ID.",
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	// Ensure a database operation doesn't take longer than utils.REQUEST_TIMEOUT_DURATION seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	// Use cController.DbConn to fetch the consultation by ID from the database
	consultation, err := cController.DbConn.FetchConsultationByID(ctx, objectID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Handle the case where no document is found with the given ID
			errMsg := fmt.Sprintf("Failed to get consultation by ID. Consultation not found: %s", err)
			log.Printf("[CONSULTATION] %s: %v", errMsg, consultationID)
			response := models.ResponseData{
				Error:   errMsg,
				Message: "Failed to get consultation by ID. Consultation not found.",
			}
			utils.RespondWithJSON(w, http.StatusNotFound, response)
			return
		}

		// Handle the case where an internal server error occurs during fetch
		errMsg := fmt.Sprintf("Internal server error: %s", err)
		log.Printf("[CONSULTATION] Failed to fetch consultation by ID: %s\n", errMsg)

		// Respond with an error using the ResponseData struct
		response := models.ResponseData{
			Error:   errMsg,
			Message: "Failed to retrieve consultation by ID. Internal server error.",
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Check if the consultation exists
	if consultation == nil {
		// Handle the case where the consultation is not found
		response := models.ResponseData{
			Error:   "Consultation not found",
			Message: "Consultation not found with the provided ID.",
		}
		utils.RespondWithJSON(w, http.StatusNotFound, response)
		return
	}

	// Log the successful retrieval of the consultation
	log.Printf("[CONSULTATION] Successfully fetched consultation %s", consultation.IDConsultatie)

	// Serialize the consultation to JSON and send the response
	response := models.ResponseData{
		Message: "Consultation retrieved successfully.",
		Payload: consultation,
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// Retrieve consultations by Doctor ID
func (cController *ConsultationController) GetConsultationsByDoctorID(w http.ResponseWriter, r *http.Request) {
	// Log the attempt to retrieve consultations by Doctor ID
	log.Printf("[CONSULTATION] Attempting to retrieve consultations by Doctor ID.")

	// Extract Doctor ID from the request URL parameters
	vars := mux.Vars(r)
	doctorID, err := strconv.Atoi(vars[utils.FETCH_CONSULTATIE_BY_DOCTOR_ID_PARAMETER])
	if err != nil {
		// Handle the case where an invalid Doctor ID is provided
		response := models.ResponseData{
			Error:   "Invalid Doctor ID",
			Message: "Invalid Doctor ID. Please provide a valid ID.",
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	// Log the Doctor ID being read
	log.Printf("Doctor ID: %d", doctorID)

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than utils.REQUEST_TIMEOUT_DURATION seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	// Use cController.DbConn to fetch consultations by Doctor ID from the database
	consultations, err := cController.DbConn.FetchConsultationsByDoctorID(ctx, doctorID, page, limit)
	if err != nil {
		// Handle the case where an internal server error occurs during fetch
		errMsg := fmt.Sprintf("Internal server error: %s", err)
		log.Printf("[CONSULTATION] Failed to fetch consultations by Doctor ID: %s\n", errMsg)

		// Respond with an error using the ResponseData struct
		response := models.ResponseData{
			Error:   errMsg,
			Message: "Failed to retrieve consultations by Doctor ID. Internal server error.",
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Log the successful retrieval of consultations
	log.Printf("[CONSULTATION] Successfully fetched %d consultations of Doctor %d", len(consultations), doctorID)

	// Serialize the consultations to JSON and send the response
	response := models.ResponseData{
		Message: fmt.Sprintf("%d Consultations retrieved successfully.", len(consultations)), Payload: consultations,
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// Retrieve consultations by Patient ID
func (cController *ConsultationController) GetConsultationsByPatientID(w http.ResponseWriter, r *http.Request) {
	// Log the attempt to retrieve consultations by Patient ID
	log.Printf("[CONSULTATION] Attempting to retrieve consultations by Patient ID.")

	// Extract Patient ID from the request URL parameters
	vars := mux.Vars(r)
	patientID, err := strconv.Atoi(vars[utils.FETCH_CONSULTATIE_BY_PATIENT_ID_PARAMETER])
	if err != nil {
		// Handle the case where an invalid Patient ID is provided
		response := models.ResponseData{
			Error:   "Invalid Patient ID",
			Message: "Invalid Patient ID. Please provide a valid ID.",
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	// Log the Patient ID being read
	log.Printf("Patient ID: %d", patientID)

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than utils.REQUEST_TIMEOUT_DURATION seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	// Use cController.DbConn to fetch consultations by Patient ID from the database
	consultations, err := cController.DbConn.FetchConsultationsByPatientID(ctx, patientID, page, limit)
	if err != nil {
		// Handle the case where an internal server error occurs during fetch
		errMsg := fmt.Sprintf("Internal server error: %s", err)
		log.Printf("[CONSULTATION] Failed to fetch consultations by Patient ID: %s\n", errMsg)

		// Respond with an error using the ResponseData struct
		response := models.ResponseData{
			Error:   errMsg,
			Message: "Failed to retrieve consultations by Patient ID. Internal server error.",
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Log the successful retrieval of consultations
	log.Printf("[CONSULTATION] Successfully fetched %d consultations of Patient %d", len(consultations), patientID)

	// Serialize the consultations to JSON and send the response
	response := models.ResponseData{
		Message: fmt.Sprintf("%d Consultations retrieved successfully.", len(consultations)), Payload: consultations,
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}

// Retrieve consultations by date
func (cController *ConsultationController) GetConsultationsByDate(w http.ResponseWriter, r *http.Request) {
	// Log the attempt to retrieve consultations by date
	log.Printf("[CONSULTATION] Attempting to retrieve consultations by date.")

	// Use mux.Vars to get the date parameter from the route
	vars := mux.Vars(r)
	dateStr := vars[utils.FETCH_CONSULTATIE_BY_DATE_PARAMETER]

	// Parse the date string into a time.Time object
	date, err := time.Parse(utils.TIME_FORMAT, dateStr)
	if err != nil {
		// Handle the case where an invalid date format is provided
		errResponse := models.ResponseData{
			Error:   "Invalid date format",
			Message: "Invalid date format. Please provide a date in the format: " + utils.TIME_FORMAT,
		}
		log.Printf("[CONSULTATION] Failed to convert date string: %s\n", err)
		utils.RespondWithJSON(w, http.StatusBadRequest, errResponse)
		return
	}

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than utils.REQUEST_TIMEOUT_DURATION seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	// Use cController.DbConn to fetch consultations by date from the database
	consultations, err := cController.DbConn.FetchConsultationsByDate(ctx, date, page, limit)
	if err != nil {
		// Handle the case where an internal server error occurs during fetch
		errMsg := fmt.Sprintf("Internal server error: %s", err)
		errResponse := models.ResponseData{
			Error:   errMsg,
			Message: "Failed to retrieve consultations by date. Internal server error.",
		}
		log.Printf("[CONSULTATION] Failed to fetch consultations by date: %s\n", errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, errResponse)
		return
	}

	// Log the successful retrieval of consultations
	log.Printf("[CONSULTATION] Successfully fetched %d consultations from %s", len(consultations), date)

	// Serialize the consultations to JSON and send the response
	response := models.ResponseData{
		Message: fmt.Sprintf("%d Consultations retrieved successfully.", len(consultations)),
		Payload: consultations,
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}

func (cController *ConsultationController) GetFilteredConsultations(w http.ResponseWriter, r *http.Request) {
	log.Printf("[CONSULTATION] Attempting to retrieve filtered consultations.")

	// Extract query filter params
	filter := utils.ExtractQueryParams(r)

	// Extract the limit and page query parameters from the request
	limit, page := utils.ExtractPaginationParams(r)

	// Ensure a database operation doesn't take longer than utils.REQUEST_TIMEOUT_DURATION seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	// Use cController.DbConn to fetch filtered consultations from the database
	consultations, err := cController.DbConn.FetchConsultationsByFilter(ctx, filter, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[CONSULTATION] Failed to fetch filtered consultations: %s\n", errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, errMsg)
		return
	}

	if len(consultations) != 0 {
		log.Printf("[CONSULTATION] Successfully fetched filtered consultations")
	} else {
		log.Printf("[CONSULTATION] No consultations found with the filter: %v", filter)
	}
	// Serialize the filtered consultations to JSON and send the response
	utils.RespondWithJSON(w, http.StatusOK, consultations)
}
