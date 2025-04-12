package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ENVIROMENT int

const (
	development ENVIROMENT = iota
	production
	test
)

var Enviroment = map[ENVIROMENT]string{
	development: "development",
	production:  "production",
	test:        "test",
}

type Config struct {
	ENV      ENVIROMENT
	SMTP     SMTPConfig
	DB       PostgresConfig
	DB_CACHE RedisConfig
	API      APIConfig
}

type RedisConfig struct {
	Username  string
	Password  string
	URL       string
	Port      string
	Host      string
	DefaultDB string
}

type SMTPConfig struct {
	Host      string
	Port      string
	HostEmail string
	Password  string
}

type APIConfig struct {
	Host         string
	Port         string
	PasswordSalt string
	JWTSalt      string
}

type PostgresConfig struct {
	Username  string
	Password  string
	URL       string
	Port      string
	Host      string
	DefaultDB string
}

var (
	enviroment ENVIROMENT
)

func Env() *Config {
	err := godotenv.Load("../../../.env")

	if err != nil {
		log.Fatal("config failed: couldn't load .env")
	}

	switch os.Getenv("ENVIROMENT") {
	case "production":
		enviroment = production
	case "test":
		enviroment = test
	default:
		enviroment = development
	}

	cfg := &Config{
		ENV: enviroment,
		SMTP: SMTPConfig{
			Port:      os.Getenv("SMTP_PORT"),
			Host:      os.Getenv("SMTP_HOST"),
			HostEmail: os.Getenv("SMTP_EMAIL"),
			Password:  os.Getenv("SMTP_PASSWORD"),
		},
		API: APIConfig{
			Port:         os.Getenv("API_PORT"),
			Host:         os.Getenv("API_HOST"),
			PasswordSalt: os.Getenv("API_PASSWORD_SALT"),
			JWTSalt:      os.Getenv("API_JWT_SALT"),
		},
		DB_CACHE: RedisConfig{
			Username:  os.Getenv("REDIS_USER"),
			Password:  os.Getenv("REDIS_PASSWORD"),
			URL:       os.Getenv("REDIS_URL"),
			Port:      os.Getenv("REDIS_PORT"),
			Host:      os.Getenv("REDIS_HOST"),
			DefaultDB: os.Getenv("REDIS_DB"),
		},
		DB: PostgresConfig{
			Username:  os.Getenv("POSTGRES_USER"),
			Password:  os.Getenv("POSTGRES_PASSWORD"),
			URL:       os.Getenv("POSTGRES_URL"),
			Port:      os.Getenv("POSTGRES_PORT"),
			Host:      os.Getenv("POSTGRES_HOST"),
			DefaultDB: os.Getenv("POSTGRES_DB"),
		},
	}

	return cfg
}
