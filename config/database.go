package config

import (
	"fmt"
	"log"
	"os"

	"github.com/srivathsav-max/backend/prisma/db"
)

var DB *db.PrismaClient

func InitDatabase() {
	if os.Getenv("DATABASE_URL") == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}
	if os.Getenv("DIRECT_URL") == "" {
		log.Fatal("DIRECT_URL environment variable is not set")
	}

	DB = db.NewClient()
	if err := DB.Prisma.Connect(); err != nil {
		log.Fatal(fmt.Sprintf("Could not connect to database: %v", err))
	}
	
	log.Println("✅ Successfully connected to database")
}

func DisconnectDatabase() {
	if DB != nil {
		if err := DB.Prisma.Disconnect(); err != nil {
			log.Printf("Warning: Error disconnecting from database: %v", err)
			return
		}
		log.Println("✅ Successfully disconnected from database")
	}
}
