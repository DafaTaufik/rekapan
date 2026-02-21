package handler

import (
	"net/http"
	"rekap-backend/config"
	"time"

	"github.com/gin-gonic/gin"
)

// DailySummaryResult holds the aggregated data for a single day
type DailySummaryResult struct {
	Date              string  `json:"date"`
	TotalTransactions int64   `json:"total_transactions"`
	TotalRevenue      float64 `json:"total_revenue"`
	TotalKg           float64 `json:"total_kg"`
	TotalPc           int64   `json:"total_pc"`
	TotalPaid         int64   `json:"total_paid"` // Count of transactions with status_pembayaran = 'lunas'
}

// GetDailySummary returns the summary for a single day.
// Query param: date (YYYY-MM-DD), defaults to today. Optional: branch_id
func GetDailySummary(c *gin.Context) {
	dateStr := c.DefaultQuery("date", time.Now().Format("2006-01-02"))

	parsed, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format, use: YYYY-MM-DD"})
		return
	}

	start := parsed
	end := parsed.Add(24 * time.Hour)

	query := config.DB.Table("transactions").
		Where("tanggal_masuk >= ? AND tanggal_masuk < ?", start, end)

	// Optional branch filter
	if branchID := c.Query("branch_id"); branchID != "" {
		query = query.Where("branch_id = ?", branchID)
	}

	var result DailySummaryResult
	query.Select(`
		TO_CHAR(tanggal_masuk, 'YYYY-MM-DD') as date,
		COUNT(*) as total_transactions,
		COALESCE(SUM(total), 0) as total_revenue,
		COALESCE(SUM(jumlah_kg), 0) as total_kg,
		COALESCE(SUM(jumlah_pc), 0) as total_pc,
		COUNT(CASE WHEN status_pembayaran = 'lunas' THEN 1 END) as total_paid
	`).Scan(&result)

	result.Date = dateStr

	c.JSON(http.StatusOK, gin.H{"data": result})
}

// RangeSummaryResult holds aggregated data grouped by day for a date range
type RangeSummaryResult struct {
	Date              string  `json:"date"`
	TotalTransactions int64   `json:"total_transactions"`
	TotalRevenue      float64 `json:"total_revenue"`
	TotalKg           float64 `json:"total_kg"`
	TotalPc           int64   `json:"total_pc"`
}

// GetRangeSummary returns a per-day breakdown within a date range.
// Query params: start_date, end_date (YYYY-MM-DD). Optional: branch_id
func GetRangeSummary(c *gin.Context) {
	startStr := c.Query("start_date")
	endStr := c.Query("end_date")

	if startStr == "" || endStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required"})
		return
	}

	startDate, err := time.Parse("2006-01-02", startStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start_date format, use: YYYY-MM-DD"})
		return
	}

	endDate, err := time.Parse("2006-01-02", endStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end_date format, use: YYYY-MM-DD"})
		return
	}

	// Make end_date inclusive by adding 1 day
	endDate = endDate.Add(24 * time.Hour)

	query := config.DB.Table("transactions").
		Where("tanggal_masuk >= ? AND tanggal_masuk < ?", startDate, endDate)

	// Optional branch filter
	if branchID := c.Query("branch_id"); branchID != "" {
		query = query.Where("branch_id = ?", branchID)
	}

	var results []RangeSummaryResult
	query.Select(`
		TO_CHAR(tanggal_masuk, 'YYYY-MM-DD') as date,
		COUNT(*) as total_transactions,
		COALESCE(SUM(total), 0) as total_revenue,
		COALESCE(SUM(jumlah_kg), 0) as total_kg,
		COALESCE(SUM(jumlah_pc), 0) as total_pc
	`).Group("TO_CHAR(tanggal_masuk, 'YYYY-MM-DD')").
		Order("date ASC").
		Scan(&results)

	c.JSON(http.StatusOK, gin.H{
		"data":       results,
		"start_date": startStr,
		"end_date":   endStr,
	})
}
