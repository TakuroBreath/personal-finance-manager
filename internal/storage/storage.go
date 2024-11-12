package storage

import (
	"fmt"
	"github.com/TakuroBreath/personal-finance-manager/internal/config"
	"github.com/TakuroBreath/personal-finance-manager/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewStorage(cfg *config.Config) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch cfg.Env {
	case "local":
		dialector = sqlite.Open(cfg.StoragePath)
	case "prod":
		panic("not implemented")
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	if err := models.AutoMigrate(db); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %w", err)
	}

	return db, nil
}
