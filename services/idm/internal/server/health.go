package server

import (
	"context"
	"log"

	"github.com/mihnea1711/POS_Project/services/idm/idm/proto_files"
)

func (s *MyIDMServer) HealthCheck(ctx context.Context, req *proto_files.HealthCheckRequest) (*proto_files.HealthCheckResponse, error) {
	status := proto_files.HealthCheckResponse_SERVING
	log.Printf("Health check for service %s returned status: %v", req.Service, status)

	return &proto_files.HealthCheckResponse{
		Status: status,
	}, nil
}
