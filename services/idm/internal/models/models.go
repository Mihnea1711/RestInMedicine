package models

type User struct {
	IDUser   int    `db:"id_user" json:"id_user"`
	Username string `db:"username" json:"username"`
	Password string `db:"password" json:"-"`
	Token    string `db:"token" json:"token"`
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
