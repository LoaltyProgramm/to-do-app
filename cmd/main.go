package main

import (
	"go1f/pkg/server"
	"go1f/pkg/db"
	"log"

	"github.com/joho/godotenv"

)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Fail read file .env: %v", err)
	}

	if err := db.InitDB(); err != nil {
		log.Fatalf("Fail init database: %v", err)
	}

	if err := server.StartServer(); err != nil {
		log.Fatalf("Server startup error: %v", err)
	}
}
