package model

import "time"

type Category struct {
	Id         int       `gorm:"primaryKey;autoIncrement"`
	Nama       string    `json:"nama_category"`
	Updated_at time.Time `gorm:"autoUpdateTime"`
	Created_at time.Time `gorm:"autoCreateTime"`
}
