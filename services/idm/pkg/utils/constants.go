package utils

type contextKey string
type UserRole string

const CONFIG_PATH = "configs/config.yaml"

const (
	IDM_TABLE = "idm"
)

const DECODED_IDM contextKey = "decodedIDM"

const (
	DEFAULT_PAGINATION_LIMIT = 20
	MAX_PAGINATION_LIMIT     = 50
	DEFAULT_PAGINATION_PAGE  = 1
)

const (
	RoleAdministrator UserRole = "administrator"
	RoleDoctor        UserRole = "doctor"
	RolePatient       UserRole = "pacient"
)
