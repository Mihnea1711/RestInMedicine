package authorization

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/rabbit/internal/middleware"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/config"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/utils"
)

// RoleMiddleware is a middleware function that checks if the user has any of the required roles.
func RoleMiddleware(allowedRoles []string, jwtConfig config.JWTConfig, next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Extract and parse the JWT token from the Authorization header
		tokenString := middleware.ExtractJWTFromHeader(r)
		if tokenString == "" {
			log.Println("[GATEWAY_AUTH] An unexpected error occurred while trying to parse the token. Token is empty")
			utils.RespondWithJSON(w, http.StatusUnauthorized, models.ResponseData{
				Error:   "An unexpected error occurred while trying to parse the token. Token is empty",
				Message: "Access denied: An unexpected error occurred",
			})
			return
		}
		log.Println("sunt aici: " + tokenString)

		// Get the claims from the jwt
		claims, err := middleware.ParseJWT(tokenString, jwtConfig)
		if err != nil {
			log.Printf("[GATEWAY_AUTH] An unexpected error occurred while trying to parse the token: %v", err)
			utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
				Error:   fmt.Sprintf("An unexpected error occurred while trying to parse the token: %v", err),
				Message: "Access denied: An unexpected error occurred",
			})
			return
		}

		log.Printf("claims %v", claims)

		// Check if the user has any of the required roles
		userRole := claims.Role
		authorized := false
		for _, allowedRole := range allowedRoles {
			if userRole == allowedRole {
				authorized = true
				break
			}
		}

		log.Printf("Role %s", userRole)
		log.Printf("Auth: %v", authorized)

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

		// Store claims in the request context
		ctx := context.WithValue(r.Context(), utils.JWT_CLAIMS_CONTEXT_KEY, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	}
}

// AdminOnlyMiddleware is a middleware function that allows only admin access.
func AdminOnlyMiddleware(jwtConfig config.JWTConfig, next http.Handler) http.HandlerFunc {
	allowedRoles := []string{utils.ADMIN_ROLE}
	return RoleMiddleware(allowedRoles, jwtConfig, next)
}
