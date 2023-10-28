package application

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/mihnea1711/POS_Project/services/doctori/internal/routes"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/config"
)

type App struct {
	router http.Handler
	// rdb    *redis.Client
	config *config.AppConfig
}

func New(config *config.AppConfig) *App {
	app := &App{
		config: config,
	}

	router := routes.SetupRoutes()
	app.router = router

	return app
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

	fmt.Println("Starting server..")

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
			fmt.Println("context channel error. channel is closed")
		}
		return err
	case <-ctx.Done():
		// Log the message indicating the server is in the process of shutting down
		fmt.Println("\nServer shutting down...")

		timeout, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		return server.Shutdown(timeout)
	}
}
