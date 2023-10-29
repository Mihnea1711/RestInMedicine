package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/mihnea1711/POS_Project/services/doctori/application"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/config"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/utils"
)

func main() {
	// Setup logging
	logFile, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
		return
	}
	defer func() {
		if err := logFile.Close(); err != nil {
			log.Printf("Error closing log file: %v", err)
		}
	}()

	log.SetOutput(logFile) // Set log output to the file
	log.Println("Application starting...")

	config, err := config.LoadConfig(utils.CONFIG_PATH)
	if err != nil {
		log.Fatalf("Error loading the config file: %s", err)
	} else {
		log.Println("Successfully loaded the config file.")
	}

	app, err := application.New(config)
	if err != nil {
		log.Fatalf("Error creating and initializing the application: %s", err)
	} else {
		log.Println("Application successfully initialized.")
	}

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	log.Println("OS interrupt signals captured. Application will gracefully shut down on interruption...")

	err = app.Start(ctx)
	if err != nil {
		log.Fatalf("Error starting the app: %v", err)
	} else {
		log.Println("Application started successfully!")
	}
}
