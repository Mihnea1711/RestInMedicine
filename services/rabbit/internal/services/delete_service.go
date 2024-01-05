package services

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models"
	twophasecommit "github.com/mihnea1711/POS_Project/services/rabbit/internal/two_phase_commit"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/utils"
)

// DeleteUserMessageHandler is an example implementation of MessageHandler.
// This function represents the message handling logic for a "delete" message.
// In the context of a distributed system, such as a message queue like RabbitMQ,
// the handler may participate in a distributed transaction that involves multiple
// systems or services. This is where concepts like Two-Phase Commit (2PC) may come into play.

// Assuming a distributed transaction context:
// 1. This function receives a "delete" message as a byte slice.
// 2. It may perform local operations related to handling the delete message.
//    This could involve deleting a user from a local database, updating local state, etc.

// 3. In a more complex distributed system, if the delete operation involves interactions
//    with other services or systems, this is where a distributed transaction protocol
//    like Two-Phase Commit (2PC) could be initiated.

// 4. The Prepare Phase:
//    - The handler may check if it is ready to commit the local changes.
//    - It might interact with other services to ensure they are also ready for the commit.

// 5. The Commit Phase:
//    - If all participants are ready, the handler may commit the local changes.
//    - If any participant indicates a problem, the handler may decide to roll back
//      the changes and take appropriate compensating actions.

// This is a placeholder implementation, and the actual logic would depend on the
// requirements and architecture of the distributed system.

func (s *ServiceContainer) DeleteUserMessageHandler(message []byte) error {
	fmt.Printf("[RABBIT] Handling delete message: %s\n", string(message))

	// t id
	transactionID := utils.StartTransaction()
	messageWrapper := models.MessageWrapper{
		TransactionID: transactionID,
	}

	// // parse jwt and roles
	// // Authorize the client based on the JWT in the message
	// messageData, err := authorization.AuthorizeClient(message, s.JWTConfig)
	// if err != nil {
	// 	// Log the authorization error
	// 	log.Printf("[RABBIT] Authorization error while deleting user: %v", err)
	// 	twophasecommit.InformClient("clientID", messageWrapper, models.ClientResponse{
	// 		Code:    http.StatusUnauthorized,
	// 		Message: "Authorization error",
	// 	})
	// 	return err
	// }

	var deleteMessageData models.DeleteMessageData
	if err := json.Unmarshal(message, &deleteMessageData); err != nil {
		log.Printf("[2PC] Failed to unmarshal message: %v", err)
		// Inform client about the error
		twophasecommit.InformClient("clientID", messageWrapper, models.ClientResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error during message decoding.",
		})

		return err
	}
	messageWrapper.IDUser = deleteMessageData.IDUser

	// hardcoded list of participants
	// participants := []models.Transactional{&participants.IDM{IDMClient: s.IDMClient}, &participants.Patient{}, &participants.Doctor{}}

	// Phase 1: Prepare Phase
	prepareResponses, err := twophasecommit.SendPrepareMessage(s.Participants, messageWrapper)
	if err != nil {
		log.Printf("[2PC] Error in Prepare Phase. Transaction ID: %s, Error: %v", transactionID, err)

		// Inform client about the error
		twophasecommit.InformClient("clientID", messageWrapper, models.ClientResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error during transaction preparation.",
		})

		return err
	}
	if twophasecommit.AnyParticipantRespondedNo(prepareResponses, messageWrapper) {
		log.Println("[2PC] Prepare Phase: One or more participants responded with 'NO'")
		twophasecommit.SendAbortMessage(s.Participants, messageWrapper)

		// Inform client about the error
		twophasecommit.InformClient("clientID", messageWrapper, models.ClientResponse{
			Code:    http.StatusServiceUnavailable,
			Message: "Transaction aborted due to participant response with 'NO'.",
		})

		return fmt.Errorf("prepare phase failed: One or more participants responded with 'NO'")
	}

	// Phase 2: Commit Phase
	commitResponses, err := twophasecommit.SendCommitMessage(s.Participants, messageWrapper)
	if err != nil {
		log.Printf("[2PC] Error in Commit Phase. Transaction ID: %s, Error: %v", transactionID, err)

		// Inform client about the error
		twophasecommit.InformClient("clientID", messageWrapper, models.ClientResponse{
			Code:    http.StatusInternalServerError,
			Message: "Internal Server Error during transaction commit.",
		})

		return err
	}
	if twophasecommit.AnyParticipantFailed(commitResponses, messageWrapper) {
		log.Println("[2PC] Commit Phase: One or more participants failed")
		twophasecommit.SendRollbackMessage(s.Participants, messageWrapper)

		// Inform client about the error
		twophasecommit.InformClient("clientID", messageWrapper, models.ClientResponse{
			Code:    http.StatusFailedDependency,
			Message: "Transaction rolled back due to participant failure.",
		})

		return fmt.Errorf("commit phase failed: One or more participants failed")
	}

	// Transaction successfully committed
	log.Printf("[2PC] Transaction successfully committed. Transaction ID: %s", transactionID)
	twophasecommit.InformClient("clientID", messageWrapper, models.ClientResponse{
		Code:    http.StatusOK,
		Message: "Delete operation completed successfully.",
	})

	// Return nil to indicate that the message was successfully processed.
	return nil
}
