package services

import "github.com/mihnea1711/POS_Project/services/rabbit/idm"

// ServiceContainer holds the dependencies needed by various services and handlers.
type ServiceContainer struct {
	IDMClient idm.IDMClient
	// Add more dependencies as needed
}
