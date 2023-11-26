package services

import (
	"fmt"
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

	// // In a real-world scenario, you might have more complex logic here.
	// participants := []models.Transactional{&participants.IDM{}, &participants.Patient{}, &participants.Doctor{}, &participants.Appointment{}, &participants.Consultation{}}

	// // Phase 1: Prepare Phase
	// prepareResponses := twophasecommit.SendPrepareMessage(participants)
	// if twophasecommit.AnyParticipantRespondedNo(prepareResponses) {
	// 	twophasecommit.SendAbortMessage(participants)
	// 	return fmt.Errorf("33")

	// }

	// // Phase 2: Commit Phase
	// commitResponses := twophasecommit.SendCommitMessage(participants)
	// if twophasecommit.AnyParticipantFailed(commitResponses) {
	// 	twophasecommit.SendRollbackMessage(participants)
	// 	return fmt.Errorf("33")
	// }

	// // Transaction successfully committed
	// twophasecommit.InformClient("clientID", "Delete operation completed successfully.")

	// Return nil to indicate that the message was successfully processed.
	return nil
}
