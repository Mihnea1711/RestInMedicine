package twophasecommit

import (
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models"
)

func SendPrepareMessage(participants []models.Transactional, msgWrapper models.MessageWrapper) ([]*models.ParticipantResponse, error) {
	var prepareResponses []*models.ParticipantResponse

	for _, participant := range participants {
		response, err := sendMessage(participant, "PREPARE", nil)
		if err != nil {
			log.Printf("[2PC] Error sending PREPARE message to participant %T for Transaction ID %s: %v\n", participant, msgWrapper.TransactionID, err)
			continue
		}
		log.Printf("[2PC] Participant %T responded to PREPARE with Status Code %d for Transaction ID %s\n", participant, response.Code, msgWrapper.TransactionID)
		prepareResponses = append(prepareResponses, response)
	}

	log.Printf("[2PC] Sent PREPARE messages to all participants for Transaction ID %s\n", msgWrapper.TransactionID)
	return prepareResponses, nil
}

func SendAbortMessage(participants []models.Transactional, msgWrapper models.MessageWrapper) ([]*models.ParticipantResponse, error) {
	var abortResponses []*models.ParticipantResponse

	for _, participant := range participants {
		response, err := sendMessage(participant, "ABORT", nil)
		if err != nil {
			log.Printf("[2PC] Error sending ABORT message to participant %T for Transaction ID %s: %v\n", participant, msgWrapper.TransactionID, err)
			continue
		}
		log.Printf("[2PC] Participant %T responded to ABORT with Status Code %d for Transaction ID %s\n", participant, response.Code, msgWrapper.TransactionID)
		abortResponses = append(abortResponses, response)
	}

	log.Printf("[2PC] Sent ABORT messages to all participants for Transaction ID %s\n", msgWrapper.TransactionID)
	return abortResponses, nil
}

func SendCommitMessage(participants []models.Transactional, msgWrapper models.MessageWrapper) ([]*models.ParticipantResponse, error) {
	var commitResponses []*models.ParticipantResponse

	for _, participant := range participants {
		response, err := sendMessage(participant, "COMMIT", msgWrapper.IDUser)
		if err != nil {
			log.Printf("[2PC] Error sending COMMIT message to participant %T for Transaction ID %s: %v\n", participant, msgWrapper.TransactionID, err)
			continue
		}
		log.Printf("[2PC] Participant %T responded to COMMIT with Status Code %d for Transaction ID %s\n", participant, response.Code, msgWrapper.TransactionID)
		commitResponses = append(commitResponses, response)
	}

	log.Printf("[2PC] Sent COMMIT messages to all participants for Transaction ID %s\n", msgWrapper.TransactionID)
	return commitResponses, nil
}

func SendRollbackMessage(participants []models.Transactional, msgWrapper models.MessageWrapper) ([]*models.ParticipantResponse, error) {
	var rollbackResponses []*models.ParticipantResponse

	for _, participant := range participants {
		response, err := sendMessage(participant, "ROLLBACK", msgWrapper.IDUser)
		if err != nil {
			log.Printf("[2PC] Error sending ROLLBACK message to participant %T for Transaction ID %s: %v\n", participant, msgWrapper.TransactionID, err)
			continue
		}
		log.Printf("[2PC] Participant %T responded to ROLLBACK with Status Code %d for Transaction ID %s\n", participant, response.Code, msgWrapper.TransactionID)
		rollbackResponses = append(rollbackResponses, response)
	}

	log.Printf("[2PC] Sent ROLLBACK messages to all participants for Transaction ID %s\n", msgWrapper.TransactionID)
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
	case "ABORT":
		response, err = participant.Abort()
	case "ROLLBACK":
		response, err = participant.Rollback(payload.(int))
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
