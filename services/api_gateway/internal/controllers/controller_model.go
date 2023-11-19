package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/gateway/idm"
)

type GatewayController struct {
	IDMClient idm.IDMClient
}

func (c *GatewayController) redirectRequestBody(ctx context.Context, methodType string, endpoint string, port int, data interface{}) (*http.Response, error) {
	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Printf("[GATEWAY] Error marshaling JSON data: %v", err)
		return nil, err
	}

	// Create a new request
	request, err := http.NewRequestWithContext(ctx, methodType, fmt.Sprintf("http://localhost:%v%v", port, endpoint), bytes.NewBuffer(jsonData))
	if err != nil {
		log.Printf("[GATEWAY] Error creating HTTP request: %v", err)
		return nil, err
	}

	// Set the Content-Type header
	request.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and make the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("[GATEWAY] Error making HTTP request: %v", err)
		return nil, err
	}

	log.Printf("[GATEWAY] Request successful")
	return response, nil
}
