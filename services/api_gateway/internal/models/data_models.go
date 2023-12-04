package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserData struct {
	IDUser   int    `json:"id_user"`
	Username string `json:"username" validate:"required,min=5,max=255"`
	// other if needed
}

type PasswordData struct {
	Password string `json:"password" validate:"required,min=8,max=255"`
}

type RoleData struct {
	Role string `json:"role" validate:"required,min=1,max=50,oneof=admin patient doctor"`
}

type BlacklistData struct {
	IDUser int    `json:"id_user" validate:"required,min=1"`
	Token  string `json:"token" validate:"required,min=1"`
}

type UserRegistrationData struct {
	Username string `json:"username" validate:"required,min=5,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Role     string `json:"role" validate:"required,min=1,max=50,oneof=admin patient doctor"`
}

type UserLoginData struct {
	Username string `json:"username" validate:"required,min=5,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
}

type PatientData struct {
	IDPatient    int       `db:"id_patient" json:"idPatient" sql:"type:int primary key"`
	IDUser       int       `db:"id_user" json:"id_user" sql:"type:int" validate:"required"`
	Nume         string    `db:"nume" json:"nume" sql:"type:varchar(50)" validate:"required,max=50"`
	Prenume      string    `db:"prenume" json:"prenume" sql:"type:varchar(50)" validate:"required,max=50"`
	Email        string    `db:"email" json:"email" sql:"type:varchar(70) unique" validate:"required,email"`
	Telefon      string    `db:"telefon" json:"telefon" sql:"type:char(10) check (telefon ~ '^[0-9]{10}$')" validate:"required,len=10,numeric"`
	CNP          string    `db:"cnp" json:"cnp" sql:"type:char(13) unique" validate:"required,len=13,numeric"`
	DataNasterii time.Time `db:"data_nasterii" json:"data_nasterii" sql:"type:date" validate:"required"`
	IsActive     bool      `db:"is_active" json:"is_active" validate:"required"`
}

type Specializare string
type DoctorData struct {
	IDDoctor     int          `db:"id_doctor" json:"idDoctor" sql:"type:int primary key"`
	IDUser       int          `db:"id_user" json:"idUser" sql:"type:int" validate:"required"`
	Nume         string       `db:"nume" json:"nume" sql:"type:varchar(50)" validate:"required,max=50"`
	Prenume      string       `db:"prenume" json:"prenume" sql:"type:varchar(50)" validate:"required,max=50"`
	Email        string       `db:"email" json:"email" sql:"type:varchar(70) unique" validate:"required,email"`
	Telefon      string       `db:"telefon" json:"telefon" sql:"type:char(10) check (telefon ~ '^[0-9]{10}$')" validate:"required,len=10,numeric"`
	Specializare Specializare `db:"specializare" json:"specializare" sql:"type:enum"`
	IsActive     bool         `db:"is_active" json:"is_active" validate:"required"`
}

type StatusProgramare string
type AppointmentData struct {
	IDProgramare int              `db:"id_programare" json:"idProgramare" sql:"type:int primary key"`
	IDPatient    int              `db:"id_patient" json:"idPatient" validate:"required"`
	IDDoctor     int              `db:"id_doctor" json:"idDoctor" validate:"required"`
	Date         time.Time        `db:"date" json:"date" validate:"required"`
	Status       StatusProgramare `db:"status" json:"status" validate:"required"`
}

type ConsultationData struct {
	IDConsultatie primitive.ObjectID `json:"id_consultation" bson:"id_consultation"`
	IDPatient     int                `json:"id_patient" bson:"id_patient" validate:"required"`
	IDDoctor      int                `json:"id_doctor" bson:"id_doctor" validate:"required"`
	Date          time.Time          `json:"date" bson:"date" validate:"required"`
	Diagnostic    string             `json:"diagnostic" bson:"diagnostic" validate:"required"`
	Investigatii  []Investigatie     `json:"investigatii" bson:"investigatii" validate:"required"`
}

type Investigatie struct {
	ID              primitive.ObjectID `json:"id_investigatie" bson:"id_investigatie" validate:"required"`
	Denumire        string             `json:"denumire" bson:"denumire" validate:"required"`
	DurataProcesare int                `json:"durata_procesare" bson:"durata_procesare" validate:"required"`
	Rezultat        string             `json:"rezultat" bson:"rezultat" validate:"required"`
}

type ActivityData struct {
	IsActive bool `json:"is_active"`
	IDUser   int  `json:"id_user" validate:"required"`
}
