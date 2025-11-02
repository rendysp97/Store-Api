package model

import "time"

type Transaction struct {
	Id                int       `gorm:"primaryKey;autoIncrement"`
	Id_user           int       `gorm:"not null"`
	Alamat_pengiriman int       `gorm:"not null"`
	Harga_total       int       `json:"harga_total"`
	Kode_invoice      string    `json:"kode_invoice"`
	Method_bayar      string    `json:"method_bayar"`
	Updated_at        time.Time `gorm:"autoUpdateTime"`
	Created_at        time.Time `gorm:"autoCreateTime"`

	User   User                `gorm:"foreignKey:Id_user;constraint:OnDelete:CASCADE" json:"-"`
	Alamat Alamat              `gorm:"foreignKey:alamat_pengiriman;constraint:OnDelete:CASCADE" json:"alamat"`
	Detail []DetailTransaction `gorm:"foreignKey:Id_trx;references:Id;constraint:OnDelete:CASCADE" json:"detail"`
}

type ProductInput struct {
	ProductID int `json:"product_id"`
	Kuantitas int `json:"kuantitas"`
}

type CreateTransactionInput struct {
	MethodBayar string         `json:"method_bayar"`
	AlamatKirim int            `json:"alamat_kirim"`
	Products    []ProductInput `json:"detail_trx"`
}
