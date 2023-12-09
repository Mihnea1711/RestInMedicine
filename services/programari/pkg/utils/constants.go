package utils

import "github.com/mihnea1711/POS_Project/services/programari/internal/models"

type contextKey string

const (
	StatusScheduled  models.StatusProgramare = "scheduled"
	StatusConfirmed  models.StatusProgramare = "confirmed"
	StatusNotPresent models.StatusProgramare = "not_present"
	StatusCanceled   models.StatusProgramare = "canceled"
	StatusHonored    models.StatusProgramare = "honored"
)

var ValidStatus = [...]models.StatusProgramare{StatusScheduled, StatusConfirmed, StatusNotPresent, StatusCanceled, StatusHonored}

const CONFIG_PATH = "configs/config.yaml"

const DECODED_APPOINTMENT contextKey = "decodedAppointment"

const (
	LIMITER_REQUESTS_ALLOWED  = 10
	LIMITER_MINUTE_MULTIPLIER = 1
)

const (
	DEFAULT_PAGINATION_LIMIT = 20
	MAX_PAGINATION_LIMIT     = 50
	DEFAULT_PAGINATION_PAGE  = 1
)

const DB_REQ_TIMEOUT_SEC_MULTIPLIER = 5
const CLEAR_DB_RESOURCES_TIMEOUT = 10

const TIME_PARSE_SYNTAX = "2006-01-02"

const (
	// Endpoints
	CREATE_APPOINTMENT_ENDPOINT       = "/appointments"
	FETCH_ALL_APPOINTMENTS_ENDPOINT   = "/appointments"
	FETCH_APPOINTMENT_BY_ID_ENDPOINT  = "/appointments/{" + FETCH_APPOINTMENT_BY_ID_PARAMETER + "}"
	UPDATE_APPOINTMENT_BY_ID_ENDPOINT = "/appointments/{" + UPDATE_APPOINTMENT_BY_ID_PARAMETER + "}"
	DELETE_APPOINTMENT_BY_ID_ENDPOINT = "/appointments/{" + DELETE_APPOINTMENT_BY_ID_PARAMETER + "}"

	HEALTH_CHECK_ENDPOINT = "/appointments/health-check"
)

const (
	// Parameters
	FETCH_APPOINTMENT_BY_ID_PARAMETER  = "appointmentID"
	UPDATE_APPOINTMENT_BY_ID_PARAMETER = "appointmentID"
	DELETE_APPOINTMENT_BY_ID_PARAMETER = "appointmentID"

	QUERY_PATIENT_ID = "patientID"
	QUERY_DOCTOR_ID  = "doctorID"
	QUERY_DATE       = "date"
	QUERY_STATUS     = "status"
	QUERY_PAGE       = "page"
	QUERY_LIMIT      = "limit"
)

const (
	DatabaseName         = "pdp_db"
	AppointmentTableName = "appointment"
	ColumnIDProgramare   = "id_appointment"
	ColumnIDPatient      = "id_patient"
	ColumnIDDoctor       = "id_doctor"
	ColumnDate           = "date"
	ColumnStatus         = "status"
)

const MySQLDuplicateEntryErrorCode = 1062
