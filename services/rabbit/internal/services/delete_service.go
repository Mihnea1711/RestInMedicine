package services

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/rabbit/internal/middleware/authorization"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models/participants"
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

	// parse jwt and roles
	// Authorize the client based on the JWT in the message
	messageData, err := authorization.AuthorizeClient(message, s.JWTConfig)
	if err != nil {
		// Log the authorization error
		log.Printf("[RABBIT] Authorization error while deleting user: %v", err)
		twophasecommit.InformClient("clientID", models.ClientResponse{
			Code:    http.StatusUnauthorized,
			Message: "Authorization error",
		})
		return err
	}

	// t id
	_ = utils.StartTransaction()

	participants := []models.Transactional{&participants.IDM{IDMClient: s.IDMClient}, &participants.Patient{}, &participants.Doctor{}}

	// Phase 1: Prepare Phase
	prepareResponses, err := twophasecommit.SendPrepareMessage(participants)
	if err != nil {
		log.Printf("[2PC] Error in Prepare Phase: %v", err)
		return err
	}
	if twophasecommit.AnyParticipantRespondedNo(prepareResponses) {
		log.Println("[2PC] Prepare Phase: One or more participants responded with 'NO'")
		twophasecommit.SendAbortMessage(participants)
		return fmt.Errorf("prepare phase failed: One or more participants responded with 'NO'")
	}

	// Phase 2: Commit Phase
	commitResponses, err := twophasecommit.SendCommitMessage(participants, messageData.IDUser)
	if err != nil {
		log.Printf("[2PC] Error in Commit Phase: %v", err)
		return err
	}
	if twophasecommit.AnyParticipantFailed(commitResponses) {
		log.Println("[2PC] Commit Phase: One or more participants failed")
		twophasecommit.SendRollbackMessage(participants)
		return fmt.Errorf("commit phase failed: One or more participants failed")
	}

	// Transaction successfully committed
	log.Println("[2PC] Transaction successfully committed")
	twophasecommit.InformClient("clientID", models.ClientResponse{
		Code:    http.StatusOK,
		Message: "Delete operation completed successfully.",
	})

	// Return nil to indicate that the message was successfully processed.
	return nil
}
