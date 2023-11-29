package twophasecommit

import (
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models"
)

// InformClient sends a message to the client with the specified outcome.
func InformClient(clientID string, outcome models.ClientResponse) error {
	// Use your messaging infrastructure to send a message to the client
	// Include the outcome (success or failure) and any other relevant information in the message

	// Example: Use a message broker, HTTP request, WebSocket, etc., to inform the client
	// Replace the following line with the actual code for your messaging mechanism.
	err := sendMessageToClient(clientID, outcome)
	if err != nil {
		log.Printf("Failed to inform client %s: %v", clientID, err)
		return fmt.Errorf("failed to inform client %s: %v", clientID, err)
	}

	log.Printf("Successfully informed client %s: Code: %d, Message: %s", clientID, outcome.Code, outcome.Message)
	return nil
}

// sendMessageToClient simulates sending a message to the client.
// In a real system, this function would use a messaging infrastructure.
func sendMessageToClient(clientID string, outcome models.ClientResponse) error {
	// Simulate sending a message to the client
	log.Printf("Sending message to client %s: Code: %d, Message: %s", clientID, outcome.Code, outcome.Message)
	return nil
}
