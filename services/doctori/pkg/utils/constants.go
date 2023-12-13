package utils

import (
	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
)

type contextKey string

const (
	Cardiology  models.Specialization = "Cardiology"
	Neurology   models.Specialization = "Neurology"
	Orthopedics models.Specialization = "Orthopedics"
	Pediatrics  models.Specialization = "Pediatrics"
	Dermatology models.Specialization = "Dermatology"
	Radiology   models.Specialization = "Radiology"
	Surgery     models.Specialization = "Surgery"
)

var ValidSpecializations = [...]models.Specialization{Cardiology, Neurology, Orthopedics, Pediatrics, Dermatology, Radiology, Surgery}

const CONFIG_PATH = "configs/config.yaml"

const DECODED_DOCTOR contextKey = "decodedDoctor"
const DECODED_DOCTOR_ACTIVITY contextKey = "decodedDoctorActivity"

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

// Doctor module endpoints with parameters
const (
	HEALTH_CHECK_ENDPOINT             = "/doctors/health-check"
	CREATE_DOCTOR_ENDPOINT            = "/doctors"
	FETCH_ALL_DOCTORS_ENDPOINT        = "/doctors"
	FETCH_DOCTOR_BY_ID_ENDPOINT       = "/doctors/{" + FETCH_DOCTOR_BY_ID_PARAMETER + "}"
	FETCH_DOCTOR_BY_EMAIL_ENDPOINT    = "/doctors/email/{" + FETCH_DOCTOR_BY_EMAIL_PARAMETER + "}"
	FETCH_DOCTOR_BY_USER_ID_ENDPOINT  = "/doctors/users/{" + FETCH_DOCTOR_BY_USER_ID_PARAMETER + "}"
	UPDATE_DOCTOR_BY_ID_ENDPOINT      = "/doctors/{" + UPDATE_DOCTOR_BY_ID_PARAMETER + "}"
	DELETE_DOCTOR_BY_ID_ENDPOINT      = "/doctors/{" + DELETE_DOCTOR_BY_ID_PARAMETER + "}"
	DELETE_DOCTOR_BY_USER_ID_ENDPOINT = "/doctors/users/{" + DELETE_DOCTOR_BY_USER_ID_PARAMETER + "}"
	TOGGLE_DOCTOR_ACTIVITY_ENDPOINT   = "/doctors/{" + PATCH_DOCTOR_BY_ID_PARAMETER + "}"
)

const (
	FETCH_DOCTOR_BY_ID_PARAMETER       = "doctorID"
	FETCH_DOCTOR_BY_EMAIL_PARAMETER    = "doctorEmail"
	FETCH_DOCTOR_BY_USER_ID_PARAMETER  = "doctorID"
	UPDATE_DOCTOR_BY_ID_PARAMETER      = "doctorID"
	PATCH_DOCTOR_BY_ID_PARAMETER       = "doctorUserID"
	DELETE_DOCTOR_BY_ID_PARAMETER      = "doctorID"
	DELETE_DOCTOR_BY_USER_ID_PARAMETER = "doctorUserID"

	QUERY_IS_ACIVE       = "isActive"
	QUERY_FIRST_NAME     = "firstName"
	QUERY_SPECIALIZATION = "specialization"
	QUERY_PAGE           = "page"
	QUERY_LIMIT          = "limit"
)

const (
	DatabaseName         = "pdp_db"
	DoctorTableName      = "doctor"
	ColumnIDDoctor       = "id_doctor"
	ColumnIDUser         = "id_user"
	ColumnFirstName      = "first_name"
	ColumnSecondName     = "second_name"
	ColumnEmail          = "email"
	ColumnPhoneNumber    = "phone_number"
	ColumnSpecialization = "specialization"
	ColumnIsActive       = "is_active"
)

const MySQLDuplicateEntryErrorCode = 1062

const (
	MaxNameLength       = 255
	MaxEmailLength      = 255
	PhoneNumberLength   = 10
	MaxRequestSizeBytes = 1 << 20 // 1MB
)
