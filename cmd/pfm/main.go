package main

import (
	"github.com/TakuroBreath/personal-finance-manager/internal/config"
	"github.com/TakuroBreath/personal-finance-manager/internal/storage"
	"github.com/TakuroBreath/personal-finance-manager/pkg/logger"
	"github.com/TakuroBreath/personal-finance-manager/pkg/sl"
	"github.com/joho/godotenv"
	"log"
	"log/slog"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ConfigPath := os.Getenv("CONFIG_PATH")

	cfg := config.MustLoad(ConfigPath)

	log := logger.SetupLogger(cfg.Env)

	log.Info("Starting app", slog.String("env", cfg.Env))

	_, err = storage.NewStorage(cfg)
	if err != nil {
		log.Error("failed to connect to database", sl.Err(err))
		os.Exit(1)
	}

	//TODO: routes

	//TODO: run

	//TODO: graceful shutdown

}
