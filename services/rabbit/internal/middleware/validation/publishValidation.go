package validation

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func ValidateUserID(r *http.Request) (int, error) {
	// Decode the JSON request body
	var requestBody map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&requestBody)
	if err != nil {
		return 0, fmt.Errorf("Invalid JSON format: %v", err)
	}

	// Extract the userID from the request body
	userIDInterface, exists := requestBody["userID"]
	if !exists {
		return 0, fmt.Errorf("User ID is missing in the request body")
	}

	// Convert userID to integer
	userID, ok := userIDInterface.(float64)
	if !ok {
		return 0, fmt.Errorf("User ID must be a number")
	}

	// Convert userID to int
	userIDInt := int(userID)

	// Your additional validation logic goes here...

	// Example: Check if userID is within a valid range
	if userIDInt < 1 || userIDInt > 1000 {
		return 0, fmt.Errorf("Invalid userID range")
	}

	// Return the validated userID
	return userIDInt, nil
}
