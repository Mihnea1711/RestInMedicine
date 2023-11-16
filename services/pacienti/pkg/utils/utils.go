package utils

import (
	"encoding/json"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

var (
	EmailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	PhoneRegex = regexp.MustCompile(`^(07[0-9]{8}|\+407[0-9]{8})$`)
)

// RespondWithJSON handles responding to HTTP requests with JSON.
func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	if payload == nil {
		respondWithError(w, http.StatusInternalServerError, "Empty response body")
		return
	}

	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[PATIENT] Error marshaling JSON: %s", err)
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	writeJSONResponse(w, status, response)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	errorResponse := map[string]string{"error": message}
	response, err := json.Marshal(errorResponse)
	if err != nil {
		log.Printf("[PATIENT] Error marshaling error response JSON: %s", err)
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

// Extract pagination parameters from the request
func ExtractPaginationParams(r *http.Request) (int, int) {
	limitStr := r.URL.Query().Get("limit")
	pageStr := r.URL.Query().Get("page")

	var limit, page int
	var err error

	// If limit parameter is provided, try to parse it
	if limitStr != "" {
		limit, err = strconv.Atoi(limitStr)
		if err != nil || limit < 1 {
			limit = DEFAULT_PAGINATION_LIMIT // Use a default limit value
		}
	} else {
		limit = MAX_PAGINATION_LIMIT // Set it to a maximum value to indicate no limit
	}

	// If page parameter is provided, try to parse it
	if pageStr != "" {
		page, err = strconv.Atoi(pageStr)
		if err != nil || page < 1 {
			page = DEFAULT_PAGINATION_PAGE // Use a default page value
		}
	} else {
		page = DEFAULT_PAGINATION_PAGE // Set it to the first page if not provided
	}

	return limit, page
}
