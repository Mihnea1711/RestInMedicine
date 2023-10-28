package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Server ServerConfig `yaml:"server"`
}

type ServerConfig struct {
	Port string `yaml:"port"`
}

// LoadConfig loads the configuration from the given file path
func LoadConfig(filePath string) *AppConfig {
	var conf AppConfig

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Fatalf("Error unmarshaling config file: %v", err)
	}

	return &conf
}
