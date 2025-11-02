package repository

import (
	"fmt"
	"store-api/model"

	"gorm.io/gorm"
)

func CreateProductRepository(db *gorm.DB, anyData any) error {

	return db.Create(anyData).Error

}

func GetProductById(db *gorm.DB, id, userId int) (*model.Product, error) {

	var product model.Product

	err := db.Joins("JOIN tokos ON tokos.id = products.id_toko").
		Where("products.id = ? AND tokos.id_user = ?", id, userId).
		Preload("Toko").
		Preload("Category").
		Preload("FotoProduk").
		First(&product).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, err
	}

	return &product, nil
}

func GetAllProductRepo(db *gorm.DB, id, limit, offset int, filters map[string]string) ([]model.Product, error) {
	var product []model.Product

	res := db.
		Joins("JOIN tokos ON tokos.id = products.id_toko").
		Preload("Toko").
		Preload("Category").
		Preload("FotoProduk")

	if nama, ok := filters["nama_produk"]; ok && nama != "" {
		res = res.Where("products.nama_produk LIKE ?", "%"+nama+"%")
	}
	if cat, ok := filters["category_id"]; ok && cat != "" {
		res = res.Where("products.id_category = ?", cat)
	}
	if toko, ok := filters["toko_id"]; ok && toko != "" {
		res = res.Where("products.id_toko = ?", toko)
	}
	if min, ok := filters["min_harga"]; ok && min != "" {
		res = res.Where("products.harga_konsumen >= ?", min)
	}
	if max, ok := filters["max_harga"]; ok && max != "" {
		res = res.Where("products.harga_konsumen <= ?", max)
	}

	err := res.Limit(limit).Offset(offset).Find(&product).Error
	if err != nil {
		return nil, err
	}

	return product, nil
}

func UpdateProductRepo(db *gorm.DB, id, user_id int, updateData map[string]any) error {

	var product model.Product

	err := db.Joins("JOIN tokos ON tokos.id = products.id_toko").
		Where("products.id = ? AND tokos.id_user = ?", id, user_id).
		First(&product).Error

	if err != nil {
		return err
	}

	return db.Model(&product).Updates(updateData).Error

}

func DeleteProductByIDRepo(db *gorm.DB, id, userid int) error {

	var product model.Product

	err := db.Joins("JOIN tokos ON tokos.id = products.id_toko").
		Where("products.id = ? AND tokos.id_user = ?", id, userid).
		First(&product).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return fmt.Errorf("product not found or not yours")
		}
		return err
	}

	return db.Delete(&product).Error

}
