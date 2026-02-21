package handler

import (
	"net/http"
	"rekap-backend/config"

	"github.com/gin-gonic/gin"
)

// BranchResult holds aggregated statistics per branch
type BranchResult struct {
	BranchID          int     `json:"branch_id"`
	TotalTransactions int64   `json:"total_transactions"`
	TotalRevenue      float64 `json:"total_revenue"`
}

// GetBranches returns a list of branches with their transaction statistics
func GetBranches(c *gin.Context) {
	var branches []BranchResult

	result := config.DB.Table("transactions").
		Select(`
			branch_id,
			COUNT(*) as total_transactions,
			COALESCE(SUM(total), 0) as total_revenue
		`).
		Group("branch_id").
		Order("branch_id ASC").
		Scan(&branches)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch branches"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": branches})
}
