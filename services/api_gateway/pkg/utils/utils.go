package utils

import (
	"encoding/json"
	"html"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
)

// RespondWithJSON handles responding to HTTP requests with JSON.
func RespondWithJSON(w http.ResponseWriter, status int, payload interface{}) {
	if payload == nil {
		log.Println("[GATEWAY] RespondWithJSON: Empty response body")
		respondWithError(w, http.StatusInternalServerError, "Empty response body")
		return
	}

	response, err := json.Marshal(payload)
	if err != nil {
		log.Printf("[GATEWAY] RespondWithJSON: Error marshaling JSON: %s", err)
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	writeJSONResponse(w, status, response)
}

func respondWithError(w http.ResponseWriter, status int, message string) {
	errorResponse := map[string]string{"error": message}
	response, err := json.Marshal(errorResponse)
	if err != nil {
		log.Printf("[GATEWAY] RespondWithError: Error marshaling error response JSON: %s", err)
		writeJSONResponse(w, http.StatusInternalServerError, []byte(`{"error":"Internal Server Error"}`))
		return
	}

	writeJSONResponse(w, status, response)
}

func writeJSONResponse(w http.ResponseWriter, status int, response []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status) // for some reason this is seen as superfluous
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

func CheckNilResponse(w http.ResponseWriter, status int, errorMessage string, nilCheckFunc func() bool, nilCheckMessage string) {
	if nilCheckFunc() {
		log.Printf("[GATEWAY] %s", nilCheckMessage)
		SendErrorResponse(w, status, errorMessage, nilCheckMessage)
	}
}

// DecodeHTML decodes HTML-encoded JSON to a struct
func DecodeHTML(s string, v interface{}) error {
	decoded := html.UnescapeString(s)
	return json.Unmarshal([]byte(decoded), v)
}

func DecodeSanitizedResponse(response *http.Response) (*models.ResponseDataWrapper, error) {
	// Read the HTML-encoded JSON string from the response body
	htmlEncodedJSON, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("[GATEWAY] Error reading response body: %v", err)
		return nil, err
	}

	// Decode HTML-encoded JSON string to ResponseData
	var decodedResponse models.ResponseDataWrapper
	if err := DecodeHTML(string(htmlEncodedJSON), &decodedResponse); err != nil {
		log.Printf("[GATEWAY] Error decoding HTML-encoded JSON: %v", err)
		return nil, err
	}

	return &decodedResponse, nil
}

func AppendQueryParam(rawQuery, key, value string) string {
	// Parse the existing query parameters
	params, _ := url.ParseQuery(rawQuery)

	// Check if the key already exists
	if _, exists := params[key]; exists {
		// Key already exists, don't append it again
		return rawQuery
	}

	// Append the new key-value pair
	params.Add(key, value)

	// Encode the updated parameters
	return params.Encode()
}
