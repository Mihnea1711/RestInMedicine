package models

type Specialization string

// se poate adauga validare pe obiect aici dar ar fi de preferat sa fie facuta intr-unmiddleware pentru flexibiliate, adaptabilitate si cuplare scazuta
type Doctor struct {
	IDDoctor       int            `db:"id_doctor" json:"idDoctor" sql:"type:int primary key generated always as identity"`
	IDUser         int            `db:"id_user" json:"idUser" sql:"type:int"`
	FirstName      string         `db:"frist_name" json:"frist_name" sql:"type:varchar(50)"`
	SecondName     string         `db:"second_name" json:"second_name" sql:"type:varchar(50)"`
	Email          string         `db:"email" json:"email" sql:"type:varchar(70) unique"`
	PhoneNumber    string         `db:"phone_number" json:"phone_number" sql:"type:char(10) check (telefon ~ '^[0-9]{10}$')"`
	Specialization Specialization `db:"specialization" json:"specialization" sql:"type:enum"`
	IsActive       bool           `db:"is_active" json:"is_active"`
}

type ResponseData struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Payload interface{} `json:"payload"`
}

type RowsAffected struct {
	RowsAffected int `json:"rows_affected"`
}

type LastInsertedID struct {
	LastInsertedID int `json:"last_inserted_id"`
}

type ComplexResponse struct {
	RowsAffected int `json:"rows_affected"`
	DeletedID    int `json:"deleted_id"`
}

type ActivityData struct {
	IsActive bool `json:"is_active"`
	IDUser   int  `json:"id_user"`
}
