package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Server ServerConfig `yaml:"server"`
	MySQL  MySQLConfig  `yaml:"mysql_db"`
}

type ServerConfig struct {
	Port int `yaml:"port"`
}

type MySQLConfig struct {
	Host            string        `yaml:"host"`
	Port            int           `yaml:"port"`
	User            string        `yaml:"user"`
	Password        string        `yaml:"password"`
	DbName          string        `yaml:"dbname"`
	Charset         string        `yaml:"charset"`
	ParseTime       bool          `yaml:"parseTime"`
	MaxOpenConns    int           `yaml:"maxOpenConns"`
	MaxIdleConns    int           `yaml:"maxIdleConns"`
	ConnMaxLifetime time.Duration `yaml:"connMaxLifetime"`
}

type RedisConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"` // usually 0 unless you're using multiple databases
}

// LoadConfig loads the configuration from the given file path
func LoadConfig(filePath string) (*AppConfig, error) {
	log.Println("Loading configuration...")
	var conf AppConfig

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading config file %s: %v\n", filePath, err)
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Printf("Error unmarshaling config file %s: %v\n", filePath, err)
		return nil, fmt.Errorf("error unmarshaling config file: %w", err)
	}

	log.Printf("Configuration loaded successfully from %s\n", filePath)
	return &conf, nil
}
