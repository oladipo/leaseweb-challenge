package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/oladipo/leaseweb-challenge/internal/config"
	"github.com/oladipo/leaseweb-challenge/internal/handlers"
	"github.com/oladipo/leaseweb-challenge/internal/repository"
)

func main() {

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	// Connect to PostgreSQL
	dbURL := "postgres://" + cfg.User + ":" + cfg.Password + "@" + cfg.Host + ":" + cfg.Port + "/" + cfg.DBName
	db, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer db.Close()

	api := gin.Default()
	api.Use(gin.Logger())
	api.Use(gin.Recovery())

	// Initialize repository and handlers
	serverRepo := repository.NewServerRepository(db)
	serverHandler := handlers.NewServerHandler(serverRepo)

	RegisterRoutes(api, serverHandler)

	api.Run(":8080") // Start the server on port 8080
}

// RegisterRoutes registers the server-related routes
func RegisterRoutes(r *gin.Engine, serverHandler *handlers.ServerHandler) {
	api := r.Group("/api/v1")
	{
		api.GET("/servers", serverHandler.GetServers)
		api.POST("/servers/filter", serverHandler.FilterServers)
	}
}
