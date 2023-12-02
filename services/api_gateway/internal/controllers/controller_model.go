package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/gateway/idm"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

type GatewayController struct {
	IDMClient idm.IDMClient
}

func (gc *GatewayController) redirectRequestBody(ctx context.Context, methodType, host, endpoint string, port int, data interface{}) (*models.ResponseData, int, error) {
	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("[GATEWAY] Error marshaling JSON data: %v", err)
		return nil, http.StatusInternalServerError, err
	}

	// Create a new request
	request, err := http.NewRequestWithContext(ctx, methodType, fmt.Sprintf("http://%s:%d%s", host, port, endpoint), bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("[GATEWAY] Error creating HTTP request: %v", err)
		return nil, http.StatusInternalServerError, err
	}

	// Set the Content-Type header
	request.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and make the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("[GATEWAY] Error making HTTP request: %v", err)
		return nil, http.StatusInternalServerError, err
	}

	// Close the response body explicitly after decoding
	defer func() {
		if cerr := response.Body.Close(); cerr != nil {
			log.Printf("[GATEWAY] Error closing response body: %v", cerr)
		}
	}()

	decodedResponse, err := utils.DecodeSanitizedResponse(response)
	if err != nil {
		log.Printf("[GATEWAY] Error decoding HTML encoded request: %v", err)
		return nil, http.StatusInternalServerError, err
	}

	return decodedResponse, response.StatusCode, nil
}
