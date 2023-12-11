package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// se poate adauga validare pe obiect aici dar ar fi de preferat sa fie facuta intr-unmiddleware pentru flexibiliate, adaptabilitate si cuplare scazuta
type Consultation struct {
	IDConsultation primitive.ObjectID `json:"idConsultation" bson:"_id"`
	IDPatient      int                `json:"idPatient" bson:"id_patient"`
	IDDoctor       int                `json:"idDoctor" bson:"id_doctor"`
	Date           time.Time          `json:"date" bson:"date"`
	Diagnostic     string             `json:"diagnostic" bson:"diagnostic"`
	Investigations []Investigation    `json:"investigations" bson:"investigations"`
}

type Investigation struct {
	IDInvestigation primitive.ObjectID `json:"idInvestigation" bson:"id_investigation"`
	Name            string             `json:"name" bson:"name"`
	ProcessingTime  int                `json:"processingTime" bson:"processing_time"`
	Result          string             `json:"result" bson:"result"`
}

type ResponseData struct {
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Payload interface{} `json:"payload,omitempty"`
}

type RowsAffected struct {
	RowsAffected int `json:"rowsAffected"`
}

type LastInsertedID struct {
	LastInsertedID string `json:"lastInsertedID"`
}
