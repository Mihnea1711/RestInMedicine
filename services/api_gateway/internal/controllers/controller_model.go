package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/gateway/idm"
)

type GatewayController struct {
	IDMClient idm.IDMClient
}

func (c *GatewayController) redirectRequestBody(endpoint string, port int, data interface{}) (interface{}, error) {
	// Convert data to JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	// Create a new request
	request, err := http.NewRequest("POST", fmt.Sprintf("http://localhost:%v%v", port, endpoint), bytes.NewBuffer(jsonData))
	if err != nil {
		return nil, err
	}

	// Set the Content-Type header
	request.Header.Set("Content-Type", "application/json")

	// Create an HTTP client and make the request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Check the response status code
	if response.StatusCode != http.StatusOK {
		// Handle the error (e.g., return a specific error or log it)
		return nil, err
	}

	// Optionally, read the response body
	var responseBody interface{}
	if err := json.NewDecoder(response.Body).Decode(&responseBody); err != nil {
		return nil, err
	}

	return responseBody, nil
}
