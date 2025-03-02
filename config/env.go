package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

// Config holds all environment configuration
type Config struct {
	// Server configuration
	Port        string `envconfig:"PORT"`
	Development bool   `envconfig:"DEVELOPMENT"`
	CorsOrigin  string `envconfig:"CORS_ORIGIN"`

	// Database configuration
	DatabaseURL string `envconfig:"DATABASE_URL" required:"true"`
	DirectURL   string `envconfig:"DIRECT_URL" required:"true"`
}

var AppConfig Config

// LoadConfig loads environment variables into the Config struct
func LoadConfig() error {
	if err := envconfig.Process("", &AppConfig); err != nil {
		return fmt.Errorf("error loading config: %v", err)
	}

	// Set CORS origin for production if not explicitly set
	if !AppConfig.Development && AppConfig.CorsOrigin == "http://localhost:3000" {
		AppConfig.CorsOrigin = "https://moxium.tech"
	}

	log.Printf("✅ Configuration loaded (Development: %v)", AppConfig.Development)
	return nil
}
