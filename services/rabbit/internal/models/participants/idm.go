package participants

import "github.com/mihnea1711/POS_Project/services/rabbit/internal/models"

// IDM represents the IDM module
type IDM struct {
	models.ParticipantData // Embed Participant struct
	// Additional fields specific to IDM, if any
}

func NewIDM(id, name string) *IDM {
	return &IDM{
		ParticipantData: models.ParticipantData{
			ID:   id,
			Name: name,
		},
		// Initialize additional fields here
	}
}

// Override the Inform method for IDM
func (idm *IDM) Inform(commit bool) error {
	// Implement IDM-specific inform logic
	return nil
}

// Implement the Transactional interface methods for Participant
func (idm *IDM) Prepare() error {
	// Implement preparation logic
	return nil
}

func (idm *IDM) Commit() error {
	// Implement commit logic
	return nil
}

func (idm *IDM) Rollback() error {
	// Implement rollback logic
	return nil
}
