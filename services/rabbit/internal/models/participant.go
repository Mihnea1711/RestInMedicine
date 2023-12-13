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
