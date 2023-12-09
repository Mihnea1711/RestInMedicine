package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

// RespondWithJSON handles responding to HTTP requests with JSON.
func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	if payload == nil {
		log.Println("[APPOINTMENT] RespondWithJSON: Empty response body")
		respondWithError(w, http.StatusInternalServerError, "Empty response body")
		return
	}

	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[APPOINTMENT] RespondWithJSON: Error marshaling JSON: %s", err)
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	writeJSONResponse(w, status, response)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	errorResponse := map[string]string{"error": message}
	response, err := json.Marshal(errorResponse)
	if err != nil {
		log.Printf("[APPOINTMENT] RespondWithError: Error marshaling error response JSON: %s", err)
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
	limitStr := r.URL.Query().Get(QUERY_LIMIT)
	pageStr := r.URL.Query().Get(QUERY_PAGE)

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

// ExtractFiltersFromRequest extracts query parameters from the request and constructs a map of filters.
func ExtractFiltersFromRequest(r *http.Request) (map[string]interface{}, error) {
	filters := make(map[string]interface{})

	// Check for unknown filters
	for key := range r.URL.Query() {
		if !isExpectedFilter(key) {
			log.Printf("[APPOINTMENT] ExtractFiltersFromRequest: Unknown filter: %s", key)
			return nil, fmt.Errorf("unknown filter: %s", key)
		}
	}

	// Parse query parameters
	patientID := r.URL.Query().Get(QUERY_PATIENT_ID)
	doctorID := r.URL.Query().Get(QUERY_DOCTOR_ID)
	date := r.URL.Query().Get(QUERY_DATE)
	status := r.URL.Query().Get(QUERY_STATUS)

	// Convert string values to appropriate types
	if patientID != "" {
		if id, err := strconv.Atoi(patientID); err == nil {
			filters[ColumnIDPatient] = id
		} else {
			log.Printf("[APPOINTMENT] ExtractFiltersFromRequest: Failed to parse patientID: %v", err)
			return nil, fmt.Errorf("invalid patientID: %v", err)
		}
	}
	if doctorID != "" {
		if id, err := strconv.Atoi(doctorID); err == nil {
			filters[ColumnIDDoctor] = id
		} else {
			log.Printf("[APPOINTMENT] ExtractFiltersFromRequest: Failed to parse doctorID: %v", err)
			return nil, fmt.Errorf("invalid doctorID: %v", err)
		}
	}
	if date != "" {
		if t, err := time.Parse(TIME_PARSE_SYNTAX, date); err == nil {
			filters[ColumnDate] = t
		} else {
			log.Printf("[APPOINTMENT] ExtractFiltersFromRequest: Failed to parse date: %v", err)
			return nil, fmt.Errorf("invalid date: %v", err)
		}
	}
	if status != "" {
		// Check if the provided status is valid
		found := false
		for _, validStatus := range ValidStatus {
			if string(validStatus) == status {
				filters[ColumnStatus] = validStatus
				found = true
				break
			}
		}
		if !found {
			log.Printf("[APPOINTMENT] ExtractFiltersFromRequest: Invalid status: %s", status)
			return nil, fmt.Errorf("invalid status: %s", status)
		}
	}

	return filters, nil
}

// isExpectedFilter checks if a filter name is one of the expected names.
func isExpectedFilter(filterName string) bool {
	expectedFilters := map[string]struct{}{
		QUERY_PATIENT_ID: {},
		QUERY_DOCTOR_ID:  {},
		QUERY_DATE:       {},
		QUERY_STATUS:     {},
		QUERY_PAGE:       {},
		QUERY_LIMIT:      {},
	}

	_, ok := expectedFilters[filterName]
	return ok
}
