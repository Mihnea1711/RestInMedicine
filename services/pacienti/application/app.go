package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/pacienti/internal/database"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/database/mysql"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/database/redis"
	"github.com/mihnea1711/POS_Project/services/pacienti/internal/routes"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/config"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
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

	// setup mysql connection for the app
	mysqlDB, err := mysql.NewMySQL(parentCtx, &config.MySQL)
	if err != nil {
		log.Printf("[PATIENT] Error initializing MySQL: %v", err)
		return nil, fmt.Errorf("failed to initialize MySQL: %w", err)
	}
	app.database = mysqlDB

	// setup redis and init cnnection
	rdb, err := redis.NewRedisClient(parentCtx, &config.Redis)
	if err != nil {
		return nil, err
	}
	app.rdb = rdb

	// setup router for the app
	router := routes.SetupRoutes(parentCtx, app.database, app.rdb)
	app.router = router

	log.Println("[PATIENT] Application successfully initialized.")
	return app, nil
}

func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", a.config.Server.Port),
		Handler: a.router,
	}

	log.Println("[PATIENT] Starting server...") // Logging the server start

	// Log the message just before starting the server in the goroutine
	fmt.Printf("[PATIENT] Server started and listening on port %d\n", a.config.Server.Port)

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
			log.Println("[PATIENT] Context channel error. Channel is closed.")
		}
		return err
	case <-ctx.Done():
		// Log the message indicating the server is in the process of shutting down
		log.Println("[PATIENT] Server shutting down...")

		// Close MySQL database connection gracefully
		if err := a.database.Close(); err != nil {
			log.Printf("[PATIENT] Failed to close the MySQL database gracefully: %v", err)
		}

		if err := a.rdb.Close(); err != nil {
			fmt.Println("[PATIENT] Failed to close redis...", err)
		}

		timeout, cancel := context.WithTimeout(context.Background(), time.Second*utils.CLEAR_DB_RESOURCES_TIMEOUT)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
