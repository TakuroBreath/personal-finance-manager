package service

import (
	"errors"
	"github.com/TakuroBreath/personal-finance-manager/internal/models"
	"github.com/TakuroBreath/personal-finance-manager/internal/repository"
	"time"
)

type TransactionServiceImpl interface {
	CreateTransaction(request TransactionCreateRequest) (Transaction, error)
	UpdateTransaction(request TransactionUpdateRequest) error
	DeleteTransaction(request TransactionDeleteRequest) error
	GetTransactionByID(id uint) (Transaction, error)
	GetTransactionsByUserID(userID uint) ([]Transaction, error)
}

type TransactionService struct {
	repo repository.TransactionRepository
}

func NewTransactionService(repo repository.TransactionRepository) *TransactionService {
	return &TransactionService{repo: repo}
}

type TransactionCreateRequest struct {
	Amount      float64                `json:"amount"`
	Type        models.TransactionType `json:"type"`
	Date        string
	Description string `json:"description"`
	Notes       string `json:"notes"`
	CategoryID  uint   `json:"category"`
	UserID      uint
}

type TransactionUpdateRequest struct {
	ID          uint                   `json:"id"`
	Amount      float64                `json:"amount"`
	Type        models.TransactionType `json:"type"`
	Date        string
	Description string `json:"description"`
	Notes       string `json:"notes"`
	CategoryID  uint   `json:"category"`
	UserID      uint
}

type TransactionDeleteRequest struct {
	ID uint
}

type Transaction struct {
	ID          uint
	Amount      float64
	Type        models.TransactionType
	Date        string
	Description string
	Notes       string
	CategoryID  uint
	Category    models.Category
	UserID      uint
	User        models.User
}

func (s *TransactionService) CreateTransaction(request TransactionCreateRequest) (uint, error) {
	transaction := models.Transaction{
		Amount:      request.Amount,
		Type:        request.Type,
		Date:        time.Now(),
		Description: request.Description,
		Notes:       request.Notes,
		CategoryID:  request.CategoryID,
		UserID:      request.UserID,
	}

	return s.repo.CreateTransaction(transaction)
}

func (s *TransactionService) UpdateTransaction(request TransactionUpdateRequest) error {
	transaction, err := s.repo.GetTransactionByID(request.ID)
	if err != nil {
		return errors.New("transaction not found")
	}
	transaction.Amount = request.Amount
	transaction.Type = request.Type
	transaction.Date = time.Now()
	transaction.Description = request.Description
	transaction.Notes = request.Notes
	transaction.CategoryID = request.CategoryID
	transaction.UserID = request.UserID

	return s.repo.UpdateTransaction(transaction)
}

func (s *TransactionService) DeleteTransaction(request TransactionDeleteRequest) error {
	transaction, err := s.repo.GetTransactionByID(request.ID)
	if err != nil {
		return errors.New("transaction not found")
	}
	return s.repo.DeleteTransaction(transaction)
}

func (s *TransactionService) GetTransactionByID(id uint) (Transaction, error) {
	transaction, err := s.repo.GetTransactionByID(id)
	if err != nil {
		return Transaction{}, err
	}
	return Transaction{
		ID:          transaction.ID,
		Amount:      transaction.Amount,
		Type:        transaction.Type,
		Date:        transaction.Date.Format("2006-01-02"),
		Description: transaction.Description,
		Notes:       transaction.Notes,
		CategoryID:  transaction.CategoryID,
		Category:    transaction.Category,
		UserID:      transaction.UserID,
		User:        transaction.User,
	}, nil
}

func (s *TransactionService) GetTransactionsByUserID(userID uint) ([]Transaction, error) {
	transactions, err := s.repo.GetTransactionsByUserID(userID)
	if err != nil {
		return nil, err
	}
	var result []Transaction
	for _, transaction := range transactions {
		result = append(result, Transaction{
			ID:          transaction.ID,
			Amount:      transaction.Amount,
			Type:        transaction.Type,
			Date:        transaction.Date.Format("2006-01-02"),
			Description: transaction.Description,
			Notes:       transaction.Notes,
			CategoryID:  transaction.CategoryID,
			Category:    transaction.Category,
			UserID:      transaction.UserID,
			User:        transaction.User,
		})
	}
	return result, nil
}
