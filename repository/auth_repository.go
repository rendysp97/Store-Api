package repository

import (
	"store-api/model"

	"gorm.io/gorm"
)

func GetUserByPhone(db *gorm.DB, phone string) (*model.User, error) {
	var user model.User
	err := db.Where("notelp = ?", phone).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func GetUserByPhoneAndEmail(db *gorm.DB, phone, email string) (*model.User, error) {
	var user model.User
	err := db.Where("notelp = ? OR email = ?", phone, email).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func CreateUser(db *gorm.DB, user *model.User) error {
	return db.Create(user).Error
}
