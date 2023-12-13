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

// Delete a consultation by ID
func (cController *ConsultationController) DeleteConsultationByID(w http.ResponseWriter, r *http.Request) {
	// Log the attempt to delete a consultation by ID
	log.Printf("[CONSULTATION] Attempting to delete a consultation by ID.")

	// Extract the consultation ID from the request URL parameters
	vars := mux.Vars(r)
	consultationID, err := primitive.ObjectIDFromHex(vars[utils.DELETE_CONSULTATIE_BY_ID_PARAMETER])
	if err != nil {
		// Handle the case where an invalid consultation ID is provided
		response := models.ResponseData{
			Error:   "Invalid consultation ID",
			Message: "Failed to delete consultation. Invalid consultation ID provided.",
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	// Ensure a database operation doesn't take longer than utils.REQUEST_TIMEOUT_DURATION seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	cController.handleContextTimeout(ctx, w)

	// Use cController.DbConn to delete the consultation by ID from the database
	rowsAffected, err := cController.DbConn.DeleteConsultationByID(ctx, consultationID)
	if err != nil {
		handleDatabaseDeleteError(w, err, consultationID)
		return
	}

	// Check if the consultation exists and was deleted
	if rowsAffected == 0 {
		// Handle the case where the consultation is not found
		response := models.ResponseData{
			Error:   "Consultation not found",
			Message: "Failed to delete consultation. Consultation not found.",
		}
		utils.RespondWithJSON(w, http.StatusNotFound, response)
		return
	}

	// Log the successful deletion of the consultation
	log.Printf("[CONSULTATION] Successfully deleted consultation %s", consultationID.Hex())

	// Respond with a success message
	response := models.ResponseData{
		Message: "Consultation deleted successfully.",
		Payload: models.RowsAffected{
			RowsAffected: rowsAffected,
		},
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}

func handleDatabaseDeleteError(w http.ResponseWriter, err error, id interface{}) {
	if err == mongo.ErrNoDocuments {
		// Handle the case where no document is found with the given ID
		errMsg := fmt.Sprintf("Failed to delete consultation by ID. Consultation not found: %s", err)
		log.Printf("[CONSULTATION] %s: %v", errMsg, id)
		response := models.ResponseData{
			Error:   errMsg,
			Message: "Failed to delete consultation. Consultation not found.",
		}
		utils.RespondWithJSON(w, http.StatusNotFound, response)
		return
	}

	// Handle the case where an internal server error occurs during the deletion
	errMsg := fmt.Sprintf("Internal server error: %s", err)
	log.Printf("[CONSULTATION] Failed to delete consultation by ID: %s\n", errMsg)
	response := models.ResponseData{
		Error:   errMsg,
		Message: "Failed to delete consultation. Internal server error.",
	}
	utils.RespondWithJSON(w, http.StatusInternalServerError, response)
}
