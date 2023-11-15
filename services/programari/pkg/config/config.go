package config

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
	"time"

	"gopkg.in/yaml.v2"
)

type AppConfig struct {
	Server ServerConfig `yaml:"server"`
	MySQL  MySQLConfig  `yaml:"mysql_db"`
	Redis  RedisConfig  `yaml:"redis"`
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
	log.Println("[APPOINTMENT] Loading configuration...")
	var conf AppConfig

	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("[APPOINTMENT] Error reading config file %s: %v\n", filePath, err)
		return nil, fmt.Errorf("error reading config file: %w", err)
	}

	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		log.Printf("[APPOINTMENT] Error unmarshaling config file %s: %v\n", filePath, err)
		return nil, fmt.Errorf("error unmarshaling config file: %w", err)
	}

	log.Printf("[APPOINTMENT] Configuration loaded successfully from %s\n", filePath)
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
