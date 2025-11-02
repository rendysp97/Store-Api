package model

import "time"

type User struct {
	Id            int        `gorm:"primaryKey;autoIncrement"`
	Nama          string     `json:"nama"`
	Kata_sandi    string     `json:"kata_sandi"`
	Notelp        string     `gorm:"unique" json:"no_telp"`
	Tanggal_lahir *time.Time `json:"tanggal_lahir"`
	Jenis_kelamin *string    `json:"jenis_kelamin"`
	Tentang       *string    `json:"tentang"`
	Pekerjaan     *string    `json:"pekerjaan"`
	Email         string     `gorm:"unique" json:"email"`
	Id_provinsi   *int       `json:"id_provinsi"`
	Id_kota       *int       `json:"id_kota"`
	Is_admin      bool       `json:"is_admin"`
	Updated_at    time.Time  `gorm:"autoUpdateTime"`
	Created_at    time.Time  `gorm:"autoCreateTime"`

	Toko   Toko     `gorm:"foreignKey:Id_user;constraint:OnDelete:CASCADE" json:"-"`
	Alamat []Alamat `gorm:"foreignKey:Id_user;constraint:OnDelete:CASCADE" json:"-"`
}
