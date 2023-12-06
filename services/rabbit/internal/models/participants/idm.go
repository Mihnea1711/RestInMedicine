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

func NewIDM(participantID uuid.UUID, participantType models.ParticipantType, idmClient idm.IDMClient) *IDM {
	newIDM := &IDM{
		Participant: *models.NewParticipant(participantID, participantType),
		IDMClient:   idmClient,
	}

	log.Printf("New IDM registered - ID: %s\n", participantID.String())

	return newIDM
}

// Implement the Transactional interface methods for Participant
func (idm *IDM) Prepare() (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending IDM prepare request")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), utils.REQUEST_TIMEOUT_MULTIPLIER*time.Second)
	defer cancel()

	response, err := idm.IDMClient.HealthCheck(ctx, &proto_files.HealthCheckRequest{Service: "IDM"})
	if err != nil {
		log.Printf("Error making IDM prepare request: %v", err)
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

	// additionally check if user id to see if in exists before trying to delete smth that doesn t exist
	// although it should be handled ok by teh db
	// userIDResponse, err := idm.IDMClient.GetUserByID(ctx, &proto_files.UserIDRequest{UserID: &proto_files.UserID{ID: int64(userID)}})
	// if err != nil {
	// 	log.Printf("Error making IDM prepare request: %v", err)
	// 	return nil, err
	// }

	log.Printf("[2PC] IDM prepare request handled successfully. Status Code: %d", response.Status)

	return &models.ParticipantResponse{
		Code:    int(response.Status),
		Message: "IDM prepare request handled successfully",
	}, nil
}

func (idm *IDM) Commit(userID int) (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending IDM commit request")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), utils.REQUEST_TIMEOUT_MULTIPLIER*time.Second)
	defer cancel()

	response, err := idm.IDMClient.DeleteUserByID(ctx, &proto_files.UserIDRequest{UserID: &proto_files.UserID{ID: int64(userID)}})
	if err != nil {
		log.Printf("Error making IDM commit request: %v", err)
		return nil, err
	}

	log.Printf("[2PC] IDM commit request handled successfully. Status Code: %d, Rows Affected: %d", response.Info.Status, response.RowsAffected)

	return &models.ParticipantResponse{
		Code:    int(response.Info.Status),
		Message: fmt.Sprintf("%s. Rows Affected: %d", response.Info.Message, response.RowsAffected),
	}, nil
}

func (idm *IDM) Abort() (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending IDM abort request")

	return &models.ParticipantResponse{
		Code:    http.StatusOK,
		Message: "Transaction aborted successfully",
	}, nil
}

func (idm *IDM) Rollback(userID int) (*models.ParticipantResponse, error) {
	log.Println("[2PC] Sending IDM rollback request")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), utils.REQUEST_TIMEOUT_MULTIPLIER*time.Second)
	defer cancel()

	response, err := idm.IDMClient.RestoreUserByID(ctx, &proto_files.UserIDRequest{UserID: &proto_files.UserID{ID: int64(userID)}})
	if err != nil {
		log.Printf("Error making IDM rollback request: %v", err)
		return nil, err
	}

	log.Printf("[2PC] IDM rollback request handled successfully. Status Code: %d, Rows Affected: %d", response.Info.Status, response.RowsAffected)

	return &models.ParticipantResponse{
		Code:    int(response.Info.Status),
		Message: fmt.Sprintf("%s. Rows Affected: %d", response.Info.Message, response.RowsAffected),
	}, nil
}
