package main

import (
	"log"

	"golang/internal/config"
	"golang/internal/database"
	cliserver "golang/internal/server/cli"
	"golang/internal/service"
	"golang/internal/store"
	"golang/internal/utils/messaging"
)

func main() {
	// Load configuration
	cfg, err := config.Load("./config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Create database instance
	db, err := database.New(cfg)
	if err != nil {
		log.Fatalf("Failed to create database: %v", err)
	}

	err = db.Connect()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	storage, err := store.New(cfg, db.GetDB())
	if err != nil {
		log.Fatalf("Failed to create the store: %v", err)
	}

	// Integration Layer
	emailClient := messaging.NewEmailClient(cfg.Email.Token)
	emailClient.Connect()

	// Service Layer (SAME as HTTP server!)
	svc := service.NewService(storage, emailClient)

	// Presentation Layer (CLI)
	cli := cliserver.NewCLI(svc)

	// Start interactive CLI
	log.Printf("Starting CLI with store type: %s", cfg.Store.Type)
	cli.Start()
}
