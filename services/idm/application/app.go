package application

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/mihnea1711/POS_Project/services/idm/idm"
	"github.com/mihnea1711/POS_Project/services/idm/internal/database"
	"github.com/mihnea1711/POS_Project/services/idm/internal/database/mysql"
	"github.com/mihnea1711/POS_Project/services/idm/internal/database/redis"
	"github.com/mihnea1711/POS_Project/services/idm/internal/middleware"
	"github.com/mihnea1711/POS_Project/services/idm/internal/server"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/config"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
	"google.golang.org/grpc"
)

type App struct {
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
		log.Printf("[IDM] Error initializing MySQL: %v", err)
		return nil, fmt.Errorf("failed to initialize MySQL: %w", err)
	}
	app.database = mysqlDB
	log.Println("[IDM] MySQL connection successfully established.")

	// setup redis connection
	rdb, err := redis.NewRedisClient(parentCtx, &config.Redis)
	if err != nil {
		return nil, err
	}
	app.rdb = rdb
	log.Println("[IDM] Redis connection successfully established.")

	return app, nil
}

// Start initializes and starts the IDM service.
func (a *App) Start(ctx context.Context) error {
	// Create gRPC listener
	grpcListener, err := net.Listen("tcp", fmt.Sprintf(":%d", a.config.Server.Port))
	if err != nil {
		log.Fatalf("[IDM] Failed to listen: %v", err)
	}

	// Create gRPC server
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(middleware.ValidateRequestInterceptor),
	)
	grpcService := &server.MyIDMServer{
		DbConn:    a.database,
		RedisConn: a.rdb,
		JwtConfig: a.config.JWT,
	}
	idm.RegisterIDMServer(grpcServer, grpcService)

	log.Println("[IDM] Starting server...") // Log the start of the server

	channel := make(chan error, 1)
	go func() {
		log.Printf("[IDM] gRPC server started and listening on port %d\n", a.config.Server.Port)
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Fatalf("[IDM] Failed to serve gRPC server: %v", err)
		}
		close(channel)
	}()

	select {
	case err, open := <-channel:
		if !open {
			// The channel is closed
			log.Println("[IDM] Context channel error. Channel is closed.")
		}
		return err
	case <-ctx.Done():
		// Log the message indicating the server is in the process of shutting down
		log.Println("[IDM] Server shutting down...")

		// Close MySQL database connection gracefully
		if err := a.database.Close(); err != nil {
			log.Printf("[IDM] Failed to close the MySQL database gracefully: %v", err)
		}

		if err := a.rdb.Close(); err != nil {
			log.Printf("[IDM] Failed to close Redis gracefully: %v", err)
		}

		// Allow 10 seconds to close any resources
		_, cancel := context.WithTimeout(ctx, time.Second*utils.CLEAR_DB_RESOURCES_TIMEOUT)
		defer cancel()

		// Gracefully stop gRPC server
		grpcServer.GracefulStop()

		return err
	}
}
