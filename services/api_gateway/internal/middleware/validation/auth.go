package validation

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// validateUserRegistrationData validates the UserRegistrationData struct using the validator package
func validateUserRegistrationData(registrationData models.UserRegistrationData) error {
	validate := validator.New()
	return validate.Struct(registrationData)
}

// validateUserLoginData validates the UserLoginData struct using the validator package
func validateUserLoginData(loginData models.UserLoginData) error {
	validate := validator.New()
	return validate.Struct(loginData)
}

// ValidateRegistrationData is a middleware that validates UserRegistrationData
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

		// Validate UserRegistrationData
		if err := validateUserRegistrationData(registrationData); err != nil {
			utils.SendErrorResponse(w, http.StatusBadRequest, "Validation error", err.Error())
			return
		}

		// If validation passes, proceed to the next handler
		ctx := context.WithValue(r.Context(), utils.DECODED_USER_REGISTRATION_DATA, &registrationData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ValidateLoginData is a middleware that validates UserLoginData
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

		// Validate UserLoginData
		if err := validateUserLoginData(loginData); err != nil {
			utils.SendErrorResponse(w, http.StatusBadRequest, "Validation error", err.Error())
			return
		}

		// If validation passes, proceed to the next handler
		ctx := context.WithValue(r.Context(), utils.DECODED_USER_LOGIN_DATA, &loginData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
