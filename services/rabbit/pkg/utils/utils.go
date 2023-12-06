package utils

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models"
)

// DecodeHTML decodes HTML-encoded JSON to a struct
func DecodeHTML(s string, v interface{}) error {
	decoded := html.UnescapeString(s)
	return json.Unmarshal([]byte(decoded), v)
}

func DecodeSanitizedResponse(response *http.Response) (*models.ResponseData, error) {
	// Read the HTML-encoded JSON string from the response body
	htmlEncodedJSON, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("[RABBIT] Error reading response body: %v", err)
		return nil, err
	}

	// Decode HTML-encoded JSON string to ResponseData
	var decodedResponse models.ResponseData
	if err := DecodeHTML(string(htmlEncodedJSON), &decodedResponse); err != nil {
		log.Printf("[RABBIT] Error decoding HTML-encoded JSON: %v", err)
		return nil, err
	}

	return &decodedResponse, nil
}

func MakeRequest(ctx context.Context, methodType, host, endpoint string, port int, data interface{}) (*models.ResponseData, int, error) {
	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("[RABBIT] Error marshaling JSON data: %v", err)
		return nil, http.StatusInternalServerError, err
	}

	// Create a new request
	request, err := http.NewRequestWithContext(ctx, methodType, fmt.Sprintf("http://%s:%d%s", host, port, endpoint), bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("[RABBIT] Error creating HTTP request: %v", err)
		return nil, http.StatusInternalServerError, err
	}

	// Set the Content-Type header
	request.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and make the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("[RABBIT] Error making HTTP request: %v", err)
		return nil, http.StatusInternalServerError, err
	}

	// Close the response body explicitly after decoding
	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			log.Printf("[RABBIT] Error closing response body: %v", cerr)
		}
	}()

	decodedResponse, err := DecodeSanitizedResponse(response)
	if err != nil {
		log.Printf("[RABBIT] Error decoding HTML encoded request: %v", err)
		return nil, http.StatusInternalServerError, err
	}

	return decodedResponse, response.StatusCode, nil
}

func StartTransaction() string {
	// Generate a unique transaction ID, e.g., using UUID
	transactionID := uuid.New().String()

	// Log or store the transaction ID for future reference

	return transactionID
}
