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

// UpdateConsultationByID updates a consultation by its ID.
func (cController *ConsultationController) UpdateConsultationByID(w http.ResponseWriter, r *http.Request) {
	// Log the attempt to update a consultation by ID
	log.Printf("[CONSULTATION] Attempting to update a consultation by ID.")

	// Extract the consultation ID from the request URL parameters
	vars := mux.Vars(r)
	consultationID, err := primitive.ObjectIDFromHex(vars[utils.UPDATE_CONSULTATIE_BY_ID_PARAMETER])
	if err != nil {
		// Handle the case where an invalid consultation ID is provided
		response := models.ResponseData{
			Error:   "Invalid consultation ID",
			Message: "Failed to update consultation. Invalid consultation ID provided.",
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	// Retrieve the consultation data from the request context
	consultation := r.Context().Value(utils.DECODED_CONSULTATION).(*models.Consultation)
	consultation.IDConsultatie = consultationID

	// Ensure a database operation doesn't take longer than utils.REQUEST_TIMEOUT_DURATION seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	// Use cController.DbConn to update the consultation by ID in the database
	rowsAffected, err := cController.DbConn.UpdateConsultationByID(ctx, consultation)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			// Handle the case where no document is found with the given ID
			errMsg := fmt.Sprintf("Failed to update consultation by ID. Consultation not found: %s", err)
			log.Printf("[CONSULTATION] %s: %v", errMsg, consultationID)
			response := models.ResponseData{
				Error:   errMsg,
				Message: "Failed to update consultation. Consultation not found.",
			}
			utils.RespondWithJSON(w, http.StatusNotFound, response)
			return
		}

		// Check for a duplicate key error
		if writeException, ok := err.(mongo.WriteException); ok {
			for _, writeError := range writeException.WriteErrors {
				if writeError.Code == utils.DUPLICATE_KEY_ERROR_CODE {
					// Handle the case where a duplicate key error occurs
					errMsg := fmt.Sprintf("Duplicate key error: %s", writeError.Error())
					log.Printf("[CONSULTATION] Failed to update consultation to the database: %s\n", errMsg)

					// Respond with an error using the ResponseData struct
					response := models.ResponseData{
						Error:   errMsg,
						Message: "Failed to update consultation. Duplicate key violation.",
					}
					utils.RespondWithJSON(w, http.StatusConflict, response)
					return
				}
			}
		}

		// Handle the case where an internal server error occurs during the update
		errMsg := fmt.Sprintf("Internal server error: %s", err)
		log.Printf("[CONSULTATION] Failed to update consultation by ID: %s\n", errMsg)
		response := models.ResponseData{
			Error:   errMsg,
			Message: "Failed to update consultation. Internal server error.",
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Check if the consultation exists and was updated
	if rowsAffected == 0 {
		// Handle the case where the consultation is not found
		response := models.ResponseData{
			Error:   "Consultation not found",
			Message: "Failed to update consultation. Consultation not found.",
		}
		utils.RespondWithJSON(w, http.StatusNotFound, response)
		return
	}

	// Log the successful update of the consultation
	log.Printf("[CONSULTATION] Successfully updated consultation %s", consultation.IDConsultatie.Hex())

	// Respond with a success message and the number of rows affected
	response := models.ResponseData{
		Message: "Consultation updated successfully.",
		Payload: models.RowsAffected{
			RowsAffected: rowsAffected,
		},
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}
