package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/mihnea1711/POS_Project/services/doctori/application"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/config"
)

func main() {
	config := config.LoadConfig("./configs/config.yaml")
	// router := routes.SetupRoutes()

	// // Starting the server
	// log.Fatal(http.ListenAndServe(conf.Server.Port, router))

	app := application.New(config)

	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	err := app.Start(ctx)
	if err != nil {
		fmt.Println("failed to start app: ", err)
	}
}
