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
	DB_REQ_TIMEOUT_SEC_MULTIPLIER = 5
	CLEAR_DB_RESOURCES_TIMEOUT    = 10
)

const (
	RoleAdministrator UserRole = "administrator"
	RoleDoctor        UserRole = "doctor"
	RolePatient       UserRole = "pacient"
)

// UserTable and RoleTable are constants for table names
const (
	UserTable = "User"
	RoleTable = "Role"

	AliasRole = "r"
	AliasUser = "u"
)

// UserTableColumns and RoleTableColumns are constants for column names in their respective tables
const (
	ColumnIDUser       = "IDUser"
	ColumnUserName     = "Username"
	ColumnUserPassword = "Password"

	ColumnIDRole     = "IDRole"
	ColumnRoleIDUser = "IDUser"
	ColumnRole       = "Role"
)

const (
	MySQLDuplicateEntryErrorCode = 1062
)

// Constants for service and method names
const (
	IDMServiceName               = "/IDM/"
	RegisterMethodName           = "Register"
	LoginMethodName              = "Login"
	UpdateUserMethodName         = "UpdateUserByID"
	UpdateUserRoleMethodName     = "UpdateUserRole"
	UpdateUserPasswordMethodName = "UpdateUserPassword"
)
