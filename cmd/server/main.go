package main

import (
	"errors"
	"log"
	"net/http"
	"time"

	"go-typesense-app/internal/config"
	"go-typesense-app/internal/database"
	"go-typesense-app/internal/handlers"
	"go-typesense-app/internal/routes"
	"go-typesense-app/internal/search"
	"go-typesense-app/internal/users"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: .env file not found: %v", err)
	}

	cfg := config.Load()

	db := database.NewDatabase(cfg)

	// init typesense client
	typesenseClient := search.NewTypesenseClient(cfg)

	// repositories
	userRepositoryImpl := database.NewUserRepository(db)

	// modules
	mainSnapshot := search.NewSnapshot(typesenseClient)

	// user modules
	searchClient := search.NewUserSearchService(typesenseClient)
	userModule := users.NewUserModule(userRepositoryImpl, searchClient)

	if err := userModule.InitializeSearchCollection(); err != nil {
		log.Printf("Warning: Failed to create Typesense collection: %v", err)
	}

	userHandler := handlers.NewUserHandler(userModule)
	snapshotHandler := handlers.NewSnapshotHandler(mainSnapshot)

	router := routes.SetupRoutes(userHandler, snapshotHandler)

	server := &http.Server{
		Addr:         cfg.GetServerAddress(),
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatal("Server failed to start:", err)
	}
}
