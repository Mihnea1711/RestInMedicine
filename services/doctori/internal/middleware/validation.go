package middleware

import (
	"encoding/json"
	"net/http"
	"regexp"

	"github.com/mihnea1711/POS_Project/services/doctori/internal/models"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func ValidateDoctorCreation(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var doctor models.Doctor

		// Decode the request body into the Doctor struct
		if err := json.NewDecoder(r.Body).Decode(&doctor); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		// Basic validation for each field
		if doctor.Nume == "" || len(doctor.Nume) > 50 {
			http.Error(w, "Invalid or missing Nume", http.StatusBadRequest)
			return
		}

		if doctor.Prenume == "" || len(doctor.Prenume) > 50 {
			http.Error(w, "Invalid or missing Prenume", http.StatusBadRequest)
			return
		}

		// Validate email format using regex
		emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		if !regexp.MustCompile(emailRegex).MatchString(doctor.Email) || len(doctor.Email) > 70 {
			http.Error(w, "Invalid or missing Email", http.StatusBadRequest)
			return
		}

		// Validate Romanian phone number format
		phoneRegex := `^(07[0-9]{8}|\+407[0-9]{8})$`
		if !regexp.MustCompile(phoneRegex).MatchString(doctor.Telefon) {
			http.Error(w, "Invalid or missing Telefon", http.StatusBadRequest)
			return
		}

		if !isValidSpecializare(models.Specializare(doctor.Specializare)) {
			http.Error(w, "Invalid Specializare", http.StatusBadRequest)
			return
		}

		// If all validations pass, proceed to the actual controller
		next.ServeHTTP(w, r)
	})
}

func isValidSpecializare(specializare models.Specializare) bool {
	for _, validSpec := range utils.ValidSpecializari {
		if specializare == validSpec {
			return true
		}
	}
	return false
}

// add similar validation middlewares for Update, Delete, etc.
