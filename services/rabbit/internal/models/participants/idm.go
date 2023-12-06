package participants

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mihnea1711/POS_Project/services/rabbit/idm"
	"github.com/mihnea1711/POS_Project/services/rabbit/idm/proto_files"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/utils"
)

// IDM represents the IDM module
type IDM struct {
	models.Participant               // Embed Participant struct
	IDMClient          idm.IDMClient // Additional fields specific to IDM, if any
}

func NewIDM(participantID uuid.UUID, participantType models.ParticipantType) *IDM {
	return &IDM{
		Participant: *models.NewParticipant(participantID, participantType),
	}
}

// Implement the Transactional interface methods for Participant
func (idm *IDM) Prepare() (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending idm prepare request")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), utils.REQUEST_TIMEOUT_MULTIPLIER*time.Second)
	defer cancel()

	response, err := idm.IDMClient.HealthCheck(ctx, &proto_files.HealthCheckRequest{Service: "IDM"})
	if err != nil {
		log.Printf("Error making idm prepare request: %v", err)
		return nil, err
	}

	switch response.Status {
	case proto_files.HealthCheckResponse_SERVING:
		response.Status = http.StatusOK
	case proto_files.HealthCheckResponse_NOT_SERVING:
		response.Status = http.StatusServiceUnavailable
	case proto_files.HealthCheckResponse_UNKNOWN:
		response.Status = http.StatusInternalServerError
	}

	return &models.ParticipantResponse{
		Code:    int(response.Status),
		Message: "IDM prepare request handled successfully",
	}, nil
}

func (idm *IDM) Commit(userID int) (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending idm commit request")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), utils.REQUEST_TIMEOUT_MULTIPLIER*time.Second)
	defer cancel()

	response, err := idm.IDMClient.DeleteUserByID(ctx, &proto_files.UserIDRequest{UserID: &proto_files.UserID{ID: int64(userID)}})
	if err != nil {
		log.Printf("Error making idm prepare request: %v", err)
		return nil, err
	}

	return &models.ParticipantResponse{
		Code:    int(response.Info.Status),
		Message: fmt.Sprintf("%s. Rows Affected: %d", response.Info.Message, response.RowsAffected),
	}, nil
}

func (idm *IDM) Abort() (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending idm abort request")

	return &models.ParticipantResponse{
		Code:    http.StatusOK,
		Message: "Transaction aborted successfully",
	}, nil
}

func (idm *IDM) Rollback() (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending idm rollback request")

	return &models.ParticipantResponse{
		Code:    http.StatusOK,
		Message: "Transaction rolled back successfully",
	}, nil
}

func (idm *IDM) Compensate() error {
	return nil
}
