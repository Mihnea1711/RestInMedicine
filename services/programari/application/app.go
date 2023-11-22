package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/programari/internal/database"
	"github.com/mihnea1711/POS_Project/services/programari/internal/database/mysql"
	"github.com/mihnea1711/POS_Project/services/programari/internal/database/redis"
	"github.com/mihnea1711/POS_Project/services/programari/internal/routes"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/config"
	"github.com/mihnea1711/POS_Project/services/programari/pkg/utils"
)

type App struct {
	router   http.Handler
	database database.Database
	rdb      *redis.RedisClient
	config   *config.AppConfig
}

func New(config *config.AppConfig, parentCtx context.Context) (*App, error) {
	app := &App{
		config: config,
	}

	// Setup MySQL connection for the app
	mysqlDB, err := mysql.NewMySQL(parentCtx, &config.MySQL)
	if err != nil {
		log.Printf("[APPOINTMENT] Error initializing MySQL: %v", err)
		return nil, fmt.Errorf("failed to initialize MySQL: %w", err)
	}
	app.database = mysqlDB
	log.Println("[APPOINTMENT] MySQL connection successfully established.")

	// Setup Redis and initialize the connection
	rdb, err := redis.NewRedisClient(parentCtx, &config.Redis)
	if err != nil {
		log.Printf("[APPOINTMENT] Error initializing Redis: %v", err)
		return nil, fmt.Errorf("failed to initialize Redis: %w", err)
	}
	app.rdb = rdb
	log.Println("[APPOINTMENT] Redis connection successfully established.")

	// Setup router for the app
	router := routes.SetupRoutes(parentCtx, app.database, app.rdb)
	app.router = router

	log.Println("[APPOINTMENT] Application successfully initialized.")
	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.Server.Port),
		Handler: a.router,
	}

	log.Printf("[APPOINTMENT] Starting server on port %d...", a.config.Server.Port)

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
		// The second value is called open
		if !open {
			// Channel is closed
			log.Println("[APPOINTMENT] Context channel error. Channel is closed.")
		}
		return err
	case <-ctx.Done():
		// Log the message indicating the server is in the process of shutting down
		log.Println("[APPOINTMENT] Server shutting down...")

		// Close MySQL database connection gracefully
		if err := a.database.Close(); err != nil {
			log.Printf("[APPOINTMENT] Failed to close the MySQL database gracefully: %v", err)
		}

		if err := a.rdb.Close(); err != nil {
			fmt.Println("[APPOINTMENT] Failed to close Redis...", err)
		}

		timeout, cancel := context.WithTimeout(context.Background(), time.Second*utils.CLEAR_DB_RESOURCES_TIMEOUT)
		defer cancel()

		if err := server.Shutdown(timeout); err != nil {
			log.Printf("[APPOINTMENT] Failed to shut down server gracefully: %v", err)
		} else {
			log.Println("[APPOINTMENT] Server shut down gracefully.")
		}

		return nil
	}
}
