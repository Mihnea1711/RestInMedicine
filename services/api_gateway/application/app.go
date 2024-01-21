package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/gateway/idm"
	"github.com/mihnea1711/POS_Project/services/gateway/internal/routes"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/config"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	router    http.Handler
	idmClient idm.IDMClient
	config    *config.AppConfig
}

func New(config *config.AppConfig, parentCtx context.Context) (*App, error) {
	app := &App{
		config: config,
	}

	log.Println("[GATEWAY] Application successfully initialized.")
	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	log.Println("[GATEWAY] Starting server...")

	// Setup gRPC connection
	credentials := insecure.NewCredentials()
	log.Printf("[GATEWAY] Initializing IDM client connection on %s:%d.", utils.IDM_HOST, utils.IDM_PORT)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", utils.IDM_HOST, utils.IDM_PORT), grpc.WithTransportCredentials(credentials))
	if err != nil {
		return fmt.Errorf("failed to connect to IDM gRPC server: %v", err)
	}
	// Manually close te connection to the IDM
	defer func(conn *grpc.ClientConn) {
		err := conn.Close()
		if err != nil {
			log.Println("[GATEWAY] Connection error. Connection did not close.")
		}
	}(conn)
	// Create IDM client
	a.idmClient = idm.NewIDMClient(conn)

	// setup router for the app
	router := routes.SetupRoutes(a.idmClient, a.config.JWT)
	a.router = router

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.Server.Port),
		Handler: a.router,
	}

	// Log the message just before starting the server in the goroutine
	log.Printf("[GATEWAY] Server started and listening on port %d\n", a.config.Server.Port)

	channel := make(chan error, 1)
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			channel <- fmt.Errorf("failed to start server: %w", err)
		}
		close(channel)
	}()

	select {
	case err, open := <-channel:
		// second value is called open
		if !open {
			//channel is closed
			log.Println("[GATEWAY] Context channel error. Channel is closed.")
		}
		return err
	case <-ctx.Done():
		// Log the message indicating the server is in the process of shutting down
		log.Println("[GATEWAY] Server shutting down...")

		timeout, cancel := context.WithTimeout(context.Background(), time.Second*utils.DB_CLEAR_RESOURCES_MULTIPLIER)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
