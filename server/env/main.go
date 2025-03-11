package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB    PostgresConfig
	Web   WebConfig
}

type WebConfig struct {
	Host string
	Port string
}

type PostgresConfig struct {
	Username string
	Password string
	URL      string
	Port     string
	Host     string
	MFG		 string
}


func Env() (*Config) {
	err := godotenv.Load("../../../.env")

	if err != nil {
		log.Fatal("config failed: couldn't load .env")
	}

	cfg := &Config{
		Web: WebConfig {
			Port: os.Getenv("WEB_PORT"),
			Host: os.Getenv("WEB_HOST"),
		},
		DB: PostgresConfig {
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PWD"),
			URL:      os.Getenv("POSTGRES_URL"),
			Port:     os.Getenv("POSTGRES_PORT"),
			Host: 	  os.Getenv("POSTGRES_HOST"),
			MFG: 	  os.Getenv("POSTGRES_DB"),
		},
	}

	return cfg
}