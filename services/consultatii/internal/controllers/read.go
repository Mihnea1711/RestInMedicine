package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func (cController *ConsultationController) GetConsultations(w http.ResponseWriter, r *http.Request) {
	log.Printf("[CONSULTATION] Attempting to retrieve consultations.")

	// Extract query filter params
	filters, err := utils.ExtractFiltersFromRequest(r)
	if err != nil {
		errMsg := fmt.Sprintf("bad request: %s", err)
		log.Printf("[CONSULTATION] GetAppointmentsByFilter: Failed to extract filters: %s", errMsg)

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

	// Ensure a database operation doesn't take longer than utils.REQUEST_TIMEOUT_DURATION seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	cController.handleContextTimeout(ctx, w)

	// Use cController.DbConn to fetch filtered consultations from the database
	consultations, err := cController.DbConn.FetchConsultations(ctx, filters, page, limit)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[CONSULTATION] Failed to fetch filtered consultations: %s\n", errMsg)
		utils.RespondWithJSON(w, http.StatusInternalServerError, errMsg)
		return
	}

	if len(consultations) != 0 {
		log.Printf("[CONSULTATION] Successfully retrieved %d consultations", len(consultations))
	} else {
		log.Println("[CONSULTATION] No consultations found with the filter")
	}
	// Serialize the filtered consultations to JSON and send the response
	response := models.ResponseData{
		Message: "Consultations retrieved successfully.",
		Payload: consultations,
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

	cController.handleContextTimeout(ctx, w)

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
	log.Printf("[CONSULTATION] Successfully fetched consultation %s", consultation.IDConsultation)

	// Serialize the consultation to JSON and send the response
	response := models.ResponseData{
		Message: "Consultation retrieved successfully.",
		Payload: consultation,
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}
