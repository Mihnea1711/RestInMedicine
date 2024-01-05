package authorization

// func AuthorizeClient(message []byte, jwtConfig config.JWTConfig) (*models.DeleteMessageData, error) {
// 	// Unmarshal the message into DeleteMessageData struct
// 	var deleteMessageData *models.DeleteMessageData
// 	if err := json.Unmarshal(message, &deleteMessageData); err != nil {
// 		log.Printf("[RABBIT_AUTH] Failed to unmarshal message: %v", err)
// 		return nil, fmt.Errorf("failed to unmarshal message: %w", err)
// 	}

// 	// Parse and validate the JWT
// 	claims, err := middleware.ParseJWT(deleteMessageData.JWT, jwtConfig)
// 	if err != nil {
// 		log.Printf("[RABBIT_AUTH] Failed to parse JWT: %v", err)
// 		return nil, fmt.Errorf("failed to parse JWT: %w", err)
// 	}

// 	// Check the user's role for authorization
// 	if claims.Role == utils.ADMIN_ROLE {
// 		log.Printf("[RABBIT_AUTH] User with JWT %s authorized as admin", deleteMessageData.JWT)
// 		return deleteMessageData, nil
// 	}

// 	// If the role is not "admin," return false
// 	log.Printf("[RABBIT_AUTH] User with JWT %s is not authorized as admin", deleteMessageData.JWT)
// 	return nil, fmt.Errorf("not authorized")
// }
