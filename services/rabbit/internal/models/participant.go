package models

import "github.com/google/uuid"

type ParticipantType string

// Participant represents a module that participates in a distributed transaction
type Participant struct {
	_id   uuid.UUID
	_type ParticipantType
}

func NewParticipant(participantID uuid.UUID, particpantType ParticipantType) *Participant {
	return &Participant{
		_id:   participantID,
		_type: particpantType,
	}
}

func (p *Participant) GetID() uuid.UUID {
	return p._id
}

func (p *Participant) GetType() ParticipantType {
	return p._type
}

// Override the Inform method for IDM
func (p *Participant) Inform(commit bool) (*ParticipantResponse, error) {
	// Implement inform logic
	return nil, nil
}

// Implement the Transactional interface methods for Participant
func (p *Participant) Prepare() (*ParticipantResponse, error) {
	// Implement preparation logic
	return nil, nil
}

func (p *Participant) Commit() (*ParticipantResponse, error) {
	// Implement commit logic
	return nil, nil
}

func (p *Participant) Rollback() (*ParticipantResponse, error) {
	// Implement rollback logic
	return nil, nil
}
