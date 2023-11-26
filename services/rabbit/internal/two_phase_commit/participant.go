package twophasecommit

import "github.com/mihnea1711/POS_Project/services/rabbit/internal/models"

// anyParticipantRespondedNo checks if any participant responded with "No" (transaction should be aborted)
func AnyParticipantRespondedNo(responses []*models.ParticipantResponse) bool {
	for _, response := range responses {
		if !response.Success {
			// Participant responded with "No"
			return true
		}
	}
	return false
}

// anyParticipantFailed checks if any participant failed (communication issue, internal error, etc.)
func AnyParticipantFailed(responses []*models.ParticipantResponse) bool {
	for _, response := range responses {
		if !response.Success {
			// Participant failed (communication issue, internal error, etc.)
			return true
		}
	}
	return false
}
