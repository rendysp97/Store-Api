package model

import "time"

type LogProduk struct {
	Id             int       `gorm:"primaryKey;autoIncrement"`
	Nama_produk    string    `json:"nama_produk"`
	Slug           string    `json:"slug"`
	Harga_reseller string    `json:"harga_reseller"`
	Harga_konsumen string    `json:"harga_konsumen"`
	Deskripsi      string    `json:"deskripsi"`
	Updated_at     time.Time `gorm:"autoUpdateTime"`
	Created_at     time.Time `gorm:"autoCreateTime"`
	Id_toko        int       `gorm:"not null"`
	Id_category    int       `gorm:"not null"`

	DetailTransactions []DetailTransaction `gorm:"foreignKey:Id_log_produk;constraint:OnDelete:CASCADE" json:"-"`
}
