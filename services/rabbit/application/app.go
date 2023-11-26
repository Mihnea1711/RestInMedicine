package application

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/mihnea1711/POS_Project/services/rabbit/internal/rabbitmq"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/config"
)

type App struct {
	config   *config.AppConfig
	rabbitMQ *rabbitmq.RabbitMQ
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

	log.Println("[RABBIT] Application successfully initialized.")
	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	log.Printf("[RABBIT] Starting RabbitMQ server on port %d...", a.config.RabbitMq.Port)

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

	// Create a shutdown channel
	shutdownChan := make(chan struct{})

	// Use sync.WaitGroup for graceful shutdown
	var wg sync.WaitGroup

	// Start the goroutine to handle server shutdown
	wg.Add(1)
	go func() {
		defer wg.Done()

		// Wait for the context to be done
		<-ctx.Done()

		// Log the message indicating the server is in the process of shutting down
		log.Println("[RABBIT] Server shutting down...")

		// Close the shutdown channel to signal other goroutines to clean up
		close(shutdownChan)

		// Close the RabbitMQ connection
		if err := a.rabbitMQ.Close(ctx); err != nil {
			log.Fatalf("[RABBIT] Failed to close RabbitMQ: %v", err)
		}

		log.Println("[RABBIT] Server gracefully shut down")
	}()

	// Wait for all goroutines to finish before returning
	wg.Wait()

	return nil
}
