package application

import (
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/rabbit/idm"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// startWebServer starts the HTTP server
func (a *App) StartWebServer() error {
	log.Println("[RABBIT] Starting web server...")

	// Log the message just before starting the server in the goroutine
	log.Printf("[RABBIT] Web Server started and listening on port %d\n", a.config.Server.Port)
	channel := make(chan error, 1)
	go func() {
		err := a.webServer.ListenAndServe()
		if err != nil {
			channel <- fmt.Errorf("failed to start server: %w", err)
		}
		close(channel)
	}()

	return <-channel
}

func (a *App) SetupIDMClient() (*grpc.ClientConn, error) {
	creds := insecure.NewCredentials()
	log.Printf("[RABBIT] Initializing IDM client connection on %s:%d.", utils.IDM_HOST, utils.IDM_PORT)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", utils.IDM_HOST, utils.IDM_PORT), grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to IDM gRPC server: %v", err)
	}
	// Create IDM client
	a.idmClient = idm.NewIDMClient(conn)
	return conn, nil
}

func NewIDMConnection() (*grpc.ClientConn, error) {
	// Setup gRPC connection
	creds := insecure.NewCredentials()
	log.Printf("[RABBIT] Initializing IDM client connection on %s:%d.", utils.IDM_HOST, utils.IDM_PORT)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", utils.IDM_HOST, utils.IDM_PORT), grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to IDM gRPC server: %v", err)
	}
	return conn, nil
}
