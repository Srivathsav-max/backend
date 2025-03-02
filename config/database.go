package config

import (
	"log"

	"lumen/backend/prisma/db"
)

var DB *db.PrismaClient

func InitDatabase(){

	DB = db.NewClient()

	err := DB.Prisma.Connect()
	if err != nil{
		log.Fatal("Could Not Connect to Database", err)
	}
	log.Println("Connected to Database")
}

func DisconnectDatabase(){
	err := DB.Prisma.Disconnect()
	if err != nil{
		log.Fatal("Could Not Disconnect to Database", err)
	}
	log.Println("Disconnected to Database")
}