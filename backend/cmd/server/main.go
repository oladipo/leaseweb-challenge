package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/gzip"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"

	"github.com/oladipo/leaseweb-challenge/internal/config"
	"github.com/oladipo/leaseweb-challenge/internal/handlers"
	"github.com/oladipo/leaseweb-challenge/internal/models"
	"github.com/oladipo/leaseweb-challenge/internal/repository"

	"github.com/ulule/limiter/v3"
	ginmw "github.com/ulule/limiter/v3/drivers/middleware/gin"
	"github.com/ulule/limiter/v3/drivers/store/memory"
	"github.com/unrolled/secure"
	ginprometheus "github.com/zsais/go-gin-prometheus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	db, err := connectDB(cfg)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	api := gin.Default()
	// Set Gin to release mode for production
	// gin.SetMode(gin.ReleaseMode)

	// Set up middlewares
	api.Use(gin.Logger())
	api.Use(gin.Recovery())
	api.Use(cors.Default())                     // CORS
	api.Use(gzip.Gzip(gzip.DefaultCompression)) // GZIP compression
	api.Use(requestid.New())                    // Request ID
	api.Use(secureHeaders())                    // Secure headers
	api.Use(RateLimiter())                      // Rate Limiting
	// TODO: middlewares: JWT and  Rate Limiting

	// Prometheus metrics
	p := ginprometheus.NewPrometheus("gin_app")
	p.Use(api)

	// Initialize repository and handlers
	serverRepo := repository.NewServerRepository(db)
	serverHandler := handlers.NewServerHandler(serverRepo)

	registerRoutes(api, serverHandler)

	api.Run(":8080") // Start the server on port 8080
}

// RegisterRoutes registers the server-related routes
func registerRoutes(r *gin.Engine, serverHandler *handlers.ServerHandler) {
	api := r.Group("/api/v1")
	{
		api.GET("/servers", serverHandler.GetServers)
		api.POST("/servers/filter", serverHandler.FilterServers)
	}
}

func connectDB(cfg *config.DBConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(&models.Server{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	return db, nil
}

func secureHeaders() gin.HandlerFunc {
	secureMiddleware := secure.New(secure.Options{
		FrameDeny:            true,
		ContentTypeNosniff:   true,
		BrowserXssFilter:     true,
		SSLRedirect:          false,
		STSSeconds:           31536000,
		STSIncludeSubdomains: true,
	})
	return func(c *gin.Context) {
		if err := secureMiddleware.Process(c.Writer, c.Request); err != nil {
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
		c.Next()
	}
}

func RateLimiter() gin.HandlerFunc {
	rate := limiter.Rate{
		Period: 1 * time.Minute,
		Limit:  60,
	}
	store := memory.NewStore() // In-memory for now
	return ginmw.NewMiddleware(limiter.New(store, rate))
}
