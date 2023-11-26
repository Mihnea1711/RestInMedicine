package controllers

import (
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/rabbit/internal/rabbitmq"
)

type RabbitController struct {
	RabbitMQ *rabbitmq.RabbitMQ
}

func (rc *RabbitController) HandleHealthCheck(w http.ResponseWriter, r *http.Request) {
	log.Println("[RABBIT_HEALTH_CHECK] HANDLED HEALTH CHECK")

	// Write OK response to the client
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK\n"))
}
