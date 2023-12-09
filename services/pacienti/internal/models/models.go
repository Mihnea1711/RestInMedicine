package models

import (
	"time"
)

type Patient struct {
	IDPatient   int       `db:"id_patient" json:"idPatient" sql:"type:int primary key"`
	IDUser      int       `db:"id_user" json:"idUser" sql:"type:int"`
	FirstName   string    `db:"first_name" json:"firstName" sql:"type:varchar(50)"`
	SecondName  string    `db:"second_name" json:"secondName" sql:"type:varchar(50)"`
	Email       string    `db:"email" json:"email" sql:"type:varchar(70) unique"`
	PhoneNumber string    `db:"phone_number" json:"phoneNumber" sql:"type:char(10) check (phone_number ~ '^[0-9]{10}$')"`
	CNP         string    `db:"cnp" json:"cnp" sql:"type:char(13) unique"`
	BirthDay    time.Time `db:"birth_day" json:"birthDay" sql:"type:date"`
	IsActive    bool      `db:"is_active" json:"isActive"`
}

type ResponseData struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Payload interface{} `json:"payload"`
}

type RowsAffected struct {
	RowsAffected int `json:"rowsAffected"`
}

type LastInsertedID struct {
	LastInsertedID int `json:"lastInsertedID"`
}

type ComplexResponse struct {
	RowsAffected int `json:"rowsAffected"`
	DeletedID    int `json:"deletedID"`
}

type ActivityData struct {
	IsActive bool `json:"isActive"`
	IDUser   int  `json:"idUser"`
}
