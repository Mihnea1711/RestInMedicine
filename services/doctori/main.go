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
	// setup logger
	log.SetOutput(os.Stdout) // Set log output to the stdout
	log.Println("[DOCTOR] Application starting...")

	// load config file
	conf, err := config.LoadConfig(utils.CONFIG_PATH)
	if err != nil {
		log.Fatalf("[DOCTOR] Error loading the config file: %s", err)
	} else {
		log.Println("[DOCTOR] Successfully loaded the config file.")
	}

	// load .env vars into the config
	config.ReplacePlaceholdersInStruct(conf)

	// catch interrupt signal
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	log.Println("[DOCTOR] OS interrupt signals captured. Application will gracefully shut down on interruption...")

	// init the app
	app, err := application.New(conf, ctx)
	if err != nil {
		log.Fatalf("[DOCTOR] Error creating and initializing the application: %s", err)
	} else {
		log.Println("[DOCTOR] Application successfully initialized.")
	}

	// start the app with the created context
	err = app.Start(ctx)
	if err != nil {
		log.Fatalf("[DOCTOR] Error starting the app: %v", err)
	} else {
		log.Println("[DOCTOR] Application started successfully!")
	}
}
