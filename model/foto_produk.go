package model

import "time"

type FotoProduk struct {
	Id         int       `gorm:"primaryKey;autoIncrement"`
	Id_produk  int       `gorm:"not null"`
	Url        string    `json:"url"`
	Updated_at time.Time `gorm:"autoUpdateTime"`
	Created_at time.Time `gorm:"autoCreateTime"`
}
