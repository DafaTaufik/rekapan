package model

import "time"

// Transaction represents a single laundry transaction record
type Transaction struct {
	ID               int       `gorm:"primaryKey;autoIncrement" json:"id"`
	BranchID         int       `gorm:"column:branch_id" json:"branch_id"`
	NoTransaksi      string    `gorm:"column:no_transaksi" json:"no_transaksi"`
	TanggalMasuk     time.Time `gorm:"column:tanggal_masuk" json:"tanggal_masuk"`
	NamaPelanggan    string    `gorm:"column:nama_pelanggan" json:"nama_pelanggan"`
	Status           string    `gorm:"column:status" json:"status"`
	StatusPembayaran string    `gorm:"column:status_pembayaran" json:"status_pembayaran"`
	DP               int       `gorm:"column:dp" json:"dp"`
	Pelunasan        string    `gorm:"column:pelunasan" json:"pelunasan"`
	Subtotal         int       `gorm:"column:subtotal" json:"subtotal"`
	BiayaAntarJemput int       `gorm:"column:biaya_antar_jemput" json:"biaya_antar_jemput"`
	Diskon           int       `gorm:"column:diskon" json:"diskon"`
	DiskonPoin       int       `gorm:"column:diskon_poin" json:"diskon_poin"`
	Total            int       `gorm:"column:total" json:"total"`
	JumlahKg         float64   `gorm:"column:jumlah_kg" json:"jumlah_kg"`
	JumlahPc         int       `gorm:"column:jumlah_pc" json:"jumlah_pc"`
	CreatedAt        time.Time `gorm:"column:created_at" json:"created_at"`
}

// TableName specifies the database table name for GORM
func (Transaction) TableName() string {
	return "transactions"
}
