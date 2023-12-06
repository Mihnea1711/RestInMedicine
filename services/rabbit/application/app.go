package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mihnea1711/POS_Project/services/rabbit/idm"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models/participants"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/rabbitmq"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/routes"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/services"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/config"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/utils"
)

type App struct {
	webServer *http.Server
	idmClient idm.IDMClient
	rabbitMQ  *rabbitmq.RabbitMQ
	config    *config.AppConfig
}

func New(config *config.AppConfig, parentCtx context.Context) (*App, error) {
	app := &App{
		config: config,
	}

	// Initialize RabbitMQ connection
	rabbit, err := rabbitmq.NewRabbitMQ(config.RabbitMq)
	if err != nil {
		log.Printf("[RABBIT] Failed to initialize RabbitMQ: %v", err)
		return nil, fmt.Errorf("failed to initialize RabbitMQ: %w", err)
	}
	app.rabbitMQ = rabbit
	log.Println("[RABBIT] RabbitMQ connection successfully established.")

	// Setup router for the app with the rabbit in case the web part wants to publish to queues or smth
	router := routes.SetupRoutes(rabbit)
	log.Println("[RABBIT] HTTP Routes successfully loaded.")

	// Setup server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", app.config.Server.Port),
		Handler: router,
	}
	app.webServer = server
	log.Println("[RABBIT] Web Server successfully initialized.")

	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	// Start the web server concurrently
	webErrChan := make(chan error, 1)
	go func() {
		webErrChan <- a.StartWebServer()
	}()

	// Setup gRPC connection
	conn, err := a.SetupIDMClient()
	if err != nil {
		return fmt.Errorf("failed to connect to IDM gRPC server: %v", err)
	}
	// Manually close connection to IDM gRPC server
	defer conn.Close()
	log.Println("[RABBIT] IDM connection successfully established.")

	log.Println("[RABBIT] Preparing RabbitMQ server...")
	// Create the queues from the config file
	queues, err := rabbitmq.TranslateConfigToQueue(a.config.RabbitMq)
	if err != nil {
		log.Fatalf("[RABBIT] Failed to create queues from config: %v", err)
		return err
	}

	// the list should be dynamically loaded and should have the associated services unique uuids
	participantList := []models.Transactional{
		participants.NewIDM(uuid.New(), utils.IDM, a.idmClient),
		participants.NewPatient(uuid.New(), utils.PATIENT),
		participants.NewDoctor(uuid.New(), utils.DOCTOR),
	}
	// Store dependencies in the service container
	serviceContainer := &services.ServiceContainer{
		IDMClient:    a.idmClient,
		JWTConfig:    a.config.JWT,
		Participants: participantList,
	}

	// Set up the handlers for the queues with the service container
	a.rabbitMQ.SetupHandlers(serviceContainer)
	log.Println("[RABBIT] Rabbit handlers successfully established.")

	log.Printf("[RABBIT] RabbitMQ Server server started and listening on port %d...", a.config.RabbitMq.Port)
	// Declare and bind the queues
	err = a.rabbitMQ.SetupQueues(queues)
	if err != nil {
		log.Fatalf("[RABBIT] Failed to setup queues: %v", err)
		return err
	}

	// Wait for any service to finish or for the context to be done
	select {
	case err := <-webErrChan:
		if err != nil {
			return err
		}
	case <-ctx.Done():
		// Log the message indicating the server is in the process of shutting down
		log.Println("[RABBIT] Server shutting down...")

		timeout, cancel := context.WithTimeout(context.Background(), time.Second*utils.CLEAR_DB_RESOURCES_TIMEOUT)
		defer cancel()

		if err := a.webServer.Shutdown(timeout); err != nil {
			log.Printf("[RABBIT] Failed to shut down server gracefully: %v", err)
		} else {
			log.Println("[RABBIT] Server shut down gracefully.")
		}

		// Close the RabbitMQ connection
		if err := a.rabbitMQ.Close(timeout); err != nil {
			log.Fatalf("[RABBIT] Failed to close RabbitMQ: %v", err)
		}

		log.Println("[RABBIT] Server gracefully shut down")
		return nil
	}

	return nil
}
