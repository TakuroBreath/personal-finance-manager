package service

import "github.com/TakuroBreath/personal-finance-manager/internal/repository"

type Service struct {
	transactionRepository *repository.Repository
}

func NewService(transactionRepository *repository.Repository) *Service {
	return &Service{
		transactionRepository: transactionRepository,
	}
}
