package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// se poate adauga validare pe obiect aici dar ar fi de preferat sa fie facuta intr-unmiddleware pentru flexibiliate, adaptabilitate si cuplare scazuta
type Consultatie struct {
	IDConsultatie primitive.ObjectID `json:"id_consultatie" bson:"id_consultatie"`
	IDPacient     int                `json:"id_pacient" bson:"id_pacient"`
	IDDoctor      int                `json:"id_doctor" bson:"id_doctor"`
	Date          time.Time          `json:"date" bson:"date"`
	Diagnostic    string             `json:"diagnostic" bson:"diagnostic"`
	Investigatii  []Investigatie     `json:"investigatii" bson:"investigatii"`
}

type Investigatie struct {
	ID              primitive.ObjectID `json:"id_investigatie" bson:"id_investigatie"`
	Denumire        string             `json:"denumire" bson:"denumire"`
	DurataProcesare int                `json:"durata_procesare" bson:"durata_procesare"`
	Rezultat        string             `json:"rezultat" bson:"rezultat"`
}
