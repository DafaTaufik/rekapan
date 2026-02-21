package model

import "time"

// Transaction represents the transactions table structure
type Transaction struct {
	ID                 int       `gorm:"primaryKey;autoIncrement" json:"id"`
	BranchID           int       `gorm:"column:branch_id" json:"branch_id"`
	NoTransaksi        string    `gorm:"column:no_transaksi" json:"no_transaksi"`
	TanggalMasuk       time.Time `gorm:"column:tanggal_masuk" json:"tanggal_masuk"`
	NamaPelanggan      string    `gorm:"column:nama_pelanggan" json:"nama_pelanggan"`
	Status             string    `gorm:"column:status" json:"status"`
	StatusPembayaran   string    `gorm:"column:status_pembayaran" json:"status_pembayaran"`
	DP                 float64   `gorm:"column:dp" json:"dp"`
	Pelunasan          float64   `gorm:"column:pelunasan" json:"pelunasan"`
	Subtotal           float64   `gorm:"column:subtotal" json:"subtotal"`
	BiayaAntarJemput   float64   `gorm:"column:biaya_antar_jemput" json:"biaya_antar_jemput"`
	Diskon             float64   `gorm:"column:diskon" json:"diskon"`
	DiskonPoin         float64   `gorm:"column:diskon_poin" json:"diskon_poin"`
	Total              float64   `gorm:"column:total" json:"total"`
	JumlahKg           float64   `gorm:"column:jumlah_kg" json:"jumlah_kg"`
	JumlahPc           int       `gorm:"column:jumlah_pc" json:"jumlah_pc"`
	CreatedAt          time.Time `gorm:"column:created_at" json:"created_at"`
}

func (Transaction) TableName() string {
	return "tbl_transactions"
}
