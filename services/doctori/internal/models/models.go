package models

// se poate adauga validare pe obiect aici dar ar fi de preferat sa fie facuta intr-unmiddleware pentru flexibiliate, adaptabilitate si cuplare scazuta
type Doctor struct {
	IDDoctor     int    `db:"id_doctor" json:"idDoctor"`
	IDUser       int    `db:"id_user" json:"idUser"`
	Nume         string `db:"nume" json:"nume"`
	Prenume      string `db:"prenume" json:"prenume"`
	Email        string `db:"email" json:"email"`
	Telefon      string `db:"telefon" json:"telefon"`
	Specializare string `db:"specializare" json:"specializare"`
}
