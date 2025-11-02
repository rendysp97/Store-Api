package repository

import (
	"fmt"
	"store-api/model"

	"gorm.io/gorm"
)

func CreateCategoryRepo(db *gorm.DB, category *model.Category) (int, error) {

	if err := db.Create(category).Error; err != nil {
		return 0, err
	}
	return category.Id, nil

}

func GetAllCategoryRepo(db *gorm.DB) ([]model.Category, error) {

	var categories []model.Category

	err := db.Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil

}

func GetCategoryByIdRepo(db *gorm.DB, id int) (*model.Category, error) {

	var category model.Category

	err := db.Where("id = ? ", id).First(&category).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &category, nil
}

func UpdateCategoryByIdRepo(db *gorm.DB, id int, updated map[string]any) error {
	result := db.Model(&model.Category{}).Where("id = ?", id).Updates(updated)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("no category found with id %d", id)
	}

	return nil
}

func DeleteCategoryByIDRepo(db *gorm.DB, id int) (int64, error) {

	result := db.Delete(model.Category{}, "id = ?", id)

	return result.RowsAffected, nil

}
