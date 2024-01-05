package utils

import "github.com/mihnea1711/POS_Project/services/rabbit/internal/models"

type contextKey string

const CONFIG_PATH = "configs/config.yaml"
const CLEAR_DB_RESOURCES_TIMEOUT = 10
const RABBIT_CLOSE_TIMEOUT = 5

const (
	// QueueDirectionListen represents a queue for listening/consuming
	QueueDirectionListen string = "listen"
	// QueueDirectionPublish represents a queue for publishing
	QueueDirectionPublish string = "publish"
)

const JWT_CLAIMS_CONTEXT_KEY contextKey = "jwtClaims"

const (
	DELETE_QUEUE = "delete_queue"
)

const (
	IDM_HOST = "idm_app"
	IDM_PORT = 8081

	PATIENT_HOST = "patient_app"
	PATIENT_PORT = 8082

	DOCTOR_HOST = "doctor_app"
	DOCTOR_PORT = 8083

	APPOINTMENT_HOST = "appointment_app"
	APPOINTMENT_PORT = 8084

	CONSULTATION_HOST = "consultation_app"
	CONSULTATION_PORT = 8085
)

const (
	POST   = "POST"
	GET    = "GET"
	PUT    = "PUT"
	PATCH  = "PATCH"
	DELETE = "DELETE"
)

const (
	ADMIN_ROLE   = "admin"
	PATIENT_ROLE = "patient"
	DOCTOR_ROLE  = "doctor"
)

const (
	HEALTH_CHECK_ENDPOINT = "/api/rabbit/health-check"
	PUBLISH_ENDPOINT      = "/api/rabbit/publish"
)

const (
	PREPARE_PATIENT_ENDPOINT = "/patients/health-check"
	COMMIT_PATIENT_ENDPOINT  = "/patients"
	// ABORT_PATIENT_ENDPOINT    = "/patients/transaction/abort"
	// ROLLBACK_PATIENT_ENDPOINT = "/patients/transaction/rollback"

	PREPARE_DOCTOR_ENDPOINT = "/doctors/health-check"
	COMMIT_DOCTOR_ENDPOINT  = "/doctors"
	// ABORT_DOCTOR_ENDPOINT    = "/doctors/transaction/abort"
	// ROLLBACK_DOCTOR_ENDPOINT = "/doctors/transaction/rollback"
)

const (
	IDM     models.ParticipantType = "IDM"
	PATIENT models.ParticipantType = "PATIENT"
	DOCTOR  models.ParticipantType = "DOCTOR"
)

const REQUEST_TIMEOUT_MULTIPLIER = 5
