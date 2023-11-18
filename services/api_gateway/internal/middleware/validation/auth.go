package validation

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

func ValidateRegistrationData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var registrationData models.UserRegistrationData

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&registrationData)
		if checkErrorOnDecode(err, w) {
			logAndRespondWithError(w, http.StatusBadRequest, "Error decoding registration request body", err)
			return
		}

		if err := validatePassword(registrationData.Password); err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Error validating password", err)
			return
		}

		if err := validateUsername(registrationData.Username); err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Error validating username", err)
			return
		}

		if err := validateRole(registrationData.Role); err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Error validating role", err)
			return
		}

		// If all validations pass, proceed to the actual controller
		ctx := context.WithValue(r.Context(), utils.DECODED_USER_REGISTRATION_DATA, &registrationData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ValidateLoginData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var loginData models.UserLoginData

		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()

		err := dec.Decode(&loginData)
		if checkErrorOnDecode(err, w) {
			logAndRespondWithError(w, http.StatusBadRequest, "Error decoding registration request body", err)
			return
		}

		if err := validatePassword(loginData.Password); err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Error validating password", err)
			return
		}

		if err := validateUsername(loginData.Username); err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Error validating username", err)
			return
		}

		// If all validations pass, proceed to the actual controller
		ctx := context.WithValue(r.Context(), utils.DECODED_USER_LOGIN_DATA, &loginData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func validatePassword(password string) error {
	if len(password) < 5 {
		return errors.New("password must be at least 5 characters long")
	}
	return nil
}

func validateUsername(username string) error {
	if len(username) < 5 {
		return errors.New("username must be at least 5 characters long")
	}
	return nil
}

func validateRole(role string) error {
	switch role {
	case utils.ADMIN_ROLE, utils.PATIENT_ROLE, utils.DOCTOR_ROLE:
		return nil
	default:
		return errors.New("invalid role")
	}
}
