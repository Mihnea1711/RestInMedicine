package models

import (
	"time"
)

type StatusProgramare string

type Appointment struct {
	IDProgramare int              `db:"id_programare" json:"idProgramare"`
	IDPatient    int              `db:"id_pacient" json:"idPatient"`
	IDDoctor     int              `db:"id_doctor" json:"idDoctor"`
	Date         time.Time        `db:"date" json:"date"`
	Status       StatusProgramare `db:"status" json:"status"`
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
