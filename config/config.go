package config

import (
	

	"github.com/caarlos0/env"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	ServerPort string `env:"SERVER_PORT,required"`
	DBHOST     string `env:"DB_HOST,required"`
	DBNAME     string `env:"DB_NAME,required"`
	DBUSER     string `env:"DB_USER,required"`
	DBPASSWORD string `env:"DB_PASSWORD,required"`
	DBSSLMode  string `env:"DB_SSLMODE,required"`
}

func NewEnvConfig() *EnvConfig {
	
	err := godotenv.Load()
	if err != nil {
	 log.Fatalf("Error loading .env file: %e",err)
	}

	config := &EnvConfig{}

	if err := env.Parse(config); err != nil {
		log.Fatalf("Unable to load variables from the env: %e", err)
	}
	return config
}