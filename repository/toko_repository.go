package repository

import (
	"store-api/model"

	"gorm.io/gorm"
)

func CreateStoreRepo(db *gorm.DB, toko *model.Toko) error {
	return db.Create(toko).Error
}

func GetMyTokoRepo(db *gorm.DB, id int) (*model.Toko, error) {

	var toko model.Toko

	err := db.Where("id_user = ?", id).First(&toko).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &toko, nil
}

func UpdateStoreByIDRepo(db *gorm.DB, id int, updateData map[string]any) error {
	return db.Model(&model.Toko{}).
		Where("id = ?", id).
		Updates(updateData).
		Error
}

func GetTokoByIDRepo(db *gorm.DB, id int) (*model.Toko, error) {

	var toko model.Toko

	err := db.Where("id = ?", id).First(&toko).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &toko, nil

}

func GetAllTokoRepo(db *gorm.DB, limit, page int, name string) ([]model.Toko, int64, error) {

	var toko []model.Toko

	var total int64

	query := db.Model(&model.Toko{})

	if name != "" {
		query = query.Where("nama_toko like ? ", "%"+name+"%")
	}

	query.Count(&total)

	err := query.Offset((page - 1) * limit).Limit(limit).Find(&toko).Error

	if err != nil {
		return nil, 0, err
	}

	return toko, total, nil
}

func GetTokoByUserIdRepo(db *gorm.DB, id int) (*model.Toko, error) {

	var toko model.Toko

	err := db.Where("id_user = ?", id).First(&toko).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &toko, nil

}
