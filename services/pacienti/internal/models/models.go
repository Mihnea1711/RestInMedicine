package models

import (
	"time"
)

type Pacient struct {
	CNP          string    `db:"cnp" json:"cnp" sql:"type:char(13) primary key"`
	IDUser       int       `db:"id_user" json:"id_user" sql:"type:int references Utilizatori"`
	Nume         string    `db:"nume" json:"nume" sql:"type:varchar(50)"`
	Prenume      string    `db:"prenume" json:"prenume" sql:"type:varchar(50)"`
	Email        string    `db:"email" json:"email" sql:"type:varchar(70) unique"`
	Telefon      string    `db:"telefon" json:"telefon" sql:"type:char(10) check (telefon ~ '^[0-9]{10}$')"`
	DataNasterii time.Time `db:"data_nasterii" json:"data_nasterii" sql:"type:date check (data_nasterii <= current_date - interval '18 years')"`
	IsActive     bool      `db:"is_active" json:"is_active"`
}
