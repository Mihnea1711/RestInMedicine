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
		log.Printf("[APPOINTMENT] Error initializing MySQL: %v", err)
		return nil, fmt.Errorf("failed to initialize MySQL: %w", err)
	}
	app.database = mysqlDB

	// setup redis and init cnnection
	rdb, err := redis.NewRedisClient(parentCtx, &config.Redis)
	if err != nil {
		log.Fatalf("[APPOINTMENT] Failed to connect to Redis: %s", err)
	}
	app.rdb = rdb

	// setup router for the app
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

	log.Println("[APPOINTMENT] Starting server...") // Logging the server start

	// Log the message just before starting the server in the goroutine
	fmt.Printf("[APPOINTMENT] Server started and listening on port %d\n", a.config.Server.Port)

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
			fmt.Println("[APPOINTMENT] Failed to close redis...", err)
		}

		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
