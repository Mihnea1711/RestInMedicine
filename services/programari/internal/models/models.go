package models

import (
	"time"
)

type StatusProgramare string

type Programare struct {
	IDProgramare int              `db:"id_programare" json:"id"`
	IDPacient    int              `db:"id_pacient" json:"idPacient"`
	IDDoctor     int              `db:"id_doctor" json:"idDoctor"`
	Data         time.Time        `db:"data" json:"data"`
	Status       StatusProgramare `db:"status" json:"status"`
}
