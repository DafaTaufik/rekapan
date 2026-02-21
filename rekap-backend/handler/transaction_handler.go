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

	// Fetch records — latest date first, earliest time within the same day first
	result := query.Order("DATE(tanggal_masuk) DESC, tanggal_masuk ASC").Limit(limit).Offset(offset).Find(&transactions)
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

// GetTransactionByTrxID finds a transaction by the trailing sequence number of no_transaksi.
// Example: trx_id=01444 will match no_transaksi = 'TRX/260116/01444'
func GetTransactionByTrxID(c *gin.Context) {
	trxID := c.Param("trx_id")

	var transaction model.Transaction
	result := config.DB.Where("no_transaksi LIKE ?", "%/"+trxID).First(&transaction)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transaction})
}

func GetTransactionByBranchID(c *gin.Context) {
	branchID := c.Param("branch_id")

	var transactions []model.Transaction
	result := config.DB.Where("branch_id = ?", branchID).Order("DATE(tanggal_masuk) DESC, tanggal_masuk ASC").Find(&transactions)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transactions not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": transactions})
}

// TogglePaymentStatus toggles status_pembayaran between 'lunas' and 'belum lunas'
func TogglePaymentStatus(c *gin.Context) {
	id := c.Param("id")

	// Find the transaction first
	var transaction model.Transaction
	if err := config.DB.First(&transaction, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transaction not found"})
		return
	}

	// Toggle the status
	newStatus := "lunas"
	if transaction.StatusPembayaran == "lunas" {
		newStatus = "belum lunas"
	}

	// Update only the status_pembayaran column
	if err := config.DB.Model(&transaction).Update("status_pembayaran", newStatus).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update payment status"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    transaction,
		"message": "Payment status updated to: " + newStatus,
	})
}

