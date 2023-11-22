package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/database"
	"github.com/mihnea1711/POS_Project/services/consultatii/internal/database/mongo"
	"github.com/mihnea1711/POS_Project/services/consultatii/internal/database/redis"
	"github.com/mihnea1711/POS_Project/services/consultatii/internal/routes"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/config"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
)

type App struct {
	router   http.Handler
	database database.Database
	config   *config.AppConfig
	rdb      *redis.RedisClient
}

func New(config *config.AppConfig, parentCtx context.Context) (*App, error) {
	app := &App{
		config: config,
	}

	// Create a MongoDB connection
	mongoDB, err := mongo.NewMongoDB(parentCtx, &config.Mongo)
	if err != nil {
		log.Printf("[CONSULTATION] Error initializing MongoDB: %v", err)
		return nil, fmt.Errorf("failed to initialize MongoDB: %w", err)
	}
	app.database = mongoDB
	log.Println("[CONSULTATION] Mongo connection successfully established.")

	// Create a Redis connection
	rdb, err := redis.NewRedisClient(parentCtx, &config.Redis)
	if err != nil {
		log.Printf("[CONSULTATION] Error initializing Redis: %v", err)
		return nil, err
	}
	app.rdb = rdb
	log.Println("[CONSULTATION] Redis connection successfully established.")

	// setup router for the app
	router := routes.SetupRoutes(parentCtx, app.database, app.rdb)
	app.router = router

	log.Println("[CONSULTATION] Application successfully initialized.")
	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.Server.Port),
		Handler: a.router,
	}

	// Log the message just before starting the server in the goroutine
	fmt.Printf("[CONSULTATION] Server started and listening on port %d\n", a.config.Server.Port)

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
			log.Println("[CONSULTATION] Context channel error. Channel is closed.")
		}
		return err
	case <-ctx.Done():
		// Log the message indicating the server is in the process of shutting down
		log.Println("[CONSULTATION] Server shutting down...")

		// Close database connection gracefully
		if err := a.database.Close(ctx); err != nil {
			log.Printf("[CONSULTATION] Failed to close the database gracefully: %v", err)
		}

		if err := a.rdb.Close(); err != nil {
			fmt.Println("[CONSULTATION] Failed to close redis gracefully...", err)
		}

		timeout, cancel := context.WithTimeout(context.Background(), time.Second*utils.RESOURCES_CLOSE_TIMEOUT)
		defer cancel()

		if err := server.Shutdown(timeout); err != nil {
			log.Printf("[CONSULTATION] Failed to shut down server gracefully: %v", err)
			return err
		}

		log.Println("[CONSULTATION] Server shut down gracefully.")
		return nil
	}
}
