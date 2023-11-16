package models

import (
	"time"
)

type StatusProgramare string

type Programare struct {
	IDProgramare int              `db:"id_programare" json:"id"`
	IDPacient    int              `db:"id_pacient" json:"idPacient"`
	IDDoctor     int              `db:"id_doctor" json:"idDoctor"`
	Date         time.Time        `db:"date" json:"date"`
	Status       StatusProgramare `db:"status" json:"status"`
}

type ResponseData struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Payload interface{} `json:"payload"`
}
