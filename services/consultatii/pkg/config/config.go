package config

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Server ServerConfig  `yaml:"server"`
	Mongo  MongoDBConfig `yaml:"mongodb"`
	Redis  RedisConfig   `yaml:"redis"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"` // usually 0 unless you're using multiple databases
}

type MongoDBConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

// LoadConfig loads the configuration from the given file path
func LoadConfig(filePath string) (*AppConfig, error) {
	log.Println("[CONSULTATIE] Loading configuration...")
	var conf AppConfig

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("[CONSULTATIE] Error reading config file %s: %v\n", filePath, err)
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Printf("[CONSULTATIE] Error unmarshaling config file %s: %v\n", filePath, err)
		return nil, fmt.Errorf("error unmarshaling config file: %w", err)
	}

	log.Printf("[CONSULTATIE] Configuration loaded successfully from %s\n", filePath)
	return &conf, nil
}
