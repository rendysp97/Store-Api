package repository

import (
	"errors"
	"fmt"
	"store-api/model"
	"strconv"
	"time"

	"gorm.io/gorm"
)

func CreateTransactionRepo(db *gorm.DB, userID int, input model.CreateTransactionInput) (int, error) {
	var transaction model.Transaction

	err := db.Transaction(func(tx *gorm.DB) error {
		transaction = model.Transaction{
			Id_user:           userID,
			Alamat_pengiriman: input.AlamatKirim,
			Method_bayar:      input.MethodBayar,
			Created_at:        time.Now(),
		}

		if err := tx.Create(&transaction).Error; err != nil {
			return err
		}

		totalHarga := 0

		for _, p := range input.Products {
			var product model.Product
			if err := tx.First(&product, p.ProductID).Error; err != nil {
				return fmt.Errorf("produk dengan ID %d tidak ditemukan", p.ProductID)
			}

			log := model.LogProduk{
				Nama_produk:    product.Nama_produk,
				Slug:           product.Slug,
				Harga_reseller: product.Harga_reseller,
				Harga_konsumen: product.Harga_konsumen,
				Deskripsi:      product.Deskripsi,
				Id_toko:        product.Id_toko,
				Id_category:    product.Id_category,
				Created_at:     time.Now(),
			}

			if err := tx.Create(&log).Error; err != nil {
				return errors.New("gagal membuat log produk")
			}

			hargaInt, _ := strconv.Atoi(product.Harga_konsumen)
			subtotal := hargaInt * p.Kuantitas
			totalHarga += subtotal

			detail := model.DetailTransaction{
				Id_trx:        transaction.Id,
				Id_log_produk: log.Id,
				Id_toko:       product.Id_toko,
				Kuantitas:     p.Kuantitas,
				Harga_total:   subtotal,
				Created_at:    time.Now(),
			}

			if err := tx.Create(&detail).Error; err != nil {
				return err
			}
		}

		transaction.Harga_total = totalHarga
		transaction.Kode_invoice = fmt.Sprintf("INV-%s-%03d", time.Now().Format("20060102"), transaction.Id)
		transaction.Updated_at = time.Now()

		return tx.Save(&transaction).Error
	})

	if err != nil {
		return 0, err
	}

	return transaction.Id, nil
}

func GetTransactionProductByIdRepo(db *gorm.DB, id, userId int) (*model.Transaction, error) {
	var trx model.Transaction

	err := db.Preload("Alamat").
		Preload("Detail.LogProduk").
		Where("id = ? AND id_user = ?", id, userId).
		First(&trx).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return &trx, err
}

func GetAllTransactionsRepo(db *gorm.DB, userID int) ([]model.Transaction, error) {
	var transactions []model.Transaction

	err := db.
		Preload("Alamat").
		Preload("Detail.LogProduk").
		Where("transactions.id_user = ?", userID).
		Order("transactions.id DESC").
		Find(&transactions).Error

	if err != nil {
		return nil, err
	}
	return transactions, nil
}
