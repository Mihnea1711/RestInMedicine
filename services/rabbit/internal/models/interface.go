package models

// Transactional defines the methods that a participant in a distributed transaction should implement
type Transactional interface {
	// Prepare is called to prepare the participant for the transaction
	Prepare() (*ParticipantResponse, error)

	// Commit is called to commit the changes made during the transaction
	Commit(userId int) (*ParticipantResponse, error)

	// Rollback is called to undo the changes made during the transaction
	Rollback() (*ParticipantResponse, error)

	// Inform is called to inform the participant about the outcome of the transaction
	Inform(commit bool) (*ParticipantResponse, error)

	// Compensate is called to undo changes made during the transaction
	Compensate() error
}
