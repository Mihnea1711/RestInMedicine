package models

import (
	"time"
)

type Pacient struct {
	IDPacient    int       `db:"id_patient" json:"idPatient" sql:"type:int primary key"`
	IDUser       int       `db:"id_user" json:"id_user" sql:"type:int"`
	Nume         string    `db:"nume" json:"nume" sql:"type:varchar(50)"`
	Prenume      string    `db:"prenume" json:"prenume" sql:"type:varchar(50)"`
	Email        string    `db:"email" json:"email" sql:"type:varchar(70) unique"`
	Telefon      string    `db:"telefon" json:"telefon" sql:"type:char(10) check (telefon ~ '^[0-9]{10}$')"`
	CNP          string    `db:"cnp" json:"cnp" sql:"type:char(13) unique"`
	DataNasterii time.Time `db:"data_nasterii" json:"data_nasterii" sql:"type:date"`
	IsActive     bool      `db:"is_active" json:"is_active"`
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
