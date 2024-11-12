package models

import (
	"gorm.io/gorm"
	"time"
)

type TransactionType string

const (
	Income  TransactionType = "income"
	Expense TransactionType = "expense"
)

// Transaction представляет финансовую транзакцию
type Transaction struct {
	gorm.Model
	Amount      float64         `gorm:"type:decimal(10,2);not null" json:"amount"`
	Type        TransactionType `gorm:"type:varchar(10);not null" json:"type"`
	Date        time.Time       `gorm:"not null;index" json:"date"`
	Description string          `gorm:"type:varchar(200)" json:"description"`
	Notes       string          `gorm:"type:text" json:"notes,omitempty"`
	CategoryID  uint            `gorm:"not null" json:"categoryId"`
	Category    Category        `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	UserID      uint            `gorm:"not null" json:"userId"`
	User        User            `gorm:"foreignKey:UserID" json:"-"`
}

func (Transaction) TableName() string {
	return "transactions"
}

func (t *Transaction) BeforeCreate(tx *gorm.DB) error {
	if t.Date.IsZero() {
		t.Date = time.Now()
	}
	return nil
}
