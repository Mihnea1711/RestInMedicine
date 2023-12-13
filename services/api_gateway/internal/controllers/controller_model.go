package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/mihnea1711/POS_Project/services/gateway/idm"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/middleware/authorization"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

type GatewayController struct {
	IDMClient idm.IDMClient
}

func (gc *GatewayController) redirectRequestBody(ctx context.Context, methodType, host, endpoint string, port int, data interface{}) (*models.ResponseDataWrapper, int, error) {
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
	// Set the header for further use in the controllers
	decodedResponse.Header = &response.Header

	return decodedResponse, response.StatusCode, nil
}

func (gc *GatewayController) GenerateTargetURL(r *http.Request, baseEndpoint string) (string, error) {
	claims := r.Context().Value(utils.JWT_CLAIMS_CONTEXT_KEY).(*authorization.MyCustomClaims)

	// Get the userID from the claims
	userID, err := claims.GetSubject()
	if err != nil {
		return "", fmt.Errorf("failed to get user ID from claims: %v", err)
	}

	// Initialize the targetURL with the base endpoint
	targetURL := fmt.Sprintf("%s?", baseEndpoint)

	// Check the role in the claims and append appropriate query parameter
	switch claims.Role {
	case utils.PATIENT_ROLE:
		// Patients can only view their own data
		// get patient id from user id
		result, status, err := gc.redirectRequestBody(r.Context(), utils.GET, utils.PATIENT_HOST, fmt.Sprintf("%s/%s", utils.PATIENT_FETCH_PATIENT_BY_USER_ID_ENDPOINT, userID), utils.PATIENT_PORT, nil)
		if status != http.StatusOK {
			return "", fmt.Errorf("failed to redirect to patient module to get patientID: %v", err)
		}

		// Assuming result.Payload is of type interface{}
		var payloadPatient models.PatientData

		// Convert the payload to JSON
		payloadJSON, err := json.Marshal(result.Payload)
		if err != nil {
			return "", fmt.Errorf("failed to marshal result payload to JSON: %v", err)
		}

		// Unmarshal the JSON into the PatientData struct
		if err := json.Unmarshal(payloadJSON, &payloadPatient); err != nil {
			return "", fmt.Errorf("failed to unmarshal JSON into PatientData: %v", err)
		}

		targetURL += utils.AppendQueryParam(r.URL.RawQuery, utils.QUERY_ID_PATIENT, strconv.Itoa(payloadPatient.IDPatient))
	case utils.DOCTOR_ROLE:
		// Doctors can only view their own data

		// get patient id from user id
		result, status, err := gc.redirectRequestBody(r.Context(), utils.GET, utils.PATIENT_HOST, fmt.Sprintf("%s/%s", utils.PATIENT_FETCH_PATIENT_BY_USER_ID_ENDPOINT, userID), utils.PATIENT_PORT, nil)
		if status != http.StatusOK {
			return "", fmt.Errorf("failed to redirect to patient module to get patientID: %v", err)
		}

		// Assuming result.Payload is of type interface{}
		var payloadDoctor models.DoctorData

		// Convert the payload to JSON
		payloadJSON, err := json.Marshal(result.Payload)
		if err != nil {
			return "", fmt.Errorf("failed to marshal result payload to JSON: %v", err)
		}

		// Unmarshal the JSON into the PatientData struct
		if err := json.Unmarshal(payloadJSON, &payloadDoctor); err != nil {
			return "", fmt.Errorf("failed to unmarshal JSON into PatientData: %v", err)
		}

		targetURL += utils.AppendQueryParam(r.URL.RawQuery, utils.QUERY_ID_DOCTOR, strconv.Itoa(payloadDoctor.IDDoctor))
		// Add more cases if needed for other roles
	}

	return targetURL, nil
}

func (gc *GatewayController) GenerateTargetURLPatients(r *http.Request, baseEndpoint string) (string, error) {
	claims := r.Context().Value(utils.JWT_CLAIMS_CONTEXT_KEY).(*authorization.MyCustomClaims)

	// Get the userID from the claims
	userID, err := claims.GetSubject()
	if err != nil {
		return "", fmt.Errorf("failed to get user ID from claims: %v", err)
	}

	// Initialize the targetURL with the base endpoint
	targetURL := fmt.Sprintf("%s?", baseEndpoint)

	// Check the role in the claims and append appropriate query parameter
	switch claims.Role {
	case utils.DOCTOR_ROLE:
		// Doctors can only view their own data

		// get patient id from user id
		result, status, err := gc.redirectRequestBody(r.Context(), utils.GET, utils.PATIENT_HOST, fmt.Sprintf("%s/%s", utils.PATIENT_FETCH_PATIENT_BY_USER_ID_ENDPOINT, userID), utils.PATIENT_PORT, nil)
		if status != http.StatusOK {
			return "", fmt.Errorf("failed to redirect to patient module to get patientID: %v", err)
		}

		// Assuming result.Payload is of type interface{}
		var payloadDoctor models.DoctorData

		// Convert the payload to JSON
		payloadJSON, err := json.Marshal(result.Payload)
		if err != nil {
			return "", fmt.Errorf("failed to marshal result payload to JSON: %v", err)
		}

		// Unmarshal the JSON into the PatientData struct
		if err := json.Unmarshal(payloadJSON, &payloadDoctor); err != nil {
			return "", fmt.Errorf("failed to unmarshal JSON into PatientData: %v", err)
		}

		targetURL += utils.AppendQueryParam(r.URL.RawQuery, utils.QUERY_ID_DOCTOR, strconv.Itoa(payloadDoctor.IDDoctor))
		targetURL += utils.AppendQueryParam(r.URL.RawQuery, utils.QUERY_IS_ACTIVE, "true")
		// Add more cases if needed for other roles
	}

	return targetURL, nil
}
