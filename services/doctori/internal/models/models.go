package models

type Specializare string

// se poate adauga validare pe obiect aici dar ar fi de preferat sa fie facuta intr-unmiddleware pentru flexibiliate, adaptabilitate si cuplare scazuta
type Doctor struct {
	IDDoctor     int          `db:"id_doctor" json:"idDoctor" sql:"type:int primary key generated always as identity"`
	IDUser       int          `db:"id_user" json:"idUser" sql:"type:int"`
	Nume         string       `db:"nume" json:"nume" sql:"type:varchar(50)"`
	Prenume      string       `db:"prenume" json:"prenume" sql:"type:varchar(50)"`
	Email        string       `db:"email" json:"email" sql:"type:varchar(70) unique"`
	Telefon      string       `db:"telefon" json:"telefon" sql:"type:char(10) check (telefon ~ '^[0-9]{10}$')"`
	Specializare Specializare `db:"specializare" json:"specializare" sql:"type:enum"`
	IsActive     bool         `db:"is_active" json:"is_active"`
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
