package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/consultatii/internal/database"
	"github.com/mihnea1711/POS_Project/services/consultatii/internal/routes"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/config"
	"github.com/redis/go-redis/v9"
)

type App struct {
	router   http.Handler
	database database.Database
	config   *config.AppConfig
}

func New(config *config.AppConfig) (*App, error) {
	app := &App{
		config: config,
	}

	// // setup mysql connection for the app
	// mysqlDB, err := mysql.NewMySQL(&config.MySQL)
	// if err != nil {
	// 	log.Printf("[CONSULTATIE] Error initializing MySQL: %v", err)
	// 	return nil, fmt.Errorf("failed to initialize MySQL: %w", err)
	// }
	// app.database = mysqlDB

	redis_addr := fmt.Sprintf("%s:%d", config.Redis.Host, config.Redis.Port)
	// setup redis and init cnnection
	rdb := redis.NewClient(&redis.Options{
		Addr:     redis_addr,            // Redis address
		Password: config.Redis.Password, // Password for db
		DB:       config.Redis.DB,       // Default DB
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("[CONSULTATIE] Failed to connect to Redis: %s", err)
	}

	// defer close redis conn func (nu e necesara pentru ping)
	/*
		// nu aici !! (o las pt ca ar putea fi folosita)
		defer func() {
			if err := a.rdb.Close(); err != nil {
				fmt.Println("[CONSULTATIE] Failed to close redis...", err)
			}
		}()
	*/

	// setup router for the app
	router := routes.SetupRoutes(app.database, rdb)
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
		if err := a.database.Close(); err != nil {
			log.Printf("[CONSULTATIE] Failed to close the database gracefully: %v", err)
		}

		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
