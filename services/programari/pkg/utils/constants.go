package utils

import "github.com/mihnea1711/POS_Project/services/programari/internal/models"

type contextKey string

const (
	StatusOnorata    models.StatusProgramare = "onorata"
	StatusNeprezenta models.StatusProgramare = "neprezenta"
	StatusAnulata    models.StatusProgramare = "anulata"
)

var ValidStatus = [...]models.StatusProgramare{StatusOnorata, StatusNeprezenta, StatusAnulata}

const CONFIG_PATH = "configs/config.yaml"

const (
	PROGRAMARE_TABLE = "programare"
)

const DECODED_PROGRAMARE contextKey = "decodedProgramare"

const (
	DEFAULT_PAGINATION_LIMIT = 20
	MAX_PAGINATION_LIMIT     = 50
	DEFAULT_PAGINATION_PAGE  = 1
)
