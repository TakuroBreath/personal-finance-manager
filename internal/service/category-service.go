package service

import (
	"github.com/TakuroBreath/personal-finance-manager/internal/models"
	"github.com/TakuroBreath/personal-finance-manager/internal/repository"
)

type CategoryServiceImpl interface {
	CreateCategory(request CategoryCreateRequest) (uint, error)
	UpdateCategory(request CategoryUpdateRequest) error
	DeleteCategory(request CategoryDeleteRequest) error
	GetCategoryByID(id uint) (CategoryResponse, error)
	GetCategoriesByUserID(userID uint) ([]CategoryResponse, error)
	GetAllCategories() ([]CategoryResponse, error)
}

type CategoryCreateRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"userId"`
}

type CategoryUpdateRequest struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"userId"`
}

type CategoryDeleteRequest struct {
	ID uint `json:"id"`
}

type CategoryResponse struct {
	ID          uint   `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"userId"`
}

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) CreateCategory(request CategoryCreateRequest) (uint, error) {
	var category models.Category

	category.UserID = request.UserID
	category.Name = request.Name
	category.Description = request.Description

	return s.repo.CreateCategory(category)
}

func (s *CategoryService) UpdateCategory(request CategoryUpdateRequest) error {
	var category models.Category

	category.ID = request.ID
	category.UserID = request.UserID
	category.Name = request.Name
	category.Description = request.Description

	return s.repo.UpdateCategory(category)
}

func (s *CategoryService) DeleteCategory(request CategoryDeleteRequest) error {
	return s.repo.DeleteCategory(request.ID)
}

func (s *CategoryService) GetCategoryByID(id uint) (CategoryResponse, error) {
	res, err := s.repo.GetCategoryByID(id)
	if err != nil {
		return CategoryResponse{}, err
	}

	return CategoryResponse{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		UserID:      res.UserID,
	}, nil
}

func (s *CategoryService) GetCategoriesByUserID(userID uint) ([]CategoryResponse, error) {
	res, err := s.repo.GetCategoriesByUserID(userID)
	if err != nil {
		return []CategoryResponse{}, err
	}

	var categories []CategoryResponse

	for _, category := range res {
		categories = append(categories, CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			UserID:      category.UserID,
		})
	}

	return categories, nil
}

func (s *CategoryService) GetAllCategories() ([]CategoryResponse, error) {
	res, err := s.repo.GetAllCategories()
	if err != nil {
		return []CategoryResponse{}, err
	}

	var categories []CategoryResponse

	for _, category := range res {
		categories = append(categories, CategoryResponse{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
			UserID:      category.UserID,
		})
	}

	return categories, nil
}
