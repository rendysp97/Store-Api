package repository

import (
	"store-api/model"

	"gorm.io/gorm"
)

func GetAddressRepo(db *gorm.DB, id int) ([]model.Alamat, error) {

	var address []model.Alamat

	err := db.Where("id_user = ?", id).Find(&address).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return address, nil
}

func GetAlamatByIdRepo(db *gorm.DB, id, userId int) (*model.Alamat, error) {

	var address model.Alamat

	err := db.Where("id = ? AND id_user = ? ", id, userId).First(&address).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &address, err

}

func CreateAddressRepo(db *gorm.DB, address *model.Alamat) (int, error) {
	if err := db.Create(address).Error; err != nil {
		return 0, err
	}
	return address.Id, nil

}

func UpdateAddressRepo(db *gorm.DB, id, userId int, updated map[string]any) int64 {
	result := db.Model(model.Alamat{}).Where("id = ? And id_user = ?", id, userId).Updates(updated)

	return result.RowsAffected
}

func DeleteAlamatByIdRepo(db *gorm.DB, id, userId int) int64 {

	result := db.Delete(model.Alamat{}, "id  = ? and id_user = ?", id, userId)

	return result.RowsAffected
}
