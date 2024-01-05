package controllers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/models"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create a new consultation
func (cController *ConsultationController) CreateConsultation(w http.ResponseWriter, r *http.Request) {
	log.Printf("[CONSULTATION] Attempting to create a new consultation.")

	// Extract the consultation from the request context
	consultation := r.Context().Value(utils.DECODED_CONSULTATION).(*models.Consultation)

	// Ensure a database operation doesn't take longer than utils.REQUEST_TIMEOUT_DURATION seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	cController.handleContextTimeout(ctx, w)

	// Assign an ID to the consultation
	consultation.IDConsultation = primitive.NewObjectID()

	// Use cController.DbConn to save the consultation to the database
	insertedID, err := cController.DbConn.SaveConsultation(ctx, consultation)
	if err != nil {
		handleDatabaseCreateError(w, err)
		return
	}

	// Check if the inserted ID is valid
	if insertedID != primitive.NilObjectID {
		// Log the successful creation of the consultation
		log.Printf("[CONSULTATION] Successfully created consultation with ID: %s", insertedID.Hex())

		// Respond with success using the ResponseData struct
		response := models.ResponseData{
			Message: "Consultation created successfully.",
			Payload: models.LastInsertedID{
				LastInsertedID: insertedID.String(),
			},
		}
		utils.RespondWithJSON(w, http.StatusCreated, response)
		return
	} else {
		// Log the failure to get a valid inserted ID
		log.Printf("[CONSULTATION] Failed to get valid inserted ID after creating consultation.")

		// Respond with an error using the ResponseData struct
		response := models.ResponseData{
			Error:   "Failed to get valid inserted ID after creating consultation",
			Message: "Failed to create consultation. Error getting valid inserted ID.",
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}
}

func handleDatabaseCreateError(w http.ResponseWriter, err error) {
	// Check for a duplicate key error
	if writeException, ok := err.(mongo.WriteException); ok {
		for _, writeError := range writeException.WriteErrors {
			if writeError.Code == utils.DUPLICATE_KEY_ERROR_CODE {
				// Handle the case where a duplicate key error occurs
				errMsg := fmt.Sprintf("Duplicate key error: %s", writeError.Error())
				log.Printf("[CONSULTATION] Failed to save consultation to the database: %s\n", errMsg)

				// Respond with an error using the ResponseData struct
				response := models.ResponseData{
					Error:   errMsg,
					Message: "Failed to create consultation. Duplicate key violation.",
				}
				utils.RespondWithJSON(w, http.StatusConflict, response)
				return
			}
		}
	}

	// Handle the case where an internal server error occurs during the save
	errMsg := fmt.Sprintf("Internal server error: %s", err)
	log.Printf("[CONSULTATION] Failed to save consultation to the database: %s\n", errMsg)

	// Respond with an error using the ResponseData struct
	response := models.ResponseData{
		Error:   errMsg,
		Message: "Failed to create consultation. Internal server error.",
	}
	utils.RespondWithJSON(w, http.StatusInternalServerError, response)
}
