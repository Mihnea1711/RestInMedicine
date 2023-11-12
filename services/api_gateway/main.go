package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/mihnea1711/POS_Project/services/gateway/application"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/config"
	"github.com/mihnea1711/POS_Project/services/gateway/pkg/utils"
)

func main() {
	// setup logger
	log.Println("[GATEWAY] Setting log stream to stdout...")
	log.SetOutput(os.Stdout) // Set log output to the stdout

	log.Println("[GATEWAY] Application starting...")

	// load config file
	conf, err := config.LoadConfig(utils.CONFIG_PATH)
	if err != nil {
		log.Fatalf("[GATEWAY] Error loading the config file: %s", err)
	} else {
		log.Println("[GATEWAY] Successfully loaded the config file.")
	}

	// load .env vars into the config
	config.ReplacePlaceholdersInStruct(conf)

	// catch interrupt signal
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	log.Println("[CONSULTATIE] OS interrupt signals captured. Application will gracefully shut down on interruption...")

	// init the app
	app, err := application.New(conf, ctx)
	if err != nil {
		log.Fatalf("[GATEWAY] Error creating and initializing the application: %s", err)
	} else {
		log.Println("[GATEWAY] Application successfully initialized.")
	}

	// start the app with the created context
	err = app.Start(ctx)
	if err != nil {
		log.Fatalf("[GATEWAY] Error starting the app: %v", err)
	} else {
		log.Println("[GATEWAY] Application started successfully!")
	}
}
