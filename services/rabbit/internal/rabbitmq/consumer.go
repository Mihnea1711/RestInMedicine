package rabbitmq

import (
	"fmt"
	"log"
)

// Consumer Logic
func (rmq *RabbitMQ) Consume(queue Queue) error {
	msgs, err := rmq.channel.Consume(
		queue.Name, // queue name
		"",         // consumer id
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		log.Printf("[RABBIT] Failed to consume messages from queue '%s': %v", queue.Name, err)
		return err
	}

	log.Printf("[RABBIT] Started consuming messages from queue '%s'", queue.Name)

	handler, exists := rmq.handlers[queue.Name]

	if !exists {
		log.Printf("[RABBIT] No handler found for queue '%s'", queue.Name)
		return fmt.Errorf("no handler found for queue '%s'", queue.Name)
	}

	log.Printf("[RABBIT] Handler found for queue '%s'", queue.Name)

	go func() {

		for msg := range msgs {
			log.Printf("[RABBIT] Received message from queue '%s': %s", queue.Name, msg.Body)

			err := handler(msg.Body)
			if err != nil {
				log.Printf("[RABBIT] Error handling message from queue '%s': %v", queue.Name, err)
			}

			log.Printf("[RABBIT] Acknowledging message from queue '%s'", queue.Name)
			msg.Ack(false)
		}
	}()

	return nil
}
