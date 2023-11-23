package models

type User struct {
	IDUser   int    `db:"id_user" json:"id_user"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"-"`
}

type Role struct {
	IDRole int    `db:"id_role" json:"id_role"`
	IDUser int    `db:"id_user" json:"id_user"`
	Role   string `db:"role" json:"role"`
}

type BlacklistToken struct {
	IDBToken int    `db:"id_btoken" json:"id_btoken"`
	IDUser   int    `db:"id_user" json:"id_user"`
	Token    string `db:"token" json:"token"`
}

type UserRegistration struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type CredentialsRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type ResponseData struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Payload interface{} `json:"payload"`
}
