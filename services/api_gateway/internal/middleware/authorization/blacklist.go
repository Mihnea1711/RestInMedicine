package authorization

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mihnea1711/POS_Project/services/gateway/idm"
	"github.com/mihnea1711/POS_Project/services/gateway/idm/proto_files"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/models"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/config"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils/wrappers"
)

// BlacklistMiddleware is a middleware function that takes an idmClient as a parameter and checks if the user's jwt is in the blacklist or not
func BlacklistMiddleware(idmClient idm.IDMClient, jwtConfig config.JWTConfig) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Check if the request path is in the excluded paths
			for _, path := range utils.EXCLUDED_PATHS {
				// fmt.Println(r.URL.Path)
				if r.URL.Path == path {
					// If the path is excluded, skip the token check and proceed to the next handler
					next.ServeHTTP(w, r)
					return
				}
			}

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

			userIDString, err := claims.GetSubject()
			if err != nil {
				log.Printf("[GATEWAY_AUTH] Error getting user ID from claims: %v", err)
				utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
					Error:   fmt.Sprintf("Error getting user ID from claims: %v", err),
					Message: "Access denied: An unexpected error occurred",
				})
				return
			}

			userID, err := strconv.Atoi(userIDString)
			if err != nil {
				log.Printf("[GATEWAY_AUTH] Error converting user ID to integer: %v", err)
				utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
					Error:   fmt.Sprintf("Error converting user ID to integer: %v", err),
					Message: "Access denied: An unexpected error occurred",
				})
				return
			}
			response, err := idmClient.CheckUserInBlacklist(r.Context(), &proto_files.UserIDRequest{UserID: &proto_files.UserID{ID: int64(userID)}})
			if err != nil {
				log.Printf("[GATEWAY] CheckUserInBlacklist error: %v", err)
				utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
					Error:   fmt.Sprintf("CheckUserInBlacklist error: %v", err),
					Message: "Access denied: An unexpected error occurred",
				})
				return
			}
			infoResponse := &wrappers.InfoResponse{Response: response}
			utils.CheckNilResponse(w, http.StatusInternalServerError, "Check User in Blacklist response is nil", infoResponse.IsResponseNil, "Received nil response while checking the user in the blacklist.")
			utils.CheckNilResponse(w, http.StatusInternalServerError, "Check User in Blacklist response info is nil", infoResponse.IsInfoNil, "Received nil response.Info while checking the user in the blacklist.")

			// Check the gRPC response status and handle accordingly
			switch response.Info.Status {
			case http.StatusOK:
				log.Println("[GATEWAY] CheckBlacklist request handled successfully.")
				next.ServeHTTP(w, r)
				return
			case http.StatusNotFound:
				log.Println("[GATEWAY] User not found or no changes made.")
				utils.SendErrorResponse(w, http.StatusNotFound, response.Info.Message, "The user has not been found")
				return
			case http.StatusForbidden:
				log.Printf("[GATEWAY] User is in the blacklist: %v", response.Info.Message)
				utils.SendErrorResponse(w, http.StatusForbidden, response.Info.Message, "Access denied: User is in the blacklist")
				return
			default:
				log.Printf("[GATEWAY] Unexpected response status: %d", response.Info.Status)
				utils.SendErrorResponse(w, http.StatusInternalServerError, response.Info.Message, "Unexpected error occurred")
				return
			}
		})
	}
}
