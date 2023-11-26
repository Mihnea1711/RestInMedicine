package utils

import (
	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
)

type contextKey string

const (
	Cardiologie  models.Specializare = "Cardiologie"
	Neurologie   models.Specializare = "Neurologie"
	Ortopedie    models.Specializare = "Ortopedie"
	Pediatrie    models.Specializare = "Pediatrie"
	Dermatologie models.Specializare = "Dermatologie"
	Radiologie   models.Specializare = "Radiologie"
	Chirurgie    models.Specializare = "Chirurgie"
)

var ValidSpecializari = [...]models.Specializare{Cardiologie, Neurologie, Ortopedie, Pediatrie, Dermatologie, Radiologie, Chirurgie}

const CONFIG_PATH = "configs/config.yaml"

const DECODED_DOCTOR contextKey = "decodedDoctor"

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
	CREATE_DOCTOR_ENDPOINT            = "/doctors"                                                   // POST
	FETCH_ALL_DOCTORS_ENDPOINT        = "/doctors"                                                   // GET
	FETCH_DOCTOR_BY_ID_ENDPOINT       = "/doctors/{" + FETCH_DOCTOR_BY_ID_PARAMETER + "}"            // GET
	FETCH_DOCTOR_BY_EMAIL_ENDPOINT    = "/doctors/email/{" + FETCH_DOCTOR_BY_EMAIL_PARAMETER + "}"   // GET
	FETCH_DOCTOR_BY_USER_ID_ENDPOINT  = "/doctors/users/{" + FETCH_DOCTOR_BY_USER_ID_PARAMETER + "}" // GET
	UPDATE_DOCTOR_BY_ID_ENDPOINT      = "/doctors/{" + UPDATE_DOCTOR_BY_ID_PARAMETER + "}"           // PUT
	DELETE_DOCTOR_BY_ID_ENDPOINT      = "/doctors/{" + DELETE_DOCTOR_BY_ID_PARAMETER + "}"           // DELETE
	DELETE_DOCTOR_BY_USER_ID_ENDPOINT = "/doctors/users{" + DELETE_DOCTOR_BY_USER_ID_PARAMETER + "}" // DELETE
)

// Doctor module parameters
const (
	FETCH_DOCTOR_BY_ID_PARAMETER       = "doctor_id"
	FETCH_DOCTOR_BY_EMAIL_PARAMETER    = "doctor_email"
	FETCH_DOCTOR_BY_USER_ID_PARAMETER  = "doctor_id"
	UPDATE_DOCTOR_BY_ID_PARAMETER      = "doctor_id"
	DELETE_DOCTOR_BY_ID_PARAMETER      = "doctor_id"
	DELETE_DOCTOR_BY_USER_ID_PARAMETER = "doctor_user_id"
)

const (
	DatabaseName       = "pdp_db"
	DoctorTableName    = "doctor"
	ColumnIDDoctor     = "id_doctor"
	ColumnIDUser       = "id_user"
	ColumnNume         = "nume"
	ColumnPrenume      = "prenume"
	ColumnEmail        = "email"
	ColumnTelefon      = "telefon"
	ColumnSpecializare = "specializare"
)

const MySQLDuplicateEntryErrorCode = 1062
