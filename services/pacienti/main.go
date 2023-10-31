package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	"github.com/mihnea1711/POS_Project/services/pacienti/application"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/config"
	"github.com/mihnea1711/POS_Project/services/pacienti/pkg/utils"
)

func main() {
	/*
		// // Setup logging
		// logFile, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		// if err != nil {
		// 	log.Fatalf("[PACIENTI] Error opening log file: %v", err)
		// 	return
		// }
		// defer func() {
		// 	if err := logFile.Close(); err != nil {
		// 		log.Printf("[PACIENTI] Error closing log file: %v", err)
		// 	}
		// }()
	*/

	// setup logger
	log.Println("[PACIENTI] Setting log stream to stdout...")
	log.SetOutput(os.Stdout) // Set log output to the stdout

	log.Println("[PACIENTI] Application starting...")

	// load config file
	conf, err := config.LoadConfig(utils.CONFIG_PATH)
	if err != nil {
		log.Fatalf("[PACIENTI] Error loading the config file: %s", err)
	} else {
		log.Println("[PACIENTI] Successfully loaded the config file.")
	}

	/*
		// // load .env vars (ONLY WHEN TESTING LOCALLY)
		// err = godotenv.Load()
		// if err != nil {
		// 	log.Fatal("[PACIENTI] Error loading .env file")
		// }
	*/

	// load .env vars into the config
	config.ReplacePlaceholdersInStruct(conf)

	// init the app
	app, err := application.New(conf)
	if err != nil {
		log.Fatalf("[PACIENTI] Error creating and initializing the application: %s", err)
	} else {
		log.Println("[PACIENTI] Application successfully initialized.")
	}

	// catch interrupt signal
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()
	log.Println("[PACIENTI] OS interrupt signals captured. Application will gracefully shut down on interruption...")

	// start the app with the created context
	err = app.Start(ctx)
	if err != nil {
		log.Fatalf("[PACIENTI] Error starting the app: %v", err)
	} else {
		log.Println("[PACIENTI] Application started successfully!")
	}
}
