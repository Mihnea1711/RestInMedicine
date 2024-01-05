package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserData struct {
	IDUser   int    `json:"idUser"`
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
	IDUser int    `json:"idUser" validate:"required,min=1"`
	Token  string `json:"token" validate:"required,min=1"`
}

type UserRegistrationData struct {
	Username string `json:"username" validate:"required,min=5,max=255"`
	Password string `json:"password" validate:"required,min=8,max=255"`
	Role     string `json:"role" validate:"required,min=1,max=50,oneof=admin patient doctor"`
}

type UserLoginData struct {
	Username string `json:"username" validate:"required,min=5,max=255"`
	Password string `json:"password" validate:"required,min=5,max=255"`
}

type PatientData struct {
	IDPatient   int       `db:"id_patient" json:"idPatient" sql:"type:int primary key"`
	IDUser      int       `db:"id_user" json:"idUser" sql:"type:int" validate:"required"`
	FirstName   string    `db:"first_name" json:"firstName" sql:"type:varchar(50)" validate:"required,max=50"`
	SecondName  string    `db:"second_name" json:"secondName" sql:"type:varchar(50)" validate:"required,max=50"`
	Email       string    `db:"email" json:"email" sql:"type:varchar(70) unique" validate:"required,email"`
	PhoneNumber string    `db:"phone_number" json:"phoneNumber" sql:"type:char(10) check (phone_number ~ '^[0-9]{10}$')" validate:"required,len=10,numeric"`
	CNP         string    `db:"cnp" json:"cnp" sql:"type:char(13) unique" validate:"required,len=13,numeric"`
	BirthDay    time.Time `db:"birth_day" json:"birthDay" sql:"type:date" validate:"required"`
	IsActive    bool      `db:"is_active" json:"isActive" validate:"required"`
}

type Specialization string
type DoctorData struct {
	IDDoctor        int            `db:"id_doctor" json:"idDoctor" sql:"type:int primary key"`
	IDUser          int            `db:"idUser" json:"idUser" sql:"type:int" validate:"required"`
	FirstName       string         `db:"first_name" json:"firstName" sql:"type:varchar(50)" validate:"required,max=50"`
	SecondName      string         `db:"second_name" json:"secondName" sql:"type:varchar(50)" validate:"required,max=50"`
	Email           string         `db:"email" json:"email" sql:"type:varchar(70) unique" validate:"required,email"`
	PhoneNumber     string         `db:"phone_number" json:"phoneNumber" sql:"type:char(10) check (telefon ~ '^[0-9]{10}$')" validate:"required,len=10,numeric"`
	SSpecialization Specialization `db:"specialization" json:"specialization" sql:"type:enum" validate:"required"`
	IsActive        bool           `db:"is_active" json:"isActive" validate:"required"`
}

type StatusAppointment string
type AppointmentData struct {
	IDProgramare int               `db:"id_programare" json:"idProgramare" sql:"type:int primary key"`
	IDPatient    int               `db:"id_patient" json:"idPatient" validate:"required"`
	IDDoctor     int               `db:"id_doctor" json:"idDoctor" validate:"required"`
	Date         time.Time         `db:"date" json:"date" validate:"required"`
	Status       StatusAppointment `db:"status" json:"status" validate:"required"`
}

type ConsultationData struct {
	IDConsultation primitive.ObjectID `json:"idConsultation" bson:"_id"`
	IDPatient      int                `json:"idPatient" bson:"id_patient" validate:"required"`
	IDDoctor       int                `json:"idDoctor" bson:"id_doctor" validate:"required"`
	Date           time.Time          `json:"date" bson:"date" validate:"required"`
	Diagnostic     string             `json:"diagnostic" bson:"diagnostic" validate:"required"`
	Investigations []Investigation    `json:"investigations" bson:"investigations" validate:"required"`
}

type Investigation struct {
	ID             primitive.ObjectID `json:"idInvestigation" bson:"id_investigatie"`
	Name           string             `json:"name" bson:"name" validate:"required"`
	ProcessingTime int                `json:"processingTime" bson:"processing_time" validate:"required"`
	Result         string             `json:"result" bson:"result" validate:"required"`
}

type ActivityData struct {
	IsActive bool `json:"isActive"`
}
