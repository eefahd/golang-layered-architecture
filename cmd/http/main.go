package main

import (
	"log"
	"net/http"

	"golang/internal/config"
	"golang/internal/database"
	httpserver "golang/internal/server/http"
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

	// Create Store Layer
	storage, err := store.New(cfg, db.GetDB())
	if err != nil {
		log.Fatalf("Failed to create the store: %v", err)
	}

	// Integration Layer
	emailClient := messaging.NewEmailClient(cfg.Email.Token)
	emailClient.Connect()

	// Service Layer
	svc := service.NewService(storage, emailClient)

	// Presentation Layer (HTTP)
	server := httpserver.NewServer(svc)

	log.Printf("HTTP Server listening on :%s", cfg.Server.Port)
	log.Printf("  Store type: %s", cfg.Store.Type)
	log.Printf("  GET    /contacts")
	log.Printf("  GET    /contacts/{id}")
	log.Printf("  POST   /contacts")
	log.Printf("  PUT    /contacts/{id}")
	log.Printf("  DELETE /contacts/{id}")

	// TODO: handle graceful shutdown
	http.ListenAndServe(":"+cfg.Server.Port, server)
}
