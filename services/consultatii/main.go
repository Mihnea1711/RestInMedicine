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
	log.Println("[CONSULTATIE] Application starting...")

	// load config file
	conf, err := config.LoadConfig(utils.CONFIG_PATH)
	if err != nil {
		log.Fatalf("[CONSULTATIE] Error loading the config file: %s", err)
	} else {
		log.Println("[CONSULTATIE] Successfully loaded the config file.")
	}

	// load .env vars into the config
	utils.ReplacePlaceholdersInStruct(conf)

	// catch interrupt signal
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	log.Println("[CONSULTATIE] OS interrupt signals captured. Application will gracefully shut down on interruption...")

	// init the app
	app, err := application.New(conf, ctx)
	if err != nil {
		log.Fatalf("[CONSULTATIE] Error creating and initializing the application: %s", err)
	} else {
		log.Println("[CONSULTATIE] Application successfully initialized.")
	}

	// start the app with the created context
	err = app.Start(ctx)
	if err != nil {
		log.Fatalf("[CONSULTATIE] Error starting the app: %v", err)
	} else {
		log.Println("[CONSULTATIE] Application started successfully!")
	}
}
