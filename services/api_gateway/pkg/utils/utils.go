package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
)

// RespondWithJSON handles responding to HTTP requests with JSON.
func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	if payload == nil {
		respondWithError(w, http.StatusInternalServerError, "Empty response body")
		return
	}

	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[GATEWAY] Error marshaling JSON: %s", err)
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	writeJSONResponse(w, status, response)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	errorResponse := map[string]string{"error": message}
	response, err := json.Marshal(errorResponse)
	if err != nil {
		log.Printf("[GATEWAY] Error marshaling error response JSON: %s", err)
		writeJSONResponse(w, http.StatusInternalServerError, []byte(`{"error":"Internal Server Error"}`))
		return
	}

	writeJSONResponse(w, status, response)
}

func writeJSONResponse(w http.ResponseWriter, status int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(response)
	w.Write([]byte("\n"))
}

func SendMessageResponse(w http.ResponseWriter, status int, message string, payload interface{}) {
	responseData := models.ResponseData{
		Message: message,
		Payload: payload,
	}
	RespondWithJSON(w, status, responseData)
}

func SendErrorResponse(w http.ResponseWriter, status int, message string, errString string) {
	responseData := models.ResponseData{
		Message: message,
		Error:   errString,
	}
	RespondWithJSON(w, status, responseData)
}
