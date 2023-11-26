package config

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	// Server   ServerConfig   `yaml:"server"`
	RabbitMq RabbitMqConfig `yaml:"rabbitmq"`
}

// type ServerConfig struct {
// 	Port int `yaml:"port"`
// }

// RabbitMqConfig represents the RabbitMQ configuration
type RabbitMqConfig struct {
	Schema   string        `yaml:"schema"`
	Host     string        `yaml:"host"`
	Port     int           `yaml:"port"`
	Username string        `yaml:"username"`
	Password string        `yaml:"password"`
	Queues   []QueueConfig `yaml:"queues"`
}

// QueueConfig represents the configuration for a RabbitMQ queue
type QueueConfig struct {
	Name         string `yaml:"name"`
	RoutingKey   string `yaml:"routingKey"`
	Exchange     string `yaml:"exchange"`
	ExchangeType string `yaml:"exchange_type"`
	Durable      bool   `yaml:"durable"`
	Consumer     bool   `yaml:"consumer"`
}

// LoadConfig loads the configuration from the given file path
func LoadConfig(filePath string) (*AppConfig, error) {
	log.Println("[RABBIT] Loading configuration...")
	var conf AppConfig

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("[RABBIT] Error reading config file %s: %v\n", filePath, err)
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Printf("[RABBIT] Error unmarshaling config file %s: %v\n", filePath, err)
		return nil, fmt.Errorf("error unmarshaling config file: %w", err)
	}

	log.Printf("[RABBIT] Configuration loaded successfully from %s\n", filePath)
	return &conf, nil
}

func ReplaceWithEnvVars(input string) string {
	if strings.HasPrefix(input, "${") && strings.HasSuffix(input, "}") {
		envVar := strings.TrimSuffix(strings.TrimPrefix(input, "${"), "}")
		return os.Getenv(envVar)
	}
	return input
}

func ReplacePlaceholdersInStruct(s interface{}) {
	val := reflect.ValueOf(s)

	// Check if pointer and get the underlying value
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := field.Type()

		switch fieldType.Kind() {
		case reflect.String:
			if field.CanSet() {
				updatedValue := ReplaceWithEnvVars(field.String())
				field.SetString(updatedValue)
			}
		case reflect.Struct, reflect.Ptr:
			ReplacePlaceholdersInStruct(field.Addr().Interface())
		}
	}
}
