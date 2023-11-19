package utils

type contextKey string

const CONFIG_PATH = "configs/config.yaml"

const (
	COOKIE_MAX_AGE = 3600
	COOKIE_PATH    = "/api"
	COOKIE_NAME    = "jws"
)

const (
	IDM_HOST          = "localhost"
	IDM_PORT          = 8080
	PATIENT_PORT      = 8082
	DOCTOR_PORT       = 8083
	APPOINTMENT_PORT  = 8084
	CONSULTATION_PORT = 8085
)

const (
	POST   = "POST"
	GET    = "GET"
	PUT    = "PUT"
	DELETE = "DELETE"
)

const (
	ADMIN_ROLE   = "admin"
	PATIENT_ROLE = "patient"
	DOCTOR_ROLE  = "doctor"
)

const (
	DECODED_USER_REGISTRATION_DATA contextKey = "register_data"
	DECODED_USER_LOGIN_DATA        contextKey = "login_data"
)

const (
	// Endpoints
	REGISTER_USER_ENDPOINT     = "/api/users"
	LOGIN_USER_ENDPOINT        = "/api/login"
	GET_ALL_USERS_ENDPOINT     = "/api/users"
	GET_USER_BY_ID_ENDPOINT    = "/api/users/{" + GET_USER_ID_PARAMETER + "}"
	UPDATE_USER_BY_ID_ENDPOINT = "/api/users/{" + UPDATE_USER_ID_PARAMETER + "}"
	DELETE_USER_BY_ID_ENDPOINT = "/api/users/{" + DELETE_USER_ID_PARAMETER + "}"
	UPDATE_PASSWORD_ENDPOINT   = "/api/users/{" + UPDATE_USER_ROLE_ID_PARAMETER + "}/update-password"
	UPDATE_ROLE_ENDPOINT       = "/api/users/{" + UPDATE_USER_ROLE_ID_PARAMETER + "}/update-role"

	// Parameters
	GET_USER_ID_PARAMETER             = "userID"
	UPDATE_USER_ID_PARAMETER          = "userID"
	DELETE_USER_ID_PARAMETER          = "userID"
	UPDATE_USER_PASSWORD_ID_PARAMETER = "userID"
	UPDATE_USER_ROLE_ID_PARAMETER     = "userID"
)

const (
	ADD_TO_BLACKLIST_ENDPOINT      = "/api/blacklist"
	CHECK_BLACKLIST_ENDPOINT       = "/api/blacklist/{" + BLACKLIST_USER_ID_PARAMETER + "}"
	DELETE_FROM_BLACKLIST_ENDPOINT = "/api/blacklist/{" + BLACKLIST_USER_ID_PARAMETER + "}"

	BLACKLIST_USER_ID_PARAMETER = "userID"
)

const (
	// Endpoints
	CREATE_PATIENT_ENDPOINT         = "/api/patients"
	GET_ALL_PATIENTS_ENDPOINT       = "/api/patients"
	GET_PATIENT_BY_ID_ENDPOINT      = "/api/patients/{" + GET_PATIENT_ID_PARAMETER + "}"
	GET_PATIENT_BY_EMAIL_ENDPOINT   = "/api/patients/email/{" + GET_PATIENT_EMAIL_PARAMETER + "}"
	GET_PATIENT_BY_USER_ID_ENDPOINT = "/api/patients/users/{" + GET_PATIENT_USER_ID_PARAMETER + "}"
	UPDATE_PATIENT_BY_ID_ENDPOINT   = "/api/patients/{" + UPDATE_PATIENT_ID_PARAMETER + "}"
	DELETE_PATIENT_BY_ID_ENDPOINT   = "/api/patients/{" + DELETE_PATIENT_ID_PARAMETER + "}"

	// Parameters
	GET_PATIENT_ID_PARAMETER      = "patientID"
	GET_PATIENT_EMAIL_PARAMETER   = "patientEmail"
	GET_PATIENT_USER_ID_PARAMETER = "patientUserID"
	UPDATE_PATIENT_ID_PARAMETER   = "patientID"
	DELETE_PATIENT_ID_PARAMETER   = "patientID"

	// PATIENT_Endpoints
	PATIENT_CREATE_PATIENT_ENDPOINT           = "/patients"
	PATIENT_FETCH_ALL_PATIENTS_ENDPOINT       = "/patients"
	PATIENT_FETCH_PATIENT_BY_ID_ENDPOINT      = "/patients"
	PATIENT_FETCH_PATIENT_BY_EMAIL_ENDPOINT   = "/patients/email"
	PATIENT_FETCH_PATIENT_BY_USER_ID_ENDPOINT = "/patients/users"
	PATIENT_UPDATE_PATIENT_BY_ID_ENDPOINT     = "/patients"
	PATIENT_DELETE_PATIENT_BY_ID_ENDPOINT     = "/patients"
)

