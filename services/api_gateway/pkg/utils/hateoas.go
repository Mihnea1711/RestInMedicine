package utils

import (
	"strings"

	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
)

var HealthEndpoints = []models.LinkData{
	{FieldName: "health-check", EndpointData: models.EndpointData{Endpoint: CHECK_HEALTH_ENDPOINT, Method: "GET"}},
}

var UserEndpoints = []models.LinkData{
	{FieldName: "register", EndpointData: models.EndpointData{Endpoint: REGISTER_USER_ENDPOINT, Method: "POST"}},
	{FieldName: "login", EndpointData: models.EndpointData{Endpoint: LOGIN_USER_ENDPOINT, Method: "POST"}},
	{FieldName: "getAll", EndpointData: models.EndpointData{Endpoint: GET_ALL_USERS_ENDPOINT, Method: "GET"}},
	{FieldName: "getById", EndpointData: models.EndpointData{Endpoint: GET_USER_BY_ID_ENDPOINT, Method: "GET"}},
	{FieldName: "updateById", EndpointData: models.EndpointData{Endpoint: UPDATE_USER_BY_ID_ENDPOINT, Method: "PUT"}},
	{FieldName: "deleteById", EndpointData: models.EndpointData{Endpoint: DELETE_USER_BY_ID_ENDPOINT, Method: "DELETE"}},
	{FieldName: "updatePassword", EndpointData: models.EndpointData{Endpoint: UPDATE_PASSWORD_ENDPOINT, Method: "PUT"}},
	{FieldName: "updateRole", EndpointData: models.EndpointData{Endpoint: UPDATE_ROLE_ENDPOINT, Method: "PUT"}},
	{FieldName: "addToBlacklist", EndpointData: models.EndpointData{Endpoint: ADD_TO_BLACKLIST_ENDPOINT, Method: "POST"}},
	{FieldName: "checkBlacklist", EndpointData: models.EndpointData{Endpoint: CHECK_BLACKLIST_ENDPOINT, Method: "GET"}},
	{FieldName: "deleteFromBlacklist", EndpointData: models.EndpointData{Endpoint: DELETE_FROM_BLACKLIST_ENDPOINT, Method: "DELETE"}},
}

var PatientEndpoints = []models.LinkData{
	{FieldName: "create", EndpointData: models.EndpointData{Endpoint: CREATE_PATIENT_ENDPOINT, Method: "POST"}},
	{FieldName: "getAll", EndpointData: models.EndpointData{Endpoint: GET_ALL_PATIENTS_ENDPOINT, Method: "GET"}},
	{FieldName: "getById", EndpointData: models.EndpointData{Endpoint: GET_PATIENT_BY_ID_ENDPOINT, Method: "GET"}},
	{FieldName: "getByEmail", EndpointData: models.EndpointData{Endpoint: GET_PATIENT_BY_EMAIL_ENDPOINT, Method: "GET"}},
	{FieldName: "getByUserId", EndpointData: models.EndpointData{Endpoint: GET_PATIENT_BY_USER_ID_ENDPOINT, Method: "GET"}},
	{FieldName: "updateById", EndpointData: models.EndpointData{Endpoint: UPDATE_PATIENT_BY_ID_ENDPOINT, Method: "PUT"}},
	{FieldName: "deleteById", EndpointData: models.EndpointData{Endpoint: DELETE_PATIENT_BY_ID_ENDPOINT, Method: "DELETE"}},
	{FieldName: "toggleActivity", EndpointData: models.EndpointData{Endpoint: TOGGLE_PATIENT_ACTIVITY_ENDPOINT, Method: "POST"}},
}

var DoctorEndpoints = []models.LinkData{
	{FieldName: "create", EndpointData: models.EndpointData{Endpoint: CREATE_DOCTOR_ENDPOINT, Method: "POST"}},
	{FieldName: "getAll", EndpointData: models.EndpointData{Endpoint: GET_ALL_DOCTORS_ENDPOINT, Method: "GET"}},
	{FieldName: "getById", EndpointData: models.EndpointData{Endpoint: GET_DOCTOR_BY_ID_ENDPOINT, Method: "GET"}},
	{FieldName: "getByEmail", EndpointData: models.EndpointData{Endpoint: GET_DOCTOR_BY_EMAIL_ENDPOINT, Method: "GET"}},
	{FieldName: "getByUserId", EndpointData: models.EndpointData{Endpoint: GET_DOCTOR_BY_USER_ID_ENDPOINT, Method: "GET"}},
	{FieldName: "updateById", EndpointData: models.EndpointData{Endpoint: UPDATE_DOCTOR_BY_ID_ENDPOINT, Method: "PUT"}},
	{FieldName: "deleteById", EndpointData: models.EndpointData{Endpoint: DELETE_DOCTOR_BY_ID_ENDPOINT, Method: "DELETE"}},
	{FieldName: "toggleActivity", EndpointData: models.EndpointData{Endpoint: TOGGLE_DOCTOR_ACTIVITY_ENDPOINT, Method: "POST"}},
}

