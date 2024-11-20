package main

import (
	"github.com/TakuroBreath/personal-finance-manager/internal/api/handlers"
	"github.com/TakuroBreath/personal-finance-manager/internal/api/routes"
	"github.com/TakuroBreath/personal-finance-manager/internal/config"
	"github.com/TakuroBreath/personal-finance-manager/internal/repository"
	"github.com/TakuroBreath/personal-finance-manager/internal/service"
	"github.com/TakuroBreath/personal-finance-manager/internal/storage"
	"github.com/TakuroBreath/personal-finance-manager/pkg/logger"
	"github.com/TakuroBreath/personal-finance-manager/pkg/sl"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	l "log"
	"log/slog"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		l.Fatal("Error loading .env file")
	}

	ConfigPath := os.Getenv("CONFIG_PATH")
	secret := os.Getenv("JWT_SECRET")

	cfg := config.MustLoad(ConfigPath)

	log := logger.SetupLogger(cfg.Env)

	log.Info("Starting app", slog.String("env", cfg.Env))

	db, err := storage.NewStorage(cfg)
	if err != nil {
		log.Error("failed to connect to database", sl.Err(err))
		os.Exit(1)
	}

	repo := repository.NewRepository(db)

	jwtService := service.NewJWTService(secret)

	userService := service.NewUserService(repo, jwtService)
	categoryService := service.NewCategoryService(repo)
	transactionService := service.NewTransactionService(repo)

	// Initialize handlers
	userHandler := handlers.NewUserHandler(userService)
	categoryHandler := handlers.NewCategoryHandler(categoryService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)

	// Setup Gin router
	router := gin.Default()

	routes.SetupRoutes(
		router,
		userHandler,
		categoryHandler,
		transactionHandler,
		jwtService,
	)

	// Run the server
	if err := router.Run(":8080"); err != nil {
		log.Error("Failed to start server: %v", err)
	}

}
