package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/mihnea1711/POS_Project/services/idm/application"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/config"
	"github.com/mihnea1711/POS_Project/services/idm/pkg/utils"
)

func main() {
	// setup logger
	log.Println("[IDM] Setting log stream to stdout...")
	log.SetOutput(os.Stdout) // Set log output to the stdout

	log.Println("[IDM] Application starting...")

	// load config file
	conf, err := config.LoadConfig(utils.CONFIG_PATH)
	if err != nil {
		log.Fatalf("[IDM] Error loading the config file: %s", err)
	} else {
		log.Println("[IDM] Successfully loaded the config file.")
	}

	// load .env vars into the app
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("[PATIENT] Failed to load environment variables. Exitting...")
	}

	// load .env vars into the config
	utils.ReplacePlaceholdersInStruct(conf)

	// catch interrupt signal
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	log.Println("[IDM] OS interrupt signals captured. Application will gracefully shut down on interruption...")

	// init the app
	app, err := application.New(conf, ctx)
	if err != nil {
		log.Fatalf("[IDM] Error creating and initializing the application: %s", err)
	} else {
		log.Println("[IDM] Application successfully initialized.")
	}

	// start the app with the created context
	err = app.Start(ctx)
	if err != nil {
		log.Fatalf("[IDM] Error starting the app: %v", err)
	} else {
		log.Println("[IDM] Application started successfully!")
	}
}
