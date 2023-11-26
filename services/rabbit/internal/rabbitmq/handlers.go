package rabbitmq

import (
	"log"

	"github.com/mihnea1711/POS_Project/services/rabbit/internal/delete_service"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/utils"
)

// SetHandler sets the handler function for a specific queue
func (rmq *RabbitMQ) SetHandler(queueName string, handler func([]byte) error) {
	rmq.mu.Lock()
	defer rmq.mu.Unlock()
	rmq.handlers[queueName] = handler
}

// SetupHandlers sets up all the handlers for the queues. Need manual declaration
func (rmq *RabbitMQ) SetupHandlers() {
	log.Println("[RABBIT] Setting up message handlers for queues...")

	// Setting up a handler for the DELETE_QUEUE
	rmq.SetHandler(utils.DELETE_QUEUE, delete_service.DeleteUserMessageHandler)
	log.Printf("[RABBIT] Handler set up for queue '%s'", utils.DELETE_QUEUE)

	// Add more handlers for other queues as needed

	log.Println("[RABBIT] Message handlers set up successfully.")
}