var AppointmentEndpoints = []models.LinkData{
	{FieldName: "create", EndpointData: models.EndpointData{Endpoint: CREATE_APPOINTMENT_ENDPOINT, Method: "POST"}},
	{FieldName: "getAll", EndpointData: models.EndpointData{Endpoint: GET_ALL_APPOINTMENTS_ENDPOINT, Method: "GET"}},
	{FieldName: "getById", EndpointData: models.EndpointData{Endpoint: GET_APPOINTMENT_BY_ID_ENDPOINT, Method: "GET"}},
	{FieldName: "getByDoctorId", EndpointData: models.EndpointData{Endpoint: GET_APPOINTMENTS_BY_DOCTOR_ID_ENDPOINT, Method: "GET"}},
	{FieldName: "getByPatientId", EndpointData: models.EndpointData{Endpoint: GET_APPOINTMENTS_BY_PATIENT_ID_ENDPOINT, Method: "GET"}},
	{FieldName: "getByDate", EndpointData: models.EndpointData{Endpoint: GET_APPOINTMENTS_BY_DATE_ENDPOINT, Method: "GET"}},
	{FieldName: "getByStatus", EndpointData: models.EndpointData{Endpoint: GET_APPOINTMENTS_BY_STATUS_ENDPOINT, Method: "GET"}},
	{FieldName: "updateById", EndpointData: models.EndpointData{Endpoint: UPDATE_APPOINTMENT_BY_ID_ENDPOINT, Method: "PUT"}},
	{FieldName: "deleteById", EndpointData: models.EndpointData{Endpoint: DELETE_APPOINTMENT_BY_ID_ENDPOINT, Method: "DELETE"}},
}

var ConsultationEndpoints = []models.LinkData{
	{FieldName: "create", EndpointData: models.EndpointData{Endpoint: CREATE_CONSULTATION_ENDPOINT, Method: "POST"}},
	{FieldName: "getAll", EndpointData: models.EndpointData{Endpoint: GET_ALL_CONSULTATIONS_ENDPOINT, Method: "GET"}},
	{FieldName: "getByDoctorId", EndpointData: models.EndpointData{Endpoint: GET_CONSULTATION_BY_DOCTOR_ID_ENDPOINT, Method: "GET"}},
	{FieldName: "getByPatientId", EndpointData: models.EndpointData{Endpoint: GET_CONSULTATION_BY_PATIENT_ID_ENDPOINT, Method: "GET"}},
	{FieldName: "getByDate", EndpointData: models.EndpointData{Endpoint: GET_CONSULTATION_BY_DATE_ENDPOINT, Method: "GET"}},
	{FieldName: "getById", EndpointData: models.EndpointData{Endpoint: GET_CONSULTATION_BY_ID_ENDPOINT, Method: "GET"}},
	{FieldName: "updateById", EndpointData: models.EndpointData{Endpoint: UPDATE_CONSULTATION_BY_ID_ENDPOINT, Method: "PUT"}},
	{FieldName: "deleteById", EndpointData: models.EndpointData{Endpoint: DELETE_CONSULTATION_BY_ID_ENDPOINT, Method: "DELETE"}},
}

var AllEndpointsLinks = [][]models.LinkData{
	HealthEndpoints,
	UserEndpoints,
	PatientEndpoints,
	DoctorEndpoints,
	AppointmentEndpoints,
	ConsultationEndpoints,
}

func findAdjacentEndpoints(inputEndpoint, inputMethod string) []models.LinkData {
	var result []models.LinkData

	for _, endpointCategory := range AllEndpointsLinks {
		for _, endpoint := range endpointCategory {
			if strings.HasPrefix(endpoint.EndpointData.Endpoint, inputEndpoint) && (endpoint.EndpointData.Endpoint != inputEndpoint || endpoint.EndpointData.Method != inputMethod) {
				linkData := models.LinkData{
					FieldName:    endpoint.FieldName,
					EndpointData: models.EndpointData{Endpoint: endpoint.EndpointData.Endpoint, Method: endpoint.EndpointData.Method},
				}
				result = append(result, linkData)
			}
		}
	}
	return result
}

func getParentEndpoint(inputEndpoint string) string {
	// Trim any trailing slashes to ensure consistency
	inputEndpoint = strings.TrimSuffix(inputEndpoint, "/")

	// Find the last index of "/"
	lastSlashIndex := strings.LastIndex(inputEndpoint, "/")

	if lastSlashIndex == -1 {
		// No "/" found, the input is already the root
		return ""
	}

	// Extract the parent endpoint
	parentEndpoint := inputEndpoint[:lastSlashIndex]

	return parentEndpoint
}

func GetHateoasData(inputEndpoint, inputMethod string) []models.LinkData {
	var matchingEndpoints []models.LinkData

	matchingEndpoints = append(matchingEndpoints, models.LinkData{FieldName: "self", EndpointData: models.EndpointData{Endpoint: inputEndpoint, Method: inputMethod}})
	matchingEndpoints = append(matchingEndpoints, models.LinkData{FieldName: "parent", EndpointData: models.EndpointData{Endpoint: getParentEndpoint(inputEndpoint)}})
	matchingEndpoints = append(matchingEndpoints, findAdjacentEndpoints(inputEndpoint, inputMethod)...)

	return matchingEndpoints
}