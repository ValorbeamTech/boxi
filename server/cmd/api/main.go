package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/francis/projectx-api/internal/config"
	"github.com/francis/projectx-api/internal/database"
	"github.com/francis/projectx-api/internal/handler"
	"github.com/francis/projectx-api/internal/middleware"
	"github.com/francis/projectx-api/internal/repository/postgres"
	"github.com/francis/projectx-api/internal/service"
	"github.com/francis/projectx-api/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
    // Load configuration
    cfg := config.Load()

    // Initialize logger
    log := logger.New(cfg.LogLevel)

    // Initialize database
    db, err := database.NewPostgresDB(cfg.DatabaseURL)
    if err != nil {
        log.Fatal("Failed to connect to database", err)
    }
    defer db.Close()

    // Initialize repositories
    userRepo := postgres.NewUserRepository(db)

    // Initialize services
    authService := service.NewAuthService(userRepo, cfg.JWTSecret)
    userService := service.NewUserService(userRepo)

    // Initialize handlers
    authHandler := handler.NewAuthHandler(authService, log)
    userHandler := handler.NewUserHandler(userService, log)
    healthHandler := handler.NewHealthHandler(db, log)

    // Setup router
    router := setupRouter(cfg, authHandler, userHandler, healthHandler)

    // Setup server
    srv := &http.Server{
        Addr:         fmt.Sprintf(":%s", cfg.Port),
        Handler:      router,
        ReadTimeout:  15 * time.Second,
        WriteTimeout: 15 * time.Second,
        IdleTimeout:  60 * time.Second,
    }

    // Start server
    go func() {
        log.Info(fmt.Sprintf("Server starting on port %s", cfg.Port))
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatal("Failed to start server", err)
        }
    }()

    // Wait for interrupt signal to gracefully shutdown
    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    log.Info("Shutting down server...")

    // Graceful shutdown
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        log.Fatal("Server forced to shutdown", err)
    }

    log.Info("Server exited")
}

func setupRouter(cfg *config.Config, authHandler *handler.AuthHandler, userHandler *handler.UserHandler, healthHandler *handler.HealthHandler) *gin.Engine {
    if cfg.Environment == "production" {
        gin.SetMode(gin.ReleaseMode)
    }

    r := gin.New()

    // Global middleware
    r.Use(middleware.Logger())
    r.Use(middleware.CORS())
    // r.Use(middleware.RateLimit())
    r.Use(gin.Recovery())

    // Health check
    r.GET("/health", healthHandler.HealthCheck)

    // API routes
    api := r.Group("/api/v1")
    {
        // Auth routes
        auth := api.Group("/auth")
        {
            auth.POST("/register", authHandler.Register)
            auth.POST("/login", authHandler.Login)
            auth.POST("/refresh", authHandler.RefreshToken)
        }

        // Protected routes
        protected := api.Group("/")
        protected.Use(middleware.AuthMiddleware(cfg.JWTSecret))
        {
            // User routes
            users := protected.Group("/users")
            {
                users.GET("/profile", userHandler.GetProfile)
                users.PUT("/profile", userHandler.UpdateProfile)
                users.GET("", userHandler.GetUsers)
            }
        }
    }

    return r
}