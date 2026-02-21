package handler

import (
	"net/http"
	"rekap-backend/config"
	"rekap-backend/model"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// GetTransactions returns a paginated list of transactions with optional filters.
// Query params: date (YYYY-MM-DD), branch_id, status, page, limit
func GetTransactions(c *gin.Context) {
	var transactions []model.Transaction

	query := config.DB.Model(&model.Transaction{})

	// Filter by date
	if date := c.Query("date"); date != "" {
		parsed, err := time.Parse("2006-01-02", date)
		if err == nil {
			start := parsed
			end := parsed.Add(24 * time.Hour)
			query = query.Where("tanggal_masuk >= ? AND tanggal_masuk < ?", start, end)
		}
	}

	// Filter by branch
	if branchID := c.Query("branch_id"); branchID != "" {
		query = query.Where("branch_id = ?", branchID)
	}

	// Filter by status
	if status := c.Query("status"); status != "" {
		query = query.Where("status = ?", status)
	}

	// Pagination
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "20"))
	if page < 1 {
		page = 1
	}
	offset := (page - 1) * limit

	// Count total records
	var total int64
	query.Count(&total)

	// Fetch records
	result := query.Order("tanggal_masuk DESC").Limit(limit).Offset(offset).Find(&transactions)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  transactions,
		"total": total,
		"page":  page,
		"limit": limit,
	})
}

// GetTransactionByID returns the detail of a single transaction
func GetTransactionByID(c *gin.Context) {
	id := c.Param("id")

	var transaction model.Transaction
	result := config.DB.First(&transaction, id)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transaction})
}
