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
)

// Update a consultatie by ID
func (cController *ConsultatieController) UpdateConsultatieByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[CONSULTATION] Attempting to update a consultatie by ID.")
	vars := mux.Vars(r)

	log.Printf("Vars: %v", vars)

	id, err := primitive.ObjectIDFromHex(vars[utils.UPDATE_CONSULTATIE_BY_ID_PARAMETER])
	if err != nil {
		response := models.ResponseData{
			Error: "Invalid consultatie ID",
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	consultatie := r.Context().Value(utils.DECODED_CONSULTATION).(*models.Consultatie)
	consultatie.IDConsultatie = id

	// Ensure a database operation doesn't take longer than utils.REQUEST_TIMEOUT_DURATION seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	// Use cController.DbConn to update the consultatie by ID in the database
	rowsAffected, err := cController.DbConn.UpdateConsultatieByID(ctx, consultatie)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[CONSULTATION] Failed to update consultatie by ID: %s\n", errMsg)
		response := models.ResponseData{
			Error: errMsg,
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Check if the consultatie exists and was updated
	if rowsAffected == 0 {
		response := models.ResponseData{
			Error: "Consultatie not found",
		}
		utils.RespondWithJSON(w, http.StatusNotFound, response)
		return
	}

	log.Printf("[CONSULTATION] Successfully updated consultatie %d", consultatie.IDConsultatie)
	response := models.ResponseData{
		Message: "Consultatie updated",
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}
