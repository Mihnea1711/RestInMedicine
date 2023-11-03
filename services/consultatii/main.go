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
	/*
		// // Setup logging
		// logFile, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		// if err != nil {
		// 	log.Fatalf("Error opening log file: %v", err)
		// 	return
		// }
		// defer func() {
		// 	if err := logFile.Close(); err != nil {
		// 		log.Printf("Error closing log file: %v", err)
		// 	}
		// }()
	*/

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

	// load .env vars (ONLY WHEN TESTING LOCALLY)
	// err = godotenv.Load()
	// if err != nil {
	// 	log.Fatal("[CONSULTATIE] Error loading .env file")
	// }

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
