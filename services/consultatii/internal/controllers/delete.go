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

// Delete a consultatie by ID
func (cController *ConsultatieController) DeleteConsultatieByID(w http.ResponseWriter, r *http.Request) {
	log.Printf("[CONSULTATION] Attempting to delete a consultatie by ID.")
	vars := mux.Vars(r)
	id, err := primitive.ObjectIDFromHex(vars[utils.DELETE_CONSULTATIE_BY_ID_PARAMETER])
	if err != nil {
		response := models.ResponseData{
			Error: "Invalid consultatie ID",
		}
		utils.RespondWithJSON(w, http.StatusBadRequest, response)
		return
	}

	// Ensure a database operation doesn't take longer than 5 seconds
	ctx, cancel := context.WithTimeout(r.Context(), utils.REQUEST_TIMEOUT_DURATION*time.Second)
	defer cancel()

	// Use cController.DbConn to delete the consultatie by ID from the database
	rowsAffected, err := cController.DbConn.DeleteConsultatieByID(ctx, id)
	if err != nil {
		errMsg := fmt.Sprintf("internal server error: %s", err)
		log.Printf("[CONSULTATION] Failed to delete consultatie by ID: %s\n", errMsg)
		response := models.ResponseData{
			Error: errMsg,
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	// Check if the consultatie exists and was deleted
	if rowsAffected == 0 {
		response := models.ResponseData{
			Error: "Consultatie not found",
		}
		utils.RespondWithJSON(w, http.StatusNotFound, response)
		return
	}

	log.Printf("[CONSULTATION] Successfully deleted consultatie %d", id)
	response := models.ResponseData{
		Message: "Consultatie deleted",
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}
