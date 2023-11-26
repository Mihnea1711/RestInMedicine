// internal/rabbitmq/rabbitmq.go
package rabbitmq

import (
	"context"
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/config"
	"github.com/mihnea1711/POS_Project/services/rabbit/pkg/utils"
	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageHandler func([]byte) error
type RabbitMQ struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	mu       sync.Mutex
	handlers map[string]MessageHandler
}

// NewRabbitMQ creates a new RabbitMq Controller struct
func NewRabbitMQ(rabbitConfig config.RabbitMqConfig) (*RabbitMQ, error) {
	url := fmt.Sprintf(
		"%s://%s:%s@%s:%d",
		rabbitConfig.Schema,
		rabbitConfig.Username,
		rabbitConfig.Password,
		rabbitConfig.Host,
		rabbitConfig.Port,
	)

	conn, err := amqp.Dial(url)
	if err != nil {
		log.Printf("[RABBIT] Failed to connect to RabbitMQ: %v", err)
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		log.Printf("[RABBIT] Failed to open a channel: %v", err)
		return nil, err
	}

	rmq := &RabbitMQ{
		conn:     conn,
		channel:  channel,
		handlers: make(map[string]MessageHandler),
	}

	log.Printf("[RABBIT] RabbitMQ connection established: %s", url)

	return rmq, nil
}

// Close function closes any connection and queue that is open
func (rmq *RabbitMQ) Close(ctx context.Context) error {
	rmq.mu.Lock()
	defer rmq.mu.Unlock()

	// Create a context with a timeout for closing operations
	closeCtx, cancel := context.WithTimeout(ctx, time.Second*utils.RABBIT_CLOSE_TIMEOUT)
	defer cancel()

	// Close the channel in a goroutine
	closeChannel := make(chan error, 1)
	go func() {
		if err := rmq.channel.Close(); err != nil {
			log.Printf("[RABBIT] Failed to close channel: %v", err)
			closeChannel <- err
		} else {
			log.Println("[RABBIT] Channel closed successfully.")
			closeChannel <- nil
		}
	}()

	// Close the connection in a goroutine
	closeConnection := make(chan error, 1)
	go func() {
		if err := rmq.conn.Close(); err != nil {
			log.Printf("[RABBIT] Failed to close connection: %v", err)
			closeConnection <- err
		} else {
			log.Println("[RABBIT] Connection closed successfully.")
			closeConnection <- nil
		}
	}()

	// Wait for both operations to complete or the context to be done
	select {
	case channelErr := <-closeChannel:
		if channelErr != nil {
			return channelErr
		}
	case connectionErr := <-closeConnection:
		if connectionErr != nil {
			return connectionErr
		}
	case <-closeCtx.Done():
		log.Println("[RABBIT] Closing operations timed out.")
		return errors.New("closing operations timed out")
	}

	return nil
}
