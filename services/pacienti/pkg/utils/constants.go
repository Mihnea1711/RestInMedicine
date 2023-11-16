package utils

type contextKey string

const CONFIG_PATH = "configs/config.yaml"

const DECODED_PATIENT contextKey = "decodedPatient"

const (
	LIMITER_REQUESTS_ALLOWED  = 10
	LIMITER_MINUTE_MULTIPLIER = 1
)

const DB_REQ_TIMEOUT_SEC_MULTIPLIER = 5

const (
	// Endpoints
	CREATE_PATIENT_ENDPOINT           = "/patients"
	FETCH_ALL_PATIENTS_ENDPOINT       = "/patients"
	FETCH_PATIENT_BY_ID_ENDPOINT      = "/patients/{" + FETCH_PATIENT_BY_ID_PARAMETER + "}"
	FETCH_PATIENT_BY_EMAIL_ENDPOINT   = "/patients/email/{" + FETCH_PATIENT_BY_EMAIL_PARAMETER + "}"
	FETCH_PATIENT_BY_USER_ID_ENDPOINT = "/patients/users/{" + FETCH_PATIENT_BY_USER_ID_PARAMETER + "}"
	UPDATE_PATIENT_BY_ID_ENDPOINT     = "/patients/{" + UPDATE_PATIENT_BY_ID_PARAMETER + "}"
	DELETE_PATIENT_BY_ID_ENDPOINT     = "/patients/{" + DELETE_PATIENT_BY_ID_PARAMETER + "}"

	// Parameters
	FETCH_PATIENT_BY_ID_PARAMETER      = "patient_id"
	FETCH_PATIENT_BY_EMAIL_PARAMETER   = "patient_email"
	FETCH_PATIENT_BY_USER_ID_PARAMETER = "patient_id"
	UPDATE_PATIENT_BY_ID_PARAMETER     = "patient_id"
	DELETE_PATIENT_BY_ID_PARAMETER     = "patient_id"
)

const (
	DatabaseName = "pdp_db"
	TableName    = "patient"

	ColumnIDPacient    = "id_patient"
	ColumnIDUser       = "id_user"
	ColumnNume         = "nume"
	ColumnPrenume      = "prenume"
	ColumnEmail        = "email"
	ColumnTelefon      = "telefon"
	ColumnCNP          = "cnp"
	ColumnDataNasterii = "data_nasterii"
	ColumnIsActive     = "is_active"
)

const CNP_DATE_FORMAT = "060102"
