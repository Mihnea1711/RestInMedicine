package application

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/rabbit/idm"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/rabbitmq"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/routes"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/config"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type App struct {
	router    http.Handler
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

	// setup router for the app
	router := routes.SetupRoutes()
	log.Println("[RABBIT] HTTP Routes successfully loaded.")

	app.router = router

	log.Println("[RABBIT] Application successfully initialized.")
	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	// log.Println("[RABBIT] Starting web server...")
	// server := &http.Server{
	// 	Addr:    fmt.Sprintf(":%d", a.config.Server.Port),
	// 	Handler: a.router,
	// }

	// Start the web server concurrently
	webErrChan := make(chan error, 1)
	go func() {
		webErrChan <- a.StartWebServer()
	}()

	// Setup gRPC connection
	creds := insecure.NewCredentials()
	log.Printf("[RABBIT] Initializing IDM client connection on %s:%d.", utils.IDM_HOST, utils.IDM_PORT)
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", utils.IDM_HOST, utils.IDM_PORT), grpc.WithTransportCredentials(creds))
	if err != nil {
		return fmt.Errorf("failed to connect to IDM gRPC server: %v", err)
	}
	// Manually close te connection to the IDM
	defer conn.Close()
	// Create IDM client
	a.idmClient = idm.NewIDMClient(conn)
	log.Println("[RABBIT] IDM connection successfully established.")

	// conn, err := a.SetupIDMClient()
	// if err != nil {
	// 	return err
	// }

	log.Println("[RABBIT] Preparing RabbitMQ server...")
	// Create the queues from the config file
	queues, err := rabbitmq.TranslateConfigToQueue(a.config.RabbitMq)
	if err != nil {
		log.Fatalf("[RABBIT] Failed to create queues from config: %v", err)
		return err
	}

	// Set up the handlers for the queues
	a.rabbitMQ.SetupHandlers()

	// Declare and bind the queues
	a.rabbitMQ.SetupQueues(queues)
	log.Printf("[RABBIT] RabbitMQ Server server started and listening on port %d...", a.config.RabbitMq.Port)

	// // Start RabbitMQ setup concurrently
	// rabbitErrChan := make(chan error, 1)
	// go func() {
	// 	rabbitErrChan <- a.SetupRabbitMQ(ctx)
	// }()

	// // Create a shutdown channel
	// shutdownChan := make(chan struct{})

	// // Use sync.WaitGroup for graceful shutdown
	// var wg sync.WaitGroup

	// // Start the goroutine to handle server shutdown
	// wg.Add(1)
	// go func() {
	// 	defer wg.Done()

	// 	// Wait for the context to be done
	// 	<-ctx.Done()

	// 	// Log the message indicating the server is in the process of shutting down
	// 	log.Println("[RABBIT] Server shutting down...")

	// 	// Close the shutdown channel to signal other goroutines to clean up
	// 	close(shutdownChan)

	// 	// Close the RabbitMQ connection
	// 	if err := a.rabbitMQ.Close(ctx); err != nil {
	// 		log.Fatalf("[RABBIT] Failed to close RabbitMQ: %v", err)
	// 	}

	// 	log.Println("[RABBIT] Server gracefully shut down")
	// }()

	// // Wait for all goroutines to finish before returning
	// wg.Wait()

	// return nil

	// Wait for any service to finish or for the context to be done
	select {
	case err := <-webErrChan:
		if err != nil {
			return err
		}
	// case err := <-rabbitErrChan:
	// 	if err != nil {
	// 		return err
	// 	}
	case <-ctx.Done():
		// Log the message indicating the server is in the process of shutting down
		log.Println("[RABBIT] Server shutting down...")

		// Close the RabbitMQ connection
		if err := a.rabbitMQ.Close(ctx); err != nil {
			log.Fatalf("[RABBIT] Failed to close RabbitMQ: %v", err)
		}

		// Manually close connection to IDM gRPC server
		defer conn.Close()

		log.Println("[RABBIT] Server gracefully shut down")
		return nil
	}

	return nil

}
