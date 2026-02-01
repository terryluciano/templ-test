package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	// Server
	SERVER_PORT string `env:"PORT" envDefault:"3000"`

	// Database
	DATABASE_HOST     string `env:"DATABASE_HOST,required"`
	DATABASE_USER     string `env:"DATABASE_USER,required"`
	DATABASE_PASSWORD string `env:"DATABASE_PASSWORD,required"`
	DATABASE_NAME     string `env:"DATABASE_NAME,required"`
	DATABASE_PORT     string `env:"DATABASE_PORT,required"`
	DATABASE_SSL_MODE string `env:"DATABASE_SSL_MODE" envDefault:"disable"`
	DATABASE_TIMEZONE string `env:"DATABASE_TIMEZONE"`
}

var Config EnvConfig

func LoadConfig() {
	dir, _ := os.Getwd()

	envPath := filepath.Join(dir, "env", ".env")

	log.Println("Loading Env...")

	envErr := godotenv.Load(envPath)

	if envErr != nil {
		log.Printf("warning: no .env loaded from %s: %v", envPath, envErr)
	}

	configErr := env.Parse(&Config)

	if configErr != nil {
		log.Fatalf("Error loading config: %s", configErr)
	}

	log.Println("Config Loaded")
}
