package model

import "time"

type Alamat struct {
	Id            int       `gorm:"primaryKey;autoIncrement"`
	Id_user       int       `gorm:"not null"`
	Judul         string    `json:"judul_alamat"`
	Nama_penerima string    `json:"nama_penerima"`
	No_telp       string    `json:"no_telp"`
	Detail_Alamat string    `json:"detail_alamat"`
	Updated_at    time.Time `gorm:"autoUpdateTime"`
	Created_at    time.Time `gorm:"autoCreateTime"`
}
