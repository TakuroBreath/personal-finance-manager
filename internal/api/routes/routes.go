package routes

import (
	"github.com/TakuroBreath/personal-finance-manager/internal/api/handlers"
	"github.com/TakuroBreath/personal-finance-manager/internal/api/middleware"
	"github.com/TakuroBreath/personal-finance-manager/internal/service"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(
	router *gin.Engine,
	userHandler *handlers.UserHandler,
	categoryHandler *handlers.CategoryHandler,
	transactionHandler *handlers.TransactionHandler,
	jwtService *service.JWTService,
) {
	// Public routes
	public := router.Group("/api/v1")
	{
		public.POST("/users/register", userHandler.CreateUser)
		public.POST("/users/login", userHandler.LoginUser)
	}

	// Protected routes
	protected := router.Group("/api/v1")
	protected.Use(middleware.AuthMiddleware(jwtService))
	{
		// User routes
		protected.PUT("/users", userHandler.UpdateUser)
		protected.DELETE("/users", userHandler.DeleteUser)

		// Category routes
		protected.POST("/categories", categoryHandler.CreateCategory)
		protected.PUT("/categories/:cat_id", categoryHandler.UpdateCategory)
		protected.DELETE("/categories/:cat_id", categoryHandler.DeleteCategory)
		protected.GET("/categories", categoryHandler.GetCategoriesByUserID)
		protected.GET("/categories/:cat_id", categoryHandler.GetCategoryByID)

		// Transaction routes
		protected.POST("/transactions", transactionHandler.CreateTransaction)
		protected.PUT("/transactions", transactionHandler.UpdateTransaction)
		protected.DELETE("/transactions", transactionHandler.DeleteTransaction)
		protected.GET("/transactions", transactionHandler.GetTransactionsByUserID)
		protected.GET("/transactions/:trans_id", transactionHandler.GetTransactionByID)
	}
}
