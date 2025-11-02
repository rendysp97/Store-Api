package database

import (
	"fmt"
	"store-api/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func ConnectDb() {

	dsn := "root:12345@tcp(127.0.0.1:3306)/store_db?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Connection Failed:", err)
	}

	Db = db

	Db.AutoMigrate(&model.User{}, &model.Toko{}, &model.Alamat{}, &model.Category{}, &model.Product{}, &model.FotoProduk{}, &model.Transaction{}, &model.DetailTransaction{}, &model.LogProduk{})

	if err != nil {
		fmt.Println("Failed to migrate database:", err)
	}

	if err != nil {
		fmt.Println("Connection Failed to Open")
	} else {
		fmt.Println("Connection Established")
	}
}
