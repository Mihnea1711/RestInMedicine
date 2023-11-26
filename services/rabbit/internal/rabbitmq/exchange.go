package rabbitmq

import "log"

// DeclareExchange declares a RabbitMQ exchange
func (rmq *RabbitMQ) DeclareExchange(exchangeName, exchangeType string) error {
	log.Printf("[RABBIT] Declaring exchange '%s'...", exchangeName)

	err := rmq.channel.ExchangeDeclare(
		exchangeName, // exchange name
		exchangeType, // exchange type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Printf("[RABBIT] Failed to declare exchange '%s': %v", exchangeName, err)
		return err
	}

	log.Printf("[RABBIT] Exchange '%s' declared successfully", exchangeName)
	return nil
}
