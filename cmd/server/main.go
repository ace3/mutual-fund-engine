package main

import (
	"log"

	"github.com/ace3/mutual-fund-engine/internal/config"
	"github.com/ace3/mutual-fund-engine/internal/database"
	"github.com/ace3/mutual-fund-engine/internal/grpc"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}
	cfg := config.Load()
	db := database.Connect(cfg)
	dbSQL, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get database connection: %v", err)
	}
	defer dbSQL.Close()
	grpc.StartServer(cfg, db)
}
