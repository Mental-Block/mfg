package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	SMTP SMTPConfig
	DB   PostgresConfig
	Web  WebConfig
}

// SMTP server
type SMTPConfig struct {
	Host      string
	Port      string
	HostEmail string
	Password  string
}

// Go REST server API
type WebConfig struct {
	Host string
	Port string
	Salt string
}

// postgress database config
type PostgresConfig struct {
	Username string
	Password string
	URL      string
	Port     string
	Host     string
	MFG      string
}

func Env() *Config {
	err := godotenv.Load("../../../.env")

	if err != nil {
		log.Fatal("config failed: couldn't load .env")
	}

	cfg := &Config{
		SMTP: SMTPConfig{
			Port:      os.Getenv("SMTP_PORT"),
			Host:      os.Getenv("SMTP_HOST"),
			HostEmail: os.Getenv("SMTP_EMAIL"),
			Password:  os.Getenv("SMTP_PASSWORD"),
		},
		Web: WebConfig{
			Port: os.Getenv("WEB_PORT"),
			Host: os.Getenv("WEB_HOST"),
			Salt: os.Getenv("WEB_SALT"),
		},
		DB: PostgresConfig{
			Username: os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PWD"),
			URL:      os.Getenv("POSTGRES_URL"),
			Port:     os.Getenv("POSTGRES_PORT"),
			Host:     os.Getenv("POSTGRES_HOST"),
			MFG:      os.Getenv("POSTGRES_DB"),
		},
	}

	return cfg
}
