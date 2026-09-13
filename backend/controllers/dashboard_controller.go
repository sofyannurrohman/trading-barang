package controllers

import (
	"backend/db"
	"backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LowStockItem struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Stock int    `json:"stock"`
}

type MonthlySalePoint struct {
	Date  string  `json:"date"`
	Total float64 `json:"total"`
}

type DashboardStats struct {
	TotalRevenueThisMonth  float64            `json:"total_revenue_this_month"`
	TotalPPNThisMonth      float64            `json:"total_ppn_this_month"`
	TotalInvoicesThisMonth int64              `json:"total_invoices_this_month"`
	LowStockProducts       []LowStockItem     `json:"low_stock_products"`
	MonthlySalesChart      []MonthlySalePoint `json:"monthly_sales_chart"`
}

func Dashboard(c *gin.Context) {
	now := time.Now()
	startOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	thirtyDaysAgo := now.AddDate(0, 0, -30)

	// --- Total Penjualan & PPN Bulan Ini ---
	var totalRevenue, totalPPN float64
	var totalInvoices int64
	db.DB.Model(&models.Invoice{}).
		Where("created_at >= ? AND deleted_at IS NULL", startOfMonth).
		Select("COALESCE(SUM(total_dpp), 0), COALESCE(SUM(total_ppn), 0), COUNT(*)").
		Row().Scan(&totalRevenue, &totalPPN, &totalInvoices)

	// --- Produk Stok Menipis (threshold: <= 10 unit) ---
	var lowStockProducts []LowStockItem
	db.DB.Model(&models.Product{}).
		Where("current_stock <= 10 AND deleted_at IS NULL").
		Select("id, name, current_stock as stock").
		Order("current_stock asc").
		Scan(&lowStockProducts)
	if lowStockProducts == nil {
		lowStockProducts = []LowStockItem{}
	}

	// --- Grafik Penjualan 30 Hari Terakhir (per hari, grand_total) ---
	var chartData []MonthlySalePoint
	db.DB.Model(&models.Invoice{}).
		Where("created_at >= ? AND deleted_at IS NULL", thirtyDaysAgo).
		Select("TO_CHAR(created_at, 'YYYY-MM-DD') as date, SUM(grand_total) as total").
		Group("TO_CHAR(created_at, 'YYYY-MM-DD')").
		Order("date asc").
		Scan(&chartData)
	if chartData == nil {
		chartData = []MonthlySalePoint{}
	}

	c.JSON(http.StatusOK, DashboardStats{
		TotalRevenueThisMonth:  totalRevenue,
		TotalPPNThisMonth:      totalPPN,
		TotalInvoicesThisMonth: totalInvoices,
		LowStockProducts:       lowStockProducts,
		MonthlySalesChart:      chartData,
	})
}
