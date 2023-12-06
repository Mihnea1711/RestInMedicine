package twophasecommit

import (
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models"
)

func SendPrepareMessage(participants []models.Transactional) ([]*models.ParticipantResponse, error) {
	var prepareResponses []*models.ParticipantResponse

	for _, participant := range participants {
		response, err := sendMessage(participant, "PREPARE", nil)
		if err != nil {
			log.Printf("[2PC] Error sending PREPARE message to participant %T: %v\n", participant, err)
			continue
		}
		prepareResponses = append(prepareResponses, response)
	}

	log.Println("[2PC] Sent PREPARE messages to all participants")
	return prepareResponses, nil
}

func SendAbortMessage(participants []models.Transactional) ([]*models.ParticipantResponse, error) {
	var abortResponses []*models.ParticipantResponse

	for _, participant := range participants {
		response, err := sendMessage(participant, "ABORT", nil)
		if err != nil {
			log.Printf("[2PC] Error sending ABORT message to participant %T: %v\n", participant, err)
			continue
		}
		abortResponses = append(abortResponses, response)
	}

	log.Println("[2PC] Sent ABORT messages to all participants")
	return abortResponses, nil
}

func SendCommitMessage(participants []models.Transactional, userID int) ([]*models.ParticipantResponse, error) {
	var commitResponses []*models.ParticipantResponse

	for _, participant := range participants {
		response, err := sendMessage(participant, "COMMIT", userID)
		if err != nil {
			log.Printf("[2PC] Error sending COMMIT message to participant %T: %v\n", participant, err)
			continue
		}
		commitResponses = append(commitResponses, response)
	}

	log.Println("[2PC] Sent COMMIT messages to all participants")
	return commitResponses, nil
}

func SendRollbackMessage(participants []models.Transactional) ([]*models.ParticipantResponse, error) {
	var rollbackResponses []*models.ParticipantResponse

	for _, participant := range participants {
		response, err := sendMessage(participant, "ROLLBACK", nil)
		if err != nil {
			log.Printf("[2PC] Error sending ROLLBACK message to participant %T: %v\n", participant, err)
			continue
		}
		rollbackResponses = append(rollbackResponses, response)
	}

	log.Println("[2PC] Sent ROLLBACK messages to all participants.")
	return rollbackResponses, nil
}

// sendMessage is a hypothetical function to send messages to a participant
func sendMessage(participant models.Transactional, messageType string, payload interface{}) (*models.ParticipantResponse, error) {
	var response *models.ParticipantResponse
	var err error

	switch messageType {
	case "PREPARE":
		response, err = participant.Prepare()
	case "COMMIT":
		response, err = participant.Commit(payload.(int))
	case "ROLLBACK", "ABORT":
		response, err = participant.Rollback()
	default:
		err = fmt.Errorf("unsupported message type: %s", messageType)
	}

	// Construct the response based on the outcome of the transactional methods
	if err != nil {
		log.Println("[2PC] An error occured while trying to send the message")
		return nil, err
	}

	return response, nil
}
