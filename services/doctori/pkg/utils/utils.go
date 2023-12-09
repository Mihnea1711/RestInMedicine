package utils

import (
	"encoding/json"
	"fmt"
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
		log.Println("[DOCTOR] RespondWithJSON: Empty response body")
		respondWithError(w, http.StatusInternalServerError, "Empty response body")
		return
	}

	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[DOCTOR] RespondWithJSON: Error marshaling JSON: %s", err)
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	writeJSONResponse(w, status, response)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	errorResponse := map[string]string{"error": message}
	response, err := json.Marshal(errorResponse)
	if err != nil {
		log.Printf("[DOCTOR] RespondWithError: Error marshaling error response JSON: %s", err)
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

// ExtractFiltersFromRequest extracts query parameters from the request and constructs a map of filters.
func ExtractFiltersFromRequest(r *http.Request) (map[string]interface{}, error) {
	filters := make(map[string]interface{})

	// Check for unknown filters
	for key := range r.URL.Query() {
		if !isExpectedFilter(key) {
			log.Printf("[DOCTOR] ExtractFiltersFromRequest: Unknown filter: %s", key)
			return nil, fmt.Errorf("unknown filter: %s", key)
		}
	}

	// Parse query parameters
	isActiveStr := r.URL.Query().Get(QUERY_IS_ACIVE)
	firstName := r.URL.Query().Get(QUERY_FIRST_NAME)
	specialization := r.URL.Query().Get(QUERY_SPECIALIZATION)

	if isActiveStr != "" {
		isActive, err := strconv.ParseBool(isActiveStr)
		if err != nil {
			log.Printf("[DOCTOR] ExtractFiltersFromRequest: Failed to parse isActive: %v", err)
			return nil, fmt.Errorf("invalid isActive: %v", err)
		}
		filters[ColumnIsActive] = isActive
	}
	if firstName != "" {
		filters[ColumnFirstName] = firstName
	}
	if specialization != "" {
		// Check if the provided status is valid
		found := false
		for _, validSpecialization := range ValidSpecializations {
			if string(validSpecialization) == specialization {
				filters[ColumnSpecialization] = validSpecialization
				found = true
				break
			}
		}
		if !found {
			log.Printf("[APPOINTMENT] ExtractFiltersFromRequest: Invalid specialization: %s", specialization)
			return nil, fmt.Errorf("invalid specialization: %s", specialization)
		}
	}

	return filters, nil
}

// isExpectedFilter checks if a filter name is one of the expected names.
func isExpectedFilter(filterName string) bool {
	expectedFilters := map[string]struct{}{
		QUERY_IS_ACIVE:       {},
		QUERY_FIRST_NAME:     {},
		QUERY_SPECIALIZATION: {},
		QUERY_PAGE:           {},
		QUERY_LIMIT:          {},
	}

	_, ok := expectedFilters[filterName]
	return ok
}
