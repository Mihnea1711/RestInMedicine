package utils

type contextKey string

const CONFIG_PATH = "configs/config.yaml"

const DECODED_PATIENT contextKey = "decodedPatient"
const DECODED_PATIENT_ACTIVITY contextKey = "decodedPatientActivity"

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

const (
	// Endpoints
	HEALTH_CHECK_ENDPOINT              = "/patients/health-check"
	CREATE_PATIENT_ENDPOINT            = "/patients"
	FETCH_ALL_PATIENTS_ENDPOINT        = "/patients"
	FETCH_PATIENT_BY_ID_ENDPOINT       = "/patients/{" + FETCH_PATIENT_BY_ID_PARAMETER + "}"
	FETCH_PATIENT_BY_EMAIL_ENDPOINT    = "/patients/email/{" + FETCH_PATIENT_BY_EMAIL_PARAMETER + "}"
	FETCH_PATIENT_BY_USER_ID_ENDPOINT  = "/patients/users/{" + FETCH_PATIENT_BY_USER_ID_PARAMETER + "}"
	UPDATE_PATIENT_BY_ID_ENDPOINT      = "/patients/{" + UPDATE_PATIENT_BY_ID_PARAMETER + "}"
	DELETE_PATIENT_BY_ID_ENDPOINT      = "/patients/{" + DELETE_PATIENT_BY_ID_PARAMETER + "}"
	DELETE_PATIENT_BY_USER_ID_ENDPOINT = "/patients/users/{" + DELETE_PATIENT_BY_USER_ID_PARAMETER + "}"
	TOGGLE_PATIENT_ACTIVITY_ENDPOINT   = "/patients/{" + PATCH_PATIENT_BY_ID_PARAMETER + "}"

	// Parameters
	FETCH_PATIENT_BY_ID_PARAMETER       = "patientID"
	FETCH_PATIENT_BY_EMAIL_PARAMETER    = "patientEmail"
	FETCH_PATIENT_BY_USER_ID_PARAMETER  = "patientID"
	UPDATE_PATIENT_BY_ID_PARAMETER      = "patientID"
	PATCH_PATIENT_BY_ID_PARAMETER       = "patientID"
	DELETE_PATIENT_BY_ID_PARAMETER      = "patientID"
	DELETE_PATIENT_BY_USER_ID_PARAMETER = "patientUserID"

	QUERY_IS_ACIVE = "isActive"
	QUERY_PAGE     = "page"
	QUERY_LIMIT    = "limit"
)

const (
	DatabaseName     = "pdp_db"
	PatientTableName = "patient"

	ColumnIDPatient   = "id_patient"
	ColumnIDUser      = "id_user"
	ColumnFirstName   = "first_name"
	ColumnSecondName  = "second_name"
	ColumnEmail       = "email"
	ColumnPhoneNumber = "phone_number"
	ColumnCNP         = "cnp"
	ColumnBirthDay    = "birth_day"
	ColumnIsActive    = "is_active"
)

const CNP_DATE_FORMAT = "060102"

const (
	MySQLDuplicateEntryErrorCode = 1062
)
