package utils

type contextKey string

const CONFIG_PATH = "configs/config.yaml"

const DECODED_CONSULTATIE contextKey = "decodedConsultatie"

const CONSULTATIE_TABLE = "consultatie"

const (
	ID_CONSULTATIE = "id_consultatie"
	ID_PACIENT     = "id_pacient"
	ID_DOCTOR      = "id_doctor"
	DATE           = "date"
	DIAGNOSTIC     = "diagnostic"
	INVESTIGATII   = "investigatii"
)
const (
	ID_INVESTIGATIE  = "id_investigatie"
	DENUMIRE         = "denumire"
	DURATA_PROCESARE = "durata_procesare"
	REZULTAT         = "rezultat"
)

const (
	REQUEST_TIMEOUT_DURATION           = 5
	REQUEST_RATE                       = 10
	REQUEST_WINDOW_DURATION_MULTIPLIER = 1
)

const (
	DEFAULT_PAGINATION_LIMIT = 20
	MAX_PAGINATION_LIMIT     = 50
	DEFAULT_PAGINATION_PAGE  = 1
)

const TIME_FORMAT = "2006-01-02"

const (
	INSERT_CONSULTATIE_ENDPOINT = "/consultatii"

	FETCH_ALL_CONSULTATII_ENDPOINT = "/consultatii"

	FETCH_CONSULTATIE_BY_DOCTOR_ID_ENDPOINT  = "/consultatii/doctor/{" + FETCH_CONSULTATIE_BY_DOCTOR_ID_PARAMETER + "}"
	FETCH_CONSULTATIE_BY_DOCTOR_ID_PARAMETER = "id_doctor"

	FETCH_CONSULTATIE_BY_PACIENT_ID_ENDPOINT  = "/consultatii/pacient/{" + FETCH_CONSULTATIE_BY_PACIENT_ID_PARAMETER + "}"
	FETCH_CONSULTATIE_BY_PACIENT_ID_PARAMETER = "id_pacient"

	FETCH_CONSULTATIE_BY_DATE_ENDPOINT  = "/consultatii/date/{" + FETCH_CONSULTATIE_BY_DATE_PARAMETER + "}"
	FETCH_CONSULTATIE_BY_DATE_PARAMETER = "id_date"

	FETCH_CONSULTATIE_BY_ID_ENDPOINT  = "/consultatii/{" + FETCH_CONSULTATIE_BY_ID_PARAMETER + "}"
	FETCH_CONSULTATIE_BY_ID_PARAMETER = "id_consultatie"

	UPDATE_CONSULTATIE_BY_ID_ENDPOINT  = "/consultatii/{" + UPDATE_CONSULTATIE_BY_ID_PARAMETER + "}"
	UPDATE_CONSULTATIE_BY_ID_PARAMETER = "id_consultatie"

	DELETE_CONSULTATIE_BY_ID_ENDPOINT  = "/consultatii/{" + DELETE_CONSULTATIE_BY_ID_PARAMETER + "}"
	DELETE_CONSULTATIE_BY_ID_PARAMETER = "id_consultatie"
)
