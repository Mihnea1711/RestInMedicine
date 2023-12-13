package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

func (gc *GatewayController) GetDocs(w http.ResponseWriter, r *http.Request) {
	// Read the embedded OpenAPI JSON file
	openapiData, err := os.ReadFile("openapi.json")
	if err != nil {
		log.Println("[GATEWAY] ERROR: Failed to read openapi.json file")
		response := models.ResponseData{
			Message: "Failed to retrieve OpenAPI documentation",
			Error:   "Error reading OpenAPI file: " + err.Error(),
		}
		utils.RespondWithJSON(w, http.StatusInternalServerError, response)
		return
	}

	log.Println("[GATEWAY] Successfully loaded openapi.json file")
	response := models.ResponseData{
		Message: "OpenAPI documentation retrieved successfully",
		Payload: openapiData,
	}
	utils.RespondWithJSON(w, http.StatusOK, response)
}
