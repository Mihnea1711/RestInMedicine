package service

import (
	"fmt"
)

// ExampleMessageHandler is an example implementation of MessageHandler.
func DeleteUserMessageHandler(message []byte) error {
	// Your message handling logic goes here.
	// This is just a placeholder.
	fmt.Printf("[RABBIT] Handling delete message: %s\n", string(message))
	return nil
}
