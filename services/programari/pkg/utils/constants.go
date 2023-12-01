package utils

import "github.com/mihnea1711/POS_Project/services/programari/internal/models"

type contextKey string

const (
	StatusProgramata   models.StatusProgramare = "programata"
	StatusConfirmata   models.StatusProgramare = "confirmata"
	StatusNeprezentata models.StatusProgramare = "neprezentata"
	StatusAnulata      models.StatusProgramare = "anulata"
	StatusOnorata      models.StatusProgramare = "onorata"
)

var ValidStatus = [...]models.StatusProgramare{StatusOnorata, StatusNeprezentata, StatusAnulata, StatusProgramata, StatusConfirmata}

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
	CREATE_APPOINTMENT_ENDPOINT               = "/appointments"                                                               // POST
	FETCH_ALL_APPOINTMENTS_ENDPOINT           = "/appointments"                                                               // GET
	FETCH_APPOINTMENT_BY_ID_ENDPOINT          = "/appointments/{" + FETCH_APPOINTMENT_BY_ID_PARAMETER + "}"                   // GET
	FETCH_APPOINTMENTS_BY_DOCTOR_ID_ENDPOINT  = "/appointments/doctors/{" + FETCH_APPOINTMENTS_BY_DOCTOR_ID_PARAMETER + "}"   // GET
	FETCH_APPOINTMENTS_BY_PACIENT_ID_ENDPOINT = "/appointments/patients/{" + FETCH_APPOINTMENTS_BY_PACIENT_ID_PARAMETER + "}" // GET
	FETCH_APPOINTMENTS_BY_DATE_ENDPOINT       = "/appointments/date/{" + FETCH_APPOINTMENTS_BY_DATE_PARAMETER + "}"           // GET
	FETCH_APPOINTMENTS_BY_STATUS_ENDPOINT     = "/appointments/status/{" + FETCH_APPOINTMENTS_BY_STATUS_PARAMETER + "}"       // GET
	UPDATE_APPOINTMENT_BY_ID_ENDPOINT         = "/appointments/{" + UPDATE_APPOINTMENT_BY_ID_PARAMETER + "}"                  // PUT
	DELETE_APPOINTMENT_BY_ID_ENDPOINT         = "/appointments/{" + DELETE_APPOINTMENT_BY_ID_PARAMETER + "}"                  // DELETE

	HEALTH_CHECK_ENDPOINT = "/appointments/health-check"
)

// Parameters
const (
	FETCH_APPOINTMENT_BY_ID_PARAMETER          = "appointment_id"
	FETCH_APPOINTMENTS_BY_DOCTOR_ID_PARAMETER  = "doctor_id"
	FETCH_APPOINTMENTS_BY_PACIENT_ID_PARAMETER = "pacient_id"
	FETCH_APPOINTMENTS_BY_DATE_PARAMETER       = "appointment_date"
	FETCH_APPOINTMENTS_BY_STATUS_PARAMETER     = "appointment_status"
	UPDATE_APPOINTMENT_BY_ID_PARAMETER         = "appointment_id"
	DELETE_APPOINTMENT_BY_ID_PARAMETER         = "appointment_id"
)

const (
	DatabaseName         = "pdp_db"
	AppointmentTableName = "appointment"
	ColumnIDProgramare   = "id_programare"
	ColumnIDPacient      = "id_patient"
	ColumnIDDoctor       = "id_doctor"
	ColumnDate           = "date"
	ColumnStatus         = "status"
)

const MySQLDuplicateEntryErrorCode = 1062
