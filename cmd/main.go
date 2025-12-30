package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go-setup/internal/config"
	"go-setup/internal/database"
	http_delivery "go-setup/internal/delivery/http"
	"go-setup/internal/health"
	"go-setup/internal/repository/postgres"
	"go-setup/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	
	// Config
	cfg := config.Load()
	
	// Database + graceful close
	db, err := database.NewPostgres(&cfg.DB)
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	defer db.Close()
	
	// Dependency Injection
	userRepo := postgres.NewUserRepository(db)
	userUC := usecase.NewUserUsecase(userRepo)
	userHandler := http_delivery.NewUserHandler(userUC)
	healthHandler := health.NewHandler(db)
	
	// Gin router
	r := gin.Default()
	r.Use(http_delivery.CORSMiddleware())
	r.Use(http_delivery.ErrorHandlingMiddleware())
	
	// Health check (Production MUST)
	r.GET("/health", healthHandler.Check())
	
	// Business routes
	api := r.Group("/api/v1/users")
	{
		api.POST("", userHandler.CreateUser)
		api.GET("", userHandler.ListUsers)
		api.GET("/:id", userHandler.GetUser)
		api.PUT("/:id", userHandler.UpdateUser)
		api.DELETE("/:id", userHandler.DeleteUser)
	}
	
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}
	
	// Start server
	go func() {
		log.Printf("Server listening on :%s", cfg.Port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("Server failed:", err)
		}
	}()
	
	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
	
	log.Println("Server exiting")
}
