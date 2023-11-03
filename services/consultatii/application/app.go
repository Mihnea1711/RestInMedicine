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

	// Create a child context for MongoDB connection
	ctx, cancel := context.WithTimeout(parentCtx, 10*time.Second)
	defer cancel()

	// Create a MongoDB connection
	mongoDB, err := mongo.NewMongoDB(&config.Mongo, ctx)
	if err != nil {
		log.Printf("[CONSULTATIE] Error initializing MongoDB: %v", err)
		return nil, fmt.Errorf("failed to initialize MongoDB: %w", err)
	}
	app.database = mongoDB

	// Create a child context for Redis connection
	ctx, cancel = context.WithTimeout(parentCtx, 10*time.Second)
	defer cancel()
	// Create a Redis connection
	rdb, err := redis.NewRedisClient(&config.Redis, ctx)
	if err != nil {
		return nil, err
	}
	app.rdb = rdb

	// setup router for the app
	router := routes.SetupRoutes(app.database, app.rdb)
	app.router = router

	log.Println("[CONSULTATIE] Application successfully initialized.")
	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.Server.Port),
		Handler: a.router,
	}

	log.Println("[CONSULTATIE] Starting server...") // Logging the server start

	// Log the message just before starting the server in the goroutine
	fmt.Printf("[CONSULTATIE] Server started and listening on port %d\n", a.config.Server.Port)

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
			log.Println("[CONSULTATIE] Context channel error. Channel is closed.")
		}
		return err
	case <-ctx.Done():
		// Log the message indicating the server is in the process of shutting down
		log.Println("[CONSULTATIE] Server shutting down...")

		// Close database connection gracefully
		if err := a.database.Close(ctx); err != nil {
			log.Printf("[CONSULTATIE] Failed to close the database gracefully: %v", err)
		}

		if err := a.rdb.Close(); err != nil {
			fmt.Println("[CONSULTATIE] Failed to close redis gracefully...", err)
		}

		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
