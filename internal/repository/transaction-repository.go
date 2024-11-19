package repository

import (
	"errors"
	"github.com/TakuroBreath/personal-finance-manager/internal/models"
)

type TransactionRepository interface {
	CreateTransaction(transaction models.Transaction) (uint, error)
	UpdateTransaction(transaction models.Transaction) error
	DeleteTransaction(transaction models.Transaction) error
	GetTransactionByID(id uint) (models.Transaction, error)
	GetTransactionsByUserID(userID uint) ([]models.Transaction, error)
}

func (r *Repository) CreateTransaction(transaction models.Transaction) (uint, error) {
	result := r.db.Create(&transaction)
	if result.Error != nil {
		return 0, result.Error
	}
	return transaction.ID, nil
}

func (r *Repository) UpdateTransaction(transaction models.Transaction) error {
	existing, err := r.GetTransactionByID(transaction.ID)
	if err != nil {
		return errors.New("transaction not found")
	}
	existing.Amount = transaction.Amount
	existing.Type = transaction.Type
	existing.Date = transaction.Date
	existing.Description = transaction.Description
	existing.Notes = transaction.Notes
	existing.CategoryID = transaction.CategoryID
	existing.UserID = transaction.UserID

	result := r.db.Save(&existing)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) DeleteTransaction(transaction models.Transaction) error {
	result := r.db.Delete(&transaction)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) GetTransactionByID(id uint) (models.Transaction, error) {
	var transaction models.Transaction
	result := r.db.Preload("Category").Preload("User").First(&transaction, id)
	if result.Error != nil {
		return models.Transaction{}, result.Error
	}
	return transaction, nil
}

func (r *Repository) GetTransactionsByUserID(userID uint) ([]models.Transaction, error) {
	var transactions []models.Transaction
	result := r.db.Where("user_id = ?", userID).Preload("Category").Preload("User").Find(&transactions)
	if result.Error != nil {
		return nil, result.Error
	}
	return transactions, nil
}
