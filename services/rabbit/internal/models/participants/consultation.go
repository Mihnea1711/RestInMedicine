package participants

import "github.com/mihnea1711/POS_Project/services/rabbit/internal/models"

// IDM represents the IDM module
type Consultation struct {
	models.ParticipantData // Embed Participant struct
	// Additional fields specific to IDM, if any
}

func NewConsultation(id, name string) *Consultation {
	return &Consultation{
		ParticipantData: models.ParticipantData{
			ID:   id,
			Name: name,
		},
		// Initialize additional fields here
	}
}

// Override the Inform method for IDM
func (c *Consultation) Inform(commit bool) error {
	// Implement IDM-specific inform logic
	return nil
}

// Implement the Transactional interface methods for Participant
func (c *Consultation) Prepare() error {
	// Implement preparation logic
	return nil
}

func (c *Consultation) Commit() error {
	// Implement commit logic
	return nil
}

func (c *Consultation) Rollback() error {
	// Implement rollback logic
	return nil
}
