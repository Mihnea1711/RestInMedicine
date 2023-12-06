package twophasecommit

import (
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models"
)

// AnyParticipantRespondedNo checks if any participant responded with "No" (transaction should be aborted)
func AnyParticipantRespondedNo(responses []*models.ParticipantResponse, msgWrapper models.MessageWrapper) bool {
	for _, response := range responses {
		if response.Code != http.StatusOK {
			// Participant responded with "No"
			log.Printf("[2PC] Participant responded with 'No' for Transaction ID %s. Participant: %T, Code: %d, Message: %s\n", msgWrapper.TransactionID, response.ID, response.Code, response.Message)
			return true
		}
	}
	return false
}

// AnyParticipantFailed checks if any participant failed (communication issue, internal error, etc.)
func AnyParticipantFailed(responses []*models.ParticipantResponse, msgWrapper models.MessageWrapper) bool {
	for _, response := range responses {
		if response.Code != http.StatusOK && response.Code != http.StatusNotFound {
			// Participant failed (communication issue, internal error, etc.)
			log.Printf("[2PC] Participant failed for Transaction ID %s. Participant: %T, Code: %d, Message: %s\n", msgWrapper.TransactionID, response.ID, response.Code, response.Message)
			return true
		}
	}
	return false
}
