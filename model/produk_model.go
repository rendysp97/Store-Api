package model

import "time"

type Product struct {
	Id             int       `gorm:"primaryKey;autoIncrement"`
	Nama_produk    string    `json:"nama_produk"`
	Slug           string    `json:"slug"`
	Harga_reseller string    `json:"harga_reseller"`
	Harga_konsumen string    `json:"harga_konsumen"`
	Stok           int       `json:"stok"`
	Deskripsi      string    `json:"deskripsi"`
	Id_toko        int       `gorm:"not null" json:"-"`
	Id_category    int       `gorm:"not null" json:"-"`
	Created_at     time.Time `gorm:"autoCreateTime"`
	Updated_at     time.Time `gorm:"autoUpdateTime"`

	Toko       Toko         `gorm:"foreignKey:Id_toko;constraint:OnDelete:CASCADE" json:"toko"`
	Category   Category     `gorm:"foreignKey:Id_category;constraint:OnDelete:CASCADE" json:"category"`
	FotoProduk []FotoProduk `gorm:"foreignKey:Id_produk;constraint:OnDelete:CASCADE" json:"photos"`
}
