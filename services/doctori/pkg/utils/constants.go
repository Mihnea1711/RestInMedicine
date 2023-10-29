package utils

import (
	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
)

type contextKey string

const (
	Cardiologie  models.Specializare = "Cardiologie"
	Neurologie   models.Specializare = "Neurologie"
	Ortopedie    models.Specializare = "Ortopedie"
	Pediatrie    models.Specializare = "Pediatrie"
	Dermatologie models.Specializare = "Dermatologie"
	Radiologie   models.Specializare = "Radiologie"
	Chirurgie    models.Specializare = "Chirurgie"
)

var ValidSpecializari = [...]models.Specializare{Cardiologie, Neurologie, Ortopedie, Pediatrie, Dermatologie, Radiologie, Chirurgie}

const CONFIG_PATH = "configs/config.yaml"

const (
	DOCTOR_TABLE = "doctor"
)

const DECODED_DOCTOR contextKey = "decodedDoctor"
