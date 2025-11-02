package model

import "time"

type DetailTransaction struct {
	Id            int       `gorm:"primaryKey;autoIncrement"`
	Id_trx        int       `gorm:"not null"`
	Id_log_produk int       `gorm:"not null"`
	Id_toko       int       `gorm:"not null"`
	Kuantitas     int       `json:"kuantitas"`
	Harga_total   int       `json:"harga_total"`
	Updated_at    time.Time `gorm:"autoUpdateTime"`
	Created_at    time.Time `gorm:"autoCreateTime"`

	LogProduk LogProduk `gorm:"foreignKey:Id_log_produk;constraint:OnDelete:CASCADE" json:"log_produk"`
}
