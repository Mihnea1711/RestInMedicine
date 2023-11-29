package utils

import "github.com/mihnea1711/POS_Project/services/rabbit/internal/models"

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
	IDM_PORT = 8081
	IDM_HOST = "localhost"

	PATIENT_PORT = 8082
	PATIENT_HOST = "localhost"

	DOCTOR_PORT = 8083
	DOCTOR_HOST = "localhost"

	APPOINTMENT_PORT = 8084
	APPOINTMENT_HOST = "localhost"

	CONSULTATION_PORT = 8085
	CONSULTATION_HOST = "localhost"
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

const (
	PREPARE_PATIENT_ENDPOINT  = "/patients/transaction/prepare"
	COMMIT_PATIENT_ENDPOINT   = "/patients/transaction/commit"
	ABORT_PATIENT_ENDPOINT    = "/patients/transaction/abort"
	ROLLBACK_PATIENT_ENDPOINT = "/patients/transaction/rollback"

	PREPARE_DOCTOR_ENDPOINT  = "/doctors/transaction/prepare"
	COMMIT_DOCTOR_ENDPOINT   = "/doctors/transaction/commit"
	ABORT_DOCTOR_ENDPOINT    = "/doctors/transaction/abort"
	ROLLBACK_DOCTOR_ENDPOINT = "/doctors/transaction/rollback"

	PREPARE_APPOINTMENT_ENDPOINT  = "/appointments/transaction/prepare"
	COMMIT_APPOINTMENT_ENDPOINT   = "/appointments/transaction/commit"
	ABORT_APPOINTMENT_ENDPOINT    = "/appointments/transaction/abort"
	ROLLBACK_APPOINTMENT_ENDPOINT = "/appointments/transaction/rollback"

	PREPARE_CONSULTATION_ENDPOINT  = "/consultation/transaction/prepare"
	COMMIT_CONSULTATION_ENDOCTOR   = "/consultation/transaction/commit"
	ABORT_CONSULTATION_ENDPOINT    = "/consultation/transaction/abort"
	ROLLBACK_CONSULTATION_ENDPOINT = "/consultation/transaction/rollback"
)

const (
	IDM          models.ParticipantType = "IDM"
	PATIENT      models.ParticipantType = "PATIENT"
	DOCTOR       models.ParticipantType = "DOCTOR"
	APPOINTMENT  models.ParticipantType = "APPOINTMENT"
	CONSULTATION models.ParticipantType = "CONSULTATION"
)

const REQUEST_TIMEOUT_MULTIPLIER = 5

const ()
