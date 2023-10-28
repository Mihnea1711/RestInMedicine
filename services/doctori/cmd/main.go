package main

import (
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/doctori/internal/routes"
	"github.com/mihnea1711/POS_Project/services/doctori/pkg/config"
)

func main() {
	conf := config.LoadConfig("./configs/config.yaml")
	router := routes.SetupRoutes()

	// Starting the server
	log.Fatal(http.ListenAndServe(conf.Server.Port, router))
}
