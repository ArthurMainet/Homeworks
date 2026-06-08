package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	EmailConf *EmailConfig
	DB        *DbConfig
	AuthToken *AuthConfig
}

type EmailConfig struct {
	Email    string
	Password string
	Address  string
}

type DbConfig struct {
	DSN string
}

type AuthConfig struct {
	AuthToken string
}

func LoadConfig() *Config {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading download special config: ", err)
	}
	return &Config{
		EmailConf: &EmailConfig{
			Email:    "testmail@gmail.com",
			Password: "123",
			Address:  "smtp.gmail.com:587",
		},
		DB: &DbConfig{
			DSN: os.Getenv("DSN"),
		},
		AuthToken: &AuthConfig{
			AuthToken: os.Getenv("TOKEN"),
		},
	}
}
