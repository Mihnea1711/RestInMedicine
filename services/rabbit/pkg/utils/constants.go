package utils

const CONFIG_PATH = "configs/config.yaml"
const CLEAR_DB_RESOURCES_TIMEOUT = 10
const RABBIT_CLOSE_TIMEOUT = 5

const (
	// QueueDirectionListen represents a queue for listening/consuming
	QueueDirectionListen string = "listen"
	// QueueDirectionPublish represents a queue for publishing
	QueueDirectionPublish string = "publish"
)

const (
	DELETE_QUEUE = "delete_queue"
)

const (
	// ports
	IDM_HOST          = "localhost"
	IDM_PORT          = 8081
	PATIENT_PORT      = 8082
	DOCTOR_PORT       = 8083
	APPOINTMENT_PORT  = 8084
	CONSULTATION_PORT = 8085
)

const (
	ADMIN_ROLE   = "admin"
	PATIENT_ROLE = "patient"
	DOCTOR_ROLE  = "doctor"
)

const (
	// endpoints
	DELETE_PATIENT_BY_USER_ID_ENDPOINT         = "/api/patients/users"         // followed by userID
	DELETE_DOCTOR_BY_USER_ID_ENDPOINT          = "/api/doctors/users"          // followed by userID
	DELETE_APPOINTMENT_BY_PATIENT_ID_ENDPOINT  = "/api/appointments/patients"  // followed by patientID
	DELETE_APPOINTMENT_BY_DOCTOR_ID_ENDPOINT   = "/api/appointments/doctors"   // followed by doctorID
	DELETE_CONSULTATION_BY_PATIENT_ID_ENDPOINT = "/api/consultations/patients" // followed by patientID
	DELETE_CONSULTATION_BY_DOCTOR_ID_ENDPOINT  = "/api/consultations/doctors"  // followed by doctorID

	HEALTH_CHECK_ENDPOINT = "/api/rabbit/health-check"
)
