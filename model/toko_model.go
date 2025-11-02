package model

import "time"

type Toko struct {
	Id         int       `gorm:"primaryKey;autoIncrement"`
	Id_user    int       `gorm:"not null"`
	Nama_toko  string    `json:"nama_toko"`
	Url_foto   *string   `json:"url_foto"`
	Updated_at time.Time `gorm:"autoUpdateTime"`
	Created_at time.Time `gorm:"autoCreateTime"`
}
