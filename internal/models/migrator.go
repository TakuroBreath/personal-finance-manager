package models

import (
	"fmt"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	// Создаем таблицы и базовые индексы
	err := db.AutoMigrate(&User{}, &Category{}, &Transaction{})
	if err != nil {
		return fmt.Errorf("auto migration failed: %w", err)
	}

	// Создаем дополнительные индексы
	// Используем Raw для поддержки разных диалектов БД
	dialect := db.Dialector.Name()

	if dialect == "postgres" {
		// Для PostgreSQL
		if err := db.Exec(`CREATE INDEX IF NOT EXISTS idx_transactions_user_date 
            ON transactions(user_id, date)`).Error; err != nil {
			return fmt.Errorf("failed to create user_date index: %w", err)
		}

		if err := db.Exec(`CREATE INDEX IF NOT EXISTS idx_transactions_category 
            ON transactions(category_id)`).Error; err != nil {
			return fmt.Errorf("failed to create category index: %w", err)
		}
	} else if dialect == "sqlite" {
		// Для SQLite
		if err := db.Exec(`CREATE INDEX IF NOT EXISTS idx_transactions_user_date 
            ON transactions(user_id, date)`).Error; err != nil {
			return fmt.Errorf("failed to create user_date index: %w", err)
		}

		if err := db.Exec(`CREATE INDEX IF NOT EXISTS idx_transactions_category 
            ON transactions(category_id)`).Error; err != nil {
			return fmt.Errorf("failed to create category index: %w", err)
		}
	}

	return nil
}
