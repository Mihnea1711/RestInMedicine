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

	// Setup gRPC connection
	creds := insecure.NewCredentials()
	log.Printf("[GATEWAY] Initializing IDM client connection on %s:%d.", utils.IDM_HOST, utils.IDM_PORT)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", utils.IDM_HOST, utils.IDM_PORT), grpc.WithTransportCredentials(creds))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to IDM gRPC server: %v", err)
	}

	// Create IDM client
	app.idmClient = idm.NewIDMClient(conn)

	// setup router for the app
	router := routes.SetupRoutes(app.idmClient)
	app.router = router

	log.Println("[GATEWAY] Application successfully initialized.")
	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.Server.Port),
		Handler: a.router,
	}

	log.Println("[GATEWAY] Starting server...") // Logging the server start

	// Log the message just before starting the server in the goroutine
	fmt.Printf("[GATEWAY] Server started and listening on port %d\n", a.config.Server.Port)

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
