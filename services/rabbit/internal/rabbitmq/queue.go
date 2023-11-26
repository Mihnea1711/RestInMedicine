package rabbitmq

import (
	"fmt"
	"log"

	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/config"
)

// Queue represents a RabbitMQ queue
type Queue struct {
	Name         string
	RoutingKey   string
	Exchange     string
	ExchangeType string
	Durable      bool
	IsConsumer   bool
}

// NewQueue creates a new RabbitMQ queue
func NewQueue(name, routingKey, exchange, exchangeType string, durable, consumerFlag bool) *Queue {
	return &Queue{
		Name:         name,
		RoutingKey:   routingKey,
		Exchange:     exchange,
		ExchangeType: exchangeType,
		Durable:      durable,
		IsConsumer:   consumerFlag,
	}
}

// DeclareQueue declares a RabbitMQ queue
func (rmq *RabbitMQ) DeclareAndBindQueue(queue *Queue) error {
	log.Printf("[RABBIT] Declaring and binding queue '%s'...", queue.Name)

	_, err := rmq.channel.QueueDeclare(
		queue.Name,    // queue name
		queue.Durable, // durable
		false,         // delete when unused
		false,         // exclusive
		false,         // no-wait
		nil,           // arguments
	)
	if err != nil {
		log.Printf("[RABBIT] Failed to declare queue '%s': %v", queue.Name, err)
		return err
	}

	err = rmq.channel.QueueBind(
		queue.Name,       // queue name
		queue.RoutingKey, // routing key
		queue.Exchange,   // exchange
		false,            // no-wait
		nil,              // arguments
	)
	if err != nil {
		log.Printf("[RABBIT] Failed to bind queue '%s' to exchange: %v", queue.Name, err)
		return err
	}

	log.Printf("[RABBIT] Queue '%s' declared and bound to exchange successfully", queue.Name)
	return nil
}

// SetupQueues sets up each queue: declare, bind, and consume/publish
func (rmq *RabbitMQ) SetupQueues(queues []*Queue) error {
	rmq.mu.Lock()
	defer rmq.mu.Unlock()

	for _, queue := range queues {
		err := rmq.DeclareExchange(queue.Exchange, queue.ExchangeType)
		if err != nil {
			log.Printf("[RABBIT] Failed to declare exchange '%s.%s': %v", queue.Exchange, queue.ExchangeType, err)
			return err
		}
		log.Printf("[RABBIT] Exchange '%s:%s' declared and bound successfully", queue.Exchange, queue.ExchangeType)

		err = rmq.DeclareAndBindQueue(queue)
		if err != nil {
			log.Printf("[RABBIT] Failed to declare and bind queue '%s': %v", queue.Name, err)
			return err
		}
		log.Printf("[RABBIT] Queue '%s' declared and bound successfully", queue.Name)

		if queue.IsConsumer {
			err := rmq.Consume(*queue)
			if err != nil {
				log.Printf("[RABBIT] Failed to start consumer for queue '%s': %v", queue.Name, err)
				return err
			}
			log.Printf("[RABBIT] Consumer started for queue '%s'", queue.Name)
		}
	}

	return nil
}

// TranslateConfigToQueue translates RabbitMqConfig to a slice of Queue
func TranslateConfigToQueue(config config.RabbitMqConfig) ([]*Queue, error) {
	var queues []*Queue
	for _, q := range config.Queues {
		queue := NewQueue(
			q.Name,
			q.RoutingKey,
			q.Exchange,
			q.ExchangeType,
			q.Durable,
			q.Consumer,
		)
		if queue == nil {
			return nil, fmt.Errorf("failed to create Queue for config: %+v", q)
		}
		queues = append(queues, queue)
	}
	return queues, nil
}
