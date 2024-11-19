package handlers

import (
	"github.com/TakuroBreath/personal-finance-manager/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type CategoryHandler struct {
	categoryService *service.CategoryService
}

func NewCategoryHandler(categoryService *service.CategoryService) *CategoryHandler {
	return &CategoryHandler{categoryService: categoryService}
}

func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var request service.CategoryCreateRequest
	id, _ := c.Get("userID")

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	request.UserID = uint(id.(float64))
	categoryID, err := h.categoryService.CreateCategory(request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"category_id": categoryID})
}

func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	var request service.CategoryUpdateRequest
	id, _ := c.Get("userID")
	request.UserID = uint(id.(float64))

	categoryID := c.Param("cat_id")
	catID, _ := strconv.Atoi(categoryID)
	request.ID = uint(catID)

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := h.categoryService.UpdateCategory(request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Category updated successfully"})
}

func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
	var request service.CategoryDeleteRequest

	categoryID := c.Param("cat_id")
	catID, _ := strconv.Atoi(categoryID)
	request.ID = uint(catID)

	err := h.categoryService.DeleteCategory(request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Category deleted successfully"})
}

func (h *CategoryHandler) GetCategoriesByUserID(c *gin.Context) {
	id, _ := c.Get("userID")
	categories, err := h.categoryService.GetCategoriesByUserID(uint(id.(float64)))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, categories)
}

func (h *CategoryHandler) GetCategoryByID(c *gin.Context) {
	id := c.Param("cat_id")
	catID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	category, err := h.categoryService.GetCategoryByID(uint(catID))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, category)
}

func (h *CategoryHandler) GetAllCategories(c *gin.Context) {
	categories, err := h.categoryService.GetAllCategories()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, categories)
}
