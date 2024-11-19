package repository

import (
	"github.com/TakuroBreath/personal-finance-manager/internal/models"
)

type UserRepository interface {
	CreateUser(user models.User) (uint, error)
	GetUserByEmail(email string) (*models.User, error)
	GetUserByID(id uint) (*models.User, error)
	UpdateUser(user models.User) error
	DeleteUser(id uint) error
}

func (r *Repository) CreateUser(user models.User) (uint, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.ID, nil
}

func (r *Repository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *Repository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := r.db.Where("id = ?", id).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (r *Repository) UpdateUser(user models.User) error {
	result := r.db.Save(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) DeleteUser(id uint) error {
	result := r.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
