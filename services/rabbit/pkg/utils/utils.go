package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"html"
	"io"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models"
)

func MakeRequest(ctx context.Context, methodType, host string, port int, endpoint string) (*models.ParticipantResponse, int, error) {
	// Create a new request
	request, err := http.NewRequestWithContext(ctx, methodType, fmt.Sprintf("http://%s:%v%v", host, port, endpoint), nil)
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

// DecodeHTML decodes HTML-encoded JSON to a struct
func DecodeHTML(s string, v interface{}) error {
	decoded := html.UnescapeString(s)
	return json.Unmarshal([]byte(decoded), v)
}

func DecodeSanitizedResponse(response *http.Response) (*models.ParticipantResponse, error) {
	// Read the HTML-encoded JSON string from the response body
	htmlEncodedJSON, err := io.ReadAll(response.Body)
	if err != nil {
		log.Printf("[RABBIT] Error reading response body: %v", err)
		return nil, err
	}

	// Decode HTML-encoded JSON string to ResponseData
	var decodedResponse models.ParticipantResponse
	if err := DecodeHTML(string(htmlEncodedJSON), &decodedResponse); err != nil {
		log.Printf("[RABBIT] Error decoding HTML-encoded JSON: %v", err)
		return nil, err
	}

	return &decodedResponse, nil
}
