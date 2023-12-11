package utils

type contextKey string

const CONFIG_PATH = "configs/config.yaml"

const DECODED_CONSULTATION contextKey = "decodedConsultation"

const DATABASE_NAME = "consultations_db"
const CONSULTATIE_TABLE = "consultation"

const (
	COLUMN_ID_CONSULTATIE = "_id"
	COLUMN_ID_PATIENT     = "id_patient"
	COLUMN_ID_DOCTOR      = "id_doctor"
	COLUMN_DATE           = "date"
	COLUMN_DIAGNOSTIC     = "diagnostic"
	COLUMN_INVESTIGATII   = "investigations"
)
const (
	ID_INVESTIGATION = "id_investigation"
	NAME             = "name"
	PROCESSING_TIME  = "processing_time"
	RESULT           = "result"
)

const (
	REQUEST_TIMEOUT_DURATION           = 5
	CONNECTION_TIMEOUT_DB              = 10
	RESOURCES_CLOSE_TIMEOUT            = 10
	REQUEST_RATE                       = 10
	REQUEST_WINDOW_DURATION_MULTIPLIER = 1
)

const (
	DEFAULT_PAGINATION_LIMIT = 20
	MAX_PAGINATION_LIMIT     = 50
	DEFAULT_PAGINATION_PAGE  = 1
)

const (
	QUERY_PATIENT_ID = "patientID"
	QUERY_DOCTOR_ID  = "doctorID"
	QUERY_DATE       = "date"
	QUERY_PAGE       = "page"
	QUERY_LIMIT      = "limit"
)

const TIME_FORMAT = "2006-01-02"

const (
	INSERT_CONSULTATIE_ENDPOINT = "/consultations"

	FETCH_ALL_CONSULTATII_ENDPOINT = "/consultations"

	FILTER_CONSULTATII_ENDPOINT = "/consultations/filter"

	FETCH_CONSULTATIE_BY_DOCTOR_ID_ENDPOINT  = "/consultations/doctors/{" + FETCH_CONSULTATIE_BY_DOCTOR_ID_PARAMETER + "}"
	FETCH_CONSULTATIE_BY_DOCTOR_ID_PARAMETER = "id_doctor"

	FETCH_CONSULTATIE_BY_PATIENT_ID_ENDPOINT  = "/consultations/patients/{" + FETCH_CONSULTATIE_BY_PATIENT_ID_PARAMETER + "}"
	FETCH_CONSULTATIE_BY_PATIENT_ID_PARAMETER = "id_pacient"

	FETCH_CONSULTATIE_BY_DATE_ENDPOINT  = "/consultations/date/{" + FETCH_CONSULTATIE_BY_DATE_PARAMETER + "}"
	FETCH_CONSULTATIE_BY_DATE_PARAMETER = "id_date"

	FETCH_CONSULTATIE_BY_ID_ENDPOINT  = "/consultations/{" + FETCH_CONSULTATIE_BY_ID_PARAMETER + "}"
	FETCH_CONSULTATIE_BY_ID_PARAMETER = "id_consultation"

	UPDATE_CONSULTATIE_BY_ID_ENDPOINT  = "/consultations/{" + UPDATE_CONSULTATIE_BY_ID_PARAMETER + "}"
	UPDATE_CONSULTATIE_BY_ID_PARAMETER = "id_consultation"

	DELETE_CONSULTATIE_BY_ID_ENDPOINT  = "/consultations/{" + DELETE_CONSULTATIE_BY_ID_PARAMETER + "}"
	DELETE_CONSULTATIE_BY_ID_PARAMETER = "id_consultation"

	DELETE_CONSULTATIE_BY_PATIENT_DOCTOR_ID_ENDPOINT  = "/consultations/id/{" + DELETE_CONSULTATIE_BY_ID_PARAMETER + "}"
	DELETE_CONSULTATIE_BY_PATIENT_DOCTOR_ID_PARAMETER = "id_patient_doctor"

	HEALTH_CHECK_ENDPOINT = "/consultations/health-check"
)

const DUPLICATE_KEY_ERROR_CODE = 11000
