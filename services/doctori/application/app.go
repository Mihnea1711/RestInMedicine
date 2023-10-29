package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/doctori/internal/database"
	"github.com/mihnea1711/POS_Project/services/doctori/internal/routes"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/config"
)

type App struct {
	router  http.Handler
	mysqlDB *database.MySQLDatabase
	// rdb    *redis.Client
	config *config.AppConfig
}

func New(config *config.AppConfig) (*App, error) {
	app := &App{
		config: config,
	}

	// setup mysql connection for the app
	mysqlDB, err := database.NewMySQL(&config.MySQL)
	if err != nil {
		log.Printf("Error initializing MySQL: %v", err)
		return nil, fmt.Errorf("failed to initialize MySQL: %w", err)
	}
	app.mysqlDB = mysqlDB

	// setup router for the app
	router := routes.SetupRoutes(app.mysqlDB)
	app.router = router

	log.Println("Application successfully initialized.")
	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.Server.Port),
		Handler: a.router,
	}

	// err := a.rdb.Ping(ctx).Err()
	// if err != nil {
	// 	return fmt.Errorf("failed to connect to redis: %w", err)
	// }

	// defer func() {
	// 	if err := a.rdb.Close(); err != nil {
	// 		fmt.Println("failed to close redis...", err)
	// 	}
	// }()

	log.Println("Starting server...") // Logging the server start

	// Log the message just before starting the server in the goroutine
	fmt.Printf("Server started and listening on port %d\n", a.config.Server.Port)

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
			log.Println("Context channel error. Channel is closed.")
		}
		return err
	case <-ctx.Done():
		// Log the message indicating the server is in the process of shutting down
		log.Println("Server shutting down...")

		// Close MySQL database connection gracefully
		if err := a.mysqlDB.Close(); err != nil {
			log.Printf("Failed to close the MySQL database gracefully: %v", err)
		}

		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
