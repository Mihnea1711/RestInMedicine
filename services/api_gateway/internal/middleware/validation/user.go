package validation

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// validateUserData validates the UserData struct using the validator package
func validateUserData(userData models.UserData) error {
	validate := validator.New()
	return validate.Struct(userData)
}

// validatePasswordData validates the PasswordData struct using the validator package
func validatePasswordData(passwordData models.PasswordData) error {
	validate := validator.New()
	return validate.Struct(passwordData)
}

// validateRoleData validates the RoleData struct using the validator package
func validateRoleData(roleData models.RoleData) error {
	validate := validator.New()
	return validate.Struct(roleData)
}

// validateBlacklistData validates the BlacklistData struct using the validator package
func validateBlacklistData(blacklistData models.BlacklistData) error {
	validate := validator.New()
	return validate.Struct(blacklistData)
}

// ValidateUserData is a middleware that validates UserData
func ValidateUserData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userData models.UserData

		// Decode the request body into UserData
		err := json.NewDecoder(r.Body).Decode(&userData)
		if err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Error decoding user request body", err)
			return
		}

		// Validate UserData
		if err := validateUserData(userData); err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Validation error for user struct", err)
			return
		}

		// If validation passes, proceed to the next handler
		ctx := context.WithValue(r.Context(), utils.DECODED_USER_DATA, &userData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ValidatePasswordData is a middleware that validates PasswordData
func ValidatePasswordData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var passwordData models.PasswordData

		// Decode the request body into PasswordData
		err := json.NewDecoder(r.Body).Decode(&passwordData)
		if err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Error decoding password request body", err)
			return
		}

		// Validate PasswordData
		if err := validatePasswordData(passwordData); err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Validation error for password struct", err)
			return
		}

		// If validation passes, proceed to the next handler
		ctx := context.WithValue(r.Context(), utils.DECODED_PASSWORD_DATA, &passwordData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ValidateRoleData is a middleware that validates RoleData
func ValidateRoleData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var roleData models.RoleData

		// Decode the request body into RoleData
		err := json.NewDecoder(r.Body).Decode(&roleData)
		if err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Error decoding role request body", err)
			return
		}

		// Validate RoleData
		if err := validateRoleData(roleData); err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Validation error for role struct", err)
			return
		}

		// If validation passes, proceed to the next handler
		ctx := context.WithValue(r.Context(), utils.DECODED_ROLE_DATA, &roleData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// ValidateBlacklistData is a middleware that validates BlacklistData
func ValidateBlacklistData(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var blacklistData models.BlacklistData

		// Decode the request body into BlacklistData
		err := json.NewDecoder(r.Body).Decode(&blacklistData)
		if err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Error decoding blacklist request body", err)
			return
		}

		// Validate BlacklistData
		if err := validateBlacklistData(blacklistData); err != nil {
			logAndRespondWithError(w, http.StatusBadRequest, "Validation error for blacklist struct", err)
			return
		}

		// If validation passes, proceed to the next handler
		ctx := context.WithValue(r.Context(), utils.DECODED_BLACKLIST_DATA, &blacklistData)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
