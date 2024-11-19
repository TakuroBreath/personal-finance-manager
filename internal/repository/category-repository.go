package repository

import "github.com/TakuroBreath/personal-finance-manager/internal/models"

type CategoryRepository interface {
	CreateCategory(category models.Category) (uint, error)
	UpdateCategory(category models.Category) error
	DeleteCategory(id uint) error
	GetCategoryByID(id uint) (*models.Category, error)
	GetCategoriesByUserID(userID uint) ([]models.Category, error)
	GetAllCategories() ([]models.Category, error)
}

func (r *Repository) CreateCategory(category models.Category) (uint, error) {
	result := r.db.Create(&category)
	if result.Error != nil {
		return 0, result.Error
	}
	return category.ID, nil
}

func (r *Repository) UpdateCategory(category models.Category) error {
	result := r.db.Save(&category)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) DeleteCategory(id uint) error {
	result := r.db.Delete(&models.Category{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *Repository) GetCategoryByID(id uint) (*models.Category, error) {
	var category models.Category
	result := r.db.Where("id = ?", id).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (r *Repository) GetCategoriesByUserID(userID uint) ([]models.Category, error) {
	var categories []models.Category
	result := r.db.Where("user_id = ?", userID).Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}

func (r *Repository) GetAllCategories() ([]models.Category, error) {
	var categories []models.Category
	result := r.db.Find(&categories)
	if result.Error != nil {
		return nil, result.Error
	}
	return categories, nil
}