const (
	// Endpoints
	CREATE_DOCTOR_ENDPOINT         = "/api/doctors"
	GET_ALL_DOCTORS_ENDPOINT       = "/api/doctors"
	GET_DOCTOR_BY_ID_ENDPOINT      = "/api/doctors/{" + GET_DOCTOR_BY_ID_PARAMETER + "}"
	GET_DOCTOR_BY_EMAIL_ENDPOINT   = "/api/doctors/email/{" + GET_DOCTOR_BY_EMAIL_PARAMETER + "}"
	GET_DOCTOR_BY_USER_ID_ENDPOINT = "/api/doctors/users/{" + GET_DOCTOR_BY_USER_ID_PARAMETER + "}"
	UPDATE_DOCTOR_BY_ID_ENDPOINT   = "/api/doctors/{" + UPDATE_DOCTOR_BY_ID_PARAMETER + "}"
	DELETE_DOCTOR_BY_ID_ENDPOINT   = "/api/doctors/{" + DELETE_DOCTOR_BY_ID_PARAMETER + "}"

	// Parameters
	GET_DOCTOR_BY_ID_PARAMETER      = "doctorID"
	GET_DOCTOR_BY_EMAIL_PARAMETER   = "doctorEmail"
	GET_DOCTOR_BY_USER_ID_PARAMETER = "doctorID"
	UPDATE_DOCTOR_BY_ID_PARAMETER   = "doctorID"
	DELETE_DOCTOR_BY_ID_PARAMETER   = "doctorID"

	// DOCTOR_Endpoints
	DOCTOR_CREATE_DOCTOR_ENDPOINT           = "/doctors"
	DOCTOR_FETCH_ALL_DOCTORS_ENDPOINT       = "/doctors"
	DOCTOR_FETCH_DOCTOR_BY_ID_ENDPOINT      = "/doctors"
	DOCTOR_FETCH_DOCTOR_BY_EMAIL_ENDPOINT   = "/doctors/email"
	DOCTOR_FETCH_DOCTOR_BY_USER_ID_ENDPOINT = "/doctors/users"
	DOCTOR_UPDATE_DOCTOR_BY_ID_ENDPOINT     = "/doctors"
	DOCTOR_DELETE_DOCTOR_BY_ID_ENDPOINT     = "/doctors"
)

const (
	// Endpoints
	CREATE_APPOINTMENT_ENDPOINT             = "/api/appointments"
	GET_ALL_APPOINTMENTS_ENDPOINT           = "/api/appointments"
	GET_APPOINTMENT_BY_ID_ENDPOINT          = "/api/appointments/{" + GET_APPOINTMENT_ID_PARAMETER + "}"
	GET_APPOINTMENTS_BY_DOCTOR_ID_ENDPOINT  = "/api/appointments/doctors/{" + DOCTOR_ID_PARAMETER + "}"
	GET_APPOINTMENTS_BY_PACIENT_ID_ENDPOINT = "/api/appointments/pacients/{" + PACIENT_ID_PARAMETER + "}"
	GET_APPOINTMENTS_BY_DATE_ENDPOINT       = "/api/appointments/date/{" + APPOINTMENT_DATE_PARAMETER + "}"
	GET_APPOINTMENTS_BY_STATUS_ENDPOINT     = "/api/appointments/status/{" + APPOINTMENT_STATUS_PARAMETER + "}"
	UPDATE_APPOINTMENT_BY_ID_ENDPOINT       = "/api/appointments/{" + UPDATE_APPOINTMENT_ID_PARAMETER + "}"
	DELETE_APPOINTMENT_BY_ID_ENDPOINT       = "/api/appointments/{" + DELETE_APPOINTMENT_ID_PARAMETER + "}"

	// Parameters
	GET_APPOINTMENT_ID_PARAMETER    = "appointmentID"
	UPDATE_APPOINTMENT_ID_PARAMETER = "appointmentID"
	DELETE_APPOINTMENT_ID_PARAMETER = "appointmentID"
	DOCTOR_ID_PARAMETER             = "doctorID"
	PACIENT_ID_PARAMETER            = "pacientID"
	APPOINTMENT_DATE_PARAMETER      = "apointmentDate"
	APPOINTMENT_STATUS_PARAMETER    = "apointmentStatus"
)

const (
	// Endpoints
	CREATE_CONSULTATION_ENDPOINT            = "/api/consultations"
	GET_ALL_CONSULTATIONS_ENDPOINT          = "/api/consultations"
	GET_CONSULTATION_BY_DOCTOR_ID_ENDPOINT  = "/api/consultations/doctors/{" + GET_CONSULTATION_BY_DOCTOR_ID_PARAMETER + "}"
	GET_CONSULTATION_BY_PACIENT_ID_ENDPOINT = "/api/consultations/patients/{" + GET_CONSULTATION_BY_PACIENT_ID_PARAMETER + "}"
	GET_CONSULTATION_BY_DATE_ENDPOINT       = "/api/consultations/date/{" + GET_CONSULTATION_BY_DATE_PARAMETER + "}"
	GET_CONSULTATION_BY_ID_ENDPOINT         = "/api/consultations/{" + GET_CONSULTATION_BY_ID_PARAMETER + "}"
	UPDATE_CONSULTATION_BY_ID_ENDPOINT      = "/api/consultations/{" + UPDATE_CONSULTATION_BY_ID_PARAMETER + "}"
	DELETE_CONSULTATION_BY_ID_ENDPOINT      = "/api/consultations/{" + DELETE_CONSULTATION_BY_ID_PARAMETER + "}"

	// Parameters
	GET_CONSULTATION_BY_DOCTOR_ID_PARAMETER  = "doctorID"
	GET_CONSULTATION_BY_PACIENT_ID_PARAMETER = "pacientID"
	GET_CONSULTATION_BY_DATE_PARAMETER       = "consultationDate"
	GET_CONSULTATION_BY_ID_PARAMETER         = "consultationID"
	UPDATE_CONSULTATION_BY_ID_PARAMETER      = "consultationID"
	DELETE_CONSULTATION_BY_ID_PARAMETER      = "consultationID"
)

const (
	// Multipliers
	DB_CLEAR_RESOURCES_MULTIPLIER = 10

	// Timeouts
	REQUEST_CONTEXT_TIMEOUT = 10
)
