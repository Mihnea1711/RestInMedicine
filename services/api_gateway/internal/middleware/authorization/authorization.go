package authorization

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/config"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

// RoleMiddleware is a middleware function that checks if the user has any of the required roles.
func RoleMiddleware(allowedRoles []string, jwtConfig config.JWTConfig, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract and parse the JWT token from the Authorization header
		tokenString := ExtractJWTFromHeader(r)
		if tokenString == "" {
			log.Println("[GATEWAY_AUTH] An unexpected error occurred while trying to parse the token. Token is empty")
			utils.RespondWithJSON(w, http.StatusUnauthorized, models.ResponseData{
				Error:   "An unexpected error occurred while trying to parse the token. Token is empty",
				Message: "Access denied: An unexpected error occurred",
			})
			return
		}

		// Get the claims from the jwt
		claims, err := ParseJWT(tokenString, jwtConfig)
		if err != nil {
			log.Printf("[GATEWAY_AUTH] An unexpected error occurred while trying to parse the token: %v", err)
			utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
				Error:   fmt.Sprintf("An unexpected error occurred while trying to parse the token: %v", err),
				Message: "Access denied: An unexpected error occurred",
			})
			return
		}

		// Check if the user has any of the required roles
		userRole := claims.Role
		authorized := false
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				authorized = true
				break
			}
		}

		if !authorized {
			err := "You don't have access to this resource"
			log.Printf("[GATEWAY_AUTH] Unauthorized: %v", err)
			utils.RespondWithJSON(w, http.StatusUnauthorized, models.ResponseData{
				Error:   fmt.Sprintf("Unauthorized: %v", err),
				Message: "Access denied",
			})
			return
		}

		log.Printf("[GATEWAY_AUTH] Successful access for user on: %s://%s%s?%s#%s",
			r.URL.Scheme,
			r.URL.Host,
			r.URL.Path,
			r.URL.RawQuery,
			r.URL.Fragment,
		)
		next.ServeHTTP(w, r)
	}
}

// AdminOnlyMiddleware is a middleware function that allows only admin access.
func AdminOnlyMiddleware(jwtConfig config.JWTConfig, next http.Handler) http.HandlerFunc {
	allowedRoles := []string{utils.ADMIN_ROLE}
	return RoleMiddleware(allowedRoles, jwtConfig, next)
}

// AdminAndPatientMiddleware is a middleware function that allows admin and patient access.
func AdminAndPatientMiddleware(jwtConfig config.JWTConfig, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the user has either admin or patient role
		tokenString := ExtractJWTFromHeader(r)
		if tokenString == "" {
			log.Println("[GATEWAY_AUTH] An unexpected error occurred while trying to parse the token. Token is empty")
			utils.RespondWithJSON(w, http.StatusUnauthorized, models.ResponseData{
				Error:   "An unexpected error occurred while trying to parse the token. Token is empty",
				Message: "Access denied: An unexpected error occurred",
			})
			return
		}

		// Get the claims from the jwt
		claims, err := ParseJWT(tokenString, jwtConfig)
		if err != nil {
			log.Printf("[GATEWAY_AUTH] An unexpected error occurred while trying to parse the token: %v", err)
			utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
				Error:   fmt.Sprintf("An unexpected error occurred while trying to parse the token: %v", err),
				Message: "Access denied: An unexpected error occurred",
			})
			return
		}

		userRole := claims.Role
		if userRole != utils.ADMIN_ROLE && userRole != utils.PATIENT_ROLE {
			err := "You don't have access to this resource"
			log.Printf("[GATEWAY_AUTH] Unauthorized: %v", err)
			utils.RespondWithJSON(w, http.StatusUnauthorized, models.ResponseData{
				Error:   fmt.Sprintf("Unauthorized: %v", err),
				Message: "Access denied: Unauthorized",
			})
			return
		}

		next.ServeHTTP(w, r)
	}
}

// AdminAndDoctorMiddleware is a middleware function that allows admin and doctor access.
func AdminAndDoctorMiddleware(jwtConfig config.JWTConfig, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Check if the user has either admin or doctor role
		tokenString := ExtractJWTFromHeader(r)
		if tokenString == "" {
			log.Println("[GATEWAY_AUTH] An unexpected error occurred while trying to parse the token. Token is empty")
			utils.RespondWithJSON(w, http.StatusUnauthorized, models.ResponseData{
				Error:   "An unexpected error occurred while trying to parse the token. Token is empty",
				Message: "Access denied: An unexpected error occurred",
			})
			return
		}

		// Get the claims from the jwt
		claims, err := ParseJWT(tokenString, jwtConfig)
		if err != nil {
			log.Printf("[GATEWAY_AUTH] An unexpected error occurred while trying to parse the token: %v", err)
			utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
				Error:   fmt.Sprintf("An unexpected error occurred while trying to parse the token: %v", err),
				Message: "Access denied: An unexpected error occurred",
			})
			return
		}

		userRole := claims.Role
		if userRole != utils.ADMIN_ROLE && userRole != utils.DOCTOR_ROLE {
			err := "You don't have access to this resource"
			log.Printf("[GATEWAY_AUTH] Unauthorized: %v", err)
			utils.RespondWithJSON(w, http.StatusUnauthorized, models.ResponseData{
				Error:   fmt.Sprintf("Unauthorized: %v", err),
				Message: "Access denied: Unauthorized",
			})
			return
		}

		next.ServeHTTP(w, r)
	}
}

// AllRolesMiddleware is a middleware function that allows admin, doctor, and patient access.
func AllRolesMiddleware(jwtConfig config.JWTConfig, next http.Handler) http.HandlerFunc {
	allowedRoles := []string{utils.ADMIN_ROLE, utils.DOCTOR_ROLE, utils.PATIENT_ROLE}
	return RoleMiddleware(allowedRoles, jwtConfig, next)
}
