package services

import (
	"github.com/mihnea1711/POS_Project/services/rabbit/idm"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/config"
)

// ServiceContainer holds the dependencies needed by various services and handlers.
type ServiceContainer struct {
	IDMClient idm.IDMClient
	JWTConfig config.JWTConfig
	// Add more dependencies as needed

	// maybe add here list of participants
	// have a queue that listens to gateway sending info about current participants
}
