package repository

import (
	"fmt"
	"store-api/model"

	"gorm.io/gorm"
)

func GetMyProfileRepo(db *gorm.DB, id int) (*model.User, error) {

	var user model.User

	err := db.Where("id = ?", id).First(&user).Error

	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}

	return &user, nil
}

func ProfileUpdateRepo(db *gorm.DB, id int, updated map[string]any) error {
	result := db.Model(&model.User{}).Where("id = ?", id).Updates(updated)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no user found with id %d", id)
	}

	return nil
}
