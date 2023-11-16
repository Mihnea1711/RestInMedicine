package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/mihnea1711/POS_Project/services/consultatii/application"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/config"
	"github.com/mihnea1711/POS_Project/services/consultatii/pkg/utils"
)

func main() {
	// setup logger
	log.SetOutput(os.Stdout) // Set log output to the stdout
	log.Println("[CONSULTATION] Application starting...")

	// load config file
	conf, err := config.LoadConfig(utils.CONFIG_PATH)
	if err != nil {
		log.Fatalf("[CONSULTATION] Error loading the config file: %s", err)
	} else {
		log.Println("[CONSULTATION] Successfully loaded the config file.")
	}

	// load .env vars into the config
	utils.ReplacePlaceholdersInStruct(conf)

	// catch interrupt signal
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	log.Println("[CONSULTATION] OS interrupt signals captured. Application will gracefully shut down on interruption...")

	// init the app
	app, err := application.New(conf, ctx)
	if err != nil {
		log.Fatalf("[CONSULTATION] Error creating and initializing the application: %s", err)
	} else {
		log.Println("[CONSULTATION] Application successfully initialized.")
	}

	// start the app with the created context
	err = app.Start(ctx)
	if err != nil {
		log.Fatalf("[CONSULTATION] Error starting the app: %v", err)
	} else {
		log.Println("[CONSULTATION] Application started successfully!")
	}
}
