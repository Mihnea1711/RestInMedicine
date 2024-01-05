package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mihnea1711/POS_Project/services/rabbit/internal/middleware/validation"
	"github.com/mihnea1711/POS_Project/services/rabbit/internal/models"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/utils"
	"github.com/streadway/amqp"
)

func (rc *RabbitController) Publish(w http.ResponseWriter, r *http.Request) {
	log.Println("[RABBIT] Publishing message to queue...")

	// Validate and extract the userID
	userID, err := validation.ValidateUserID(r)
	if err != nil {
		log.Printf("[RABBIT] Validation error: %v\n", err)
		utils.RespondWithJSON(w, http.StatusUnprocessableEntity, models.ResponseData{
			Error: err.Error(),
		})
		return
	}

	log.Printf("[RABBIT] Validated UserID: %d\n", userID)

	// Publish the userID to RabbitMQ
	err = rc.publishToQueue(userID)
	if err != nil {
		log.Printf("[RABBIT] Error publishing message to queue: %v\n", err)
		utils.RespondWithJSON(w, http.StatusInternalServerError, models.ResponseData{
			Error: err.Error(),
		})
		return
	}

	log.Printf("[RABBIT] User ID %d published to RabbitMQ\n", userID)
	utils.RespondWithJSON(w, http.StatusOK, models.ResponseData{
		Message: "Successfully published the message to the delete queue..",
	})
}

func (rc *RabbitController) publishToQueue(userID int) error {
	// RabbitMQ server connection URL
	amqpURL := "amqp://guest:guest@rabbitmq:5672/"

	log.Printf("[RABBIT] Establishing a connection to RabbitMQ server...\n")

	// Establish a connection to RabbitMQ server
	conn, err := amqp.Dial(amqpURL)
	if err != nil {
		log.Printf("[RABBIT] Connection error: %v\n", err)
		return err
	}
	defer conn.Close()

	log.Println("[RABBIT] Connection to RabbitMQ established successfully")

	log.Println("[RABBIT] Creating a channel...")

	// Create a channel
	ch, err := conn.Channel()
	if err != nil {
		log.Printf("[RABBIT] Channel creation error: %v\n", err)
		return err
	}
	defer ch.Close()

	log.Println("[RABBIT] Channel created successfully")

	log.Printf("[RABBIT] Publishing UserID: %d to the queue...\n", userID)

	// Publish the userID to the queue
	userIDMessage := models.DeleteMessageData{IDUser: userID}
	log.Printf("[RABBIT] userIDMessage content: %+v\n", userIDMessage)
	userIDMessageBytes, err := json.Marshal(userIDMessage)
	if err != nil {
		log.Printf("[RABBIT] JSON encoding error: %v\n", err)
		return err
	}
	// Publish the userID to the queue
	err = ch.Publish(
		"",                 // exchange
		utils.DELETE_QUEUE, // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        userIDMessageBytes,
		})
	if err != nil {
		log.Printf("[RABBIT] Publishing error: %v\n", err)
		return err
	}

	log.Printf("[RABBIT] [Publish] Sent UserID: %d to RabbitMQ\n", userID)
	return nil
}
