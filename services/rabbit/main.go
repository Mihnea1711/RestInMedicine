package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/mihnea1711/POS_Project/services/rabbit/application"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/config"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/utils"
)

func main() {
	// setup logger
	log.Println("[RABBIT] Setting log stream to stdout...")
	log.SetOutput(os.Stdout) // Set log output to the stdout

	log.Println("[RABBIT] Application starting...")

	// load config file
	conf, err := config.LoadConfig(utils.CONFIG_PATH)
	if err != nil {
		log.Fatalf("[RABBIT] Error loading the config file: %s", err)
	} else {
		log.Println("[RABBIT] Successfully loaded the config file.")
	}

	log.Println(conf)

	// load .env vars into the app
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[RABBIT] Failed to load environment variables. Exitting...")
	}

	// load .env vars into the config
	config.ReplacePlaceholdersInStruct(conf)

	log.Println(conf)

	// catch interrupt signal
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	log.Println("[RABBIT] OS interrupt signals captured. Application will gracefully shut down on interruption...")

	// init the app
	app, err := application.New(conf, ctx)
	if err != nil {
		log.Fatalf("[RABBIT] Error creating and initializing the application: %s", err)
	} else {
		log.Println("[RABBIT] Application successfully initialized.")
	}

	// start the app with the created context
	err = app.Start(ctx)
	if err != nil {
		log.Fatalf("[RABBIT] Error starting the app: %v", err)
	} else {
		log.Println("[RABBIT] Application started successfully!")
	}
}
