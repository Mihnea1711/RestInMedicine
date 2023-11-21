package models

import "time"

type UserRegistrationData struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type UserLoginData struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type PacientData struct {
	IDUser       int       `db:"id_user" json:"id_user" sql:"type:int references User"`
	Nume         string    `db:"nume" json:"nume" sql:"type:varchar(50)"`
	Prenume      string    `db:"prenume" json:"prenume" sql:"type:varchar(50)"`
	Email        string    `db:"email" json:"email" sql:"type:varchar(70) unique"`
	Telefon      string    `db:"telefon" json:"telefon" sql:"type:char(10) check (telefon ~ '^[0-9]{10}$')"`
	CNP          string    `db:"cnp" json:"cnp" sql:"type:char(13) unique"`
	DataNasterii time.Time `db:"data_nasterii" json:"data_nasterii" sql:"type:date"`
	IsActive     bool      `db:"is_active" json:"is_active"`
}

type Specializare string
type DoctorData struct {
	IDDoctor     int          `db:"id_doctor" json:"idDoctor" sql:"type:int primary key generated always as identity"`
	IDUser       int          `db:"id_user" json:"idUser" sql:"type:int"`
	Nume         string       `db:"nume" json:"nume" sql:"type:varchar(50)"`
	Prenume      string       `db:"prenume" json:"prenume" sql:"type:varchar(50)"`
	Email        string       `db:"email" json:"email" sql:"type:varchar(70) unique"`
	Telefon      string       `db:"telefon" json:"telefon" sql:"type:char(10) check (telefon ~ '^[0-9]{10}$')"`
	Specializare Specializare `db:"specializare" json:"specializare" sql:"type:enum"`
}

type StatusProgramare string
type AppointmentData struct {
	IDProgramare int              `db:"id_programare" json:"idProgramare"`
	IDPacient    int              `db:"id_pacient" json:"idPacient"`
	IDDoctor     int              `db:"id_doctor" json:"idDoctor"`
	Date         time.Time        `db:"date" json:"date"`
	Status       StatusProgramare `db:"status" json:"status"`
}

type UserData struct {
	IDUser   int    `json:"id_user"`
	Username string `json:"username"`
	// other if needed
}

type PasswordData struct {
	Password string `json:"password"`
}

type RoleData struct {
	Role string `json:"role"`
}

// MakeAppointmentRequest represents the request payload for making an appointment.
type MakeAppointmentRequest struct {
	IDPatient int    `json:"id_patient"`
	IDDoctor  int    `json:"id_doctor"`
	Date      string `json:"date"`
}

// ProgramConsultationRequest represents the request payload for programming a consultation.
type ProgramConsultationRequest struct {
	IDPatient int    `json:"id_patient"`
	IDDoctor  int    `json:"id_doctor"`
	Date      string `json:"date"`
}

type ResponseData struct {
	Message string      `json:"message"`
	Error   string      `json:"error"`
	Payload interface{} `json:"payload"`
}

type RowsAffected struct {
	RowsAffected int `json:"rows_affected"`
}

type BlacklistData struct {
	IDUser int    `json:"id_user"`
	Token  string `json:"token"`
}
