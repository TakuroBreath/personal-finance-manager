package handlers

import (
	"github.com/TakuroBreath/personal-finance-manager/internal/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type TransactionHandler struct {
	transactionService *service.TransactionService
}

func NewTransactionHandler(transactionService *service.TransactionService) *TransactionHandler {
	return &TransactionHandler{transactionService: transactionService}
}

func (t *TransactionHandler) CreateTransaction(c *gin.Context) {
	var transaction service.TransactionCreateRequest
	userID, _ := c.Get("userID")
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	transaction.UserID = userID.(uint)
	id, err := t.transactionService.CreateTransaction(transaction)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"id": id})
}

func (t *TransactionHandler) UpdateTransaction(c *gin.Context) {
	var transaction service.TransactionUpdateRequest
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := t.transactionService.UpdateTransaction(transaction); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Transaction updated successfully"})
}

func (t *TransactionHandler) DeleteTransaction(c *gin.Context) {
	var transaction service.TransactionDeleteRequest
	if err := c.ShouldBindJSON(&transaction); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := t.transactionService.DeleteTransaction(transaction); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Transaction deleted successfully"})
}

func (t *TransactionHandler) GetTransactionsByUserID(c *gin.Context) {
	userID, _ := c.Get("userID")

	id := userID.(uint)

	transactions, err := t.transactionService.GetTransactionsByUserID(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, transactions)
}

func (t *TransactionHandler) GetTransactionByID(c *gin.Context) {
	id := c.Param("trans_id")
	transID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	transaction, err := t.transactionService.GetTransactionByID(uint(transID))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, transaction)
}
