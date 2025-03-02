package config

import (
	"fmt"
	"log"

	"github.com/srivathsav-max/backend/prisma/db"
)

var DB *db.PrismaClient

func InitDatabase() error {
	// Ensure configuration is loaded
	if err := LoadConfig(); err != nil {
		return fmt.Errorf("failed to load configuration: %v", err)
	}

	DB = db.NewClient()
	if err := DB.Prisma.Connect(); err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	log.Println("✅ Successfully connected to database")
	return nil
}

func DisconnectDatabase() {
	if DB != nil {
		if err := DB.Prisma.Disconnect(); err != nil {
			log.Printf("⚠️ Warning: Error disconnecting from database: %v", err)
			return
		}
		log.Println("✅ Successfully disconnected from database")
	}
}
