package product

import (
	"fmt"
	database "store-api/database/connection"
	"store-api/helper"
	"store-api/model"
	"store-api/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateProduct(ctx *fiber.Ctx) error {

	userId := ctx.Locals("user_id").(int)

	toko, err := repository.GetTokoByUserIdRepo(database.Db, userId)

	if err != nil {
		return helper.Fail(ctx, "Fail Get Store", err.Error(), nil, 500)
	}

	if toko == nil {
		return helper.Fail(ctx, "Store Not Found", "User doesn't have a store", nil, 400)
	}

	nama_produk := ctx.FormValue("nama_produk")
	categoryId, _ := strconv.Atoi(ctx.FormValue("category_id"))
	harga_reseller := ctx.FormValue("harga_reseller")
	harga_konsumen := ctx.FormValue("harga_konsumen")
	stok, _ := strconv.Atoi(ctx.FormValue("stok"))
	deskripsi := ctx.FormValue("deskripsi")

	products := model.Product{
		Nama_produk:    nama_produk,
		Id_category:    categoryId,
		Harga_reseller: harga_reseller,
		Harga_konsumen: harga_konsumen,
		Stok:           stok,
		Deskripsi:      deskripsi,
		Id_toko:        toko.Id,
	}

	if err := repository.CreateProductRepository(database.Db, &products); err != nil {
		return helper.Fail(ctx, "Fail Create Product", err.Error(), nil, 500)
	}

	form, err := ctx.MultipartForm()

	if err != nil {
		return helper.Fail(ctx, "Invalid Form", err.Error(), nil, 400)
	}

	files := form.File["photos"]

	for _, file := range files {
		filePath := fmt.Sprintf("uploads/%s", file.Filename)
		ctx.SaveFile(file, filePath)

		foto := model.FotoProduk{
			Id_produk: products.Id,
			Url:       filePath,
		}

		if err := repository.CreateProductRepository(database.Db, &foto); err != nil {
			return helper.Fail(ctx, "Fail Save Foto", err.Error(), nil, 500)
		}

	}

	return helper.Success(ctx, "Succeed Post Data ", products.Id)

}

func GetProductById(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int)

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return helper.Fail(ctx, "Id Not Valid", err.Error(), nil, 401)
	}

	product, err := repository.GetProductById(database.Db, id, userId)

	if err != nil {
		return helper.Fail(ctx, "Failed to Get Product", err.Error(), nil, 500)
	}

	if product == nil {
		return helper.Fail(ctx, "Failed To Get Data", "No Data Product", nil, 404)
	}

	return helper.Success(ctx, "Succeed Get Data", product)

}
func GetAllProduct(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int)

	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	page, _ := strconv.Atoi(ctx.Query("page", "1"))
	offset := (page - 1) * limit

	filters := map[string]string{
		"nama_produk": ctx.Query("nama_produk"),
		"category_id": ctx.Query("category_id"),
		"min_harga":   ctx.Query("min_harga"),
		"max_harga":   ctx.Query("max_harga"),
	}

	products, err := repository.GetAllProductRepo(database.Db, userId, limit, offset, filters)
	if err != nil {
		return helper.Fail(ctx, "Failed to GET data", err.Error(), nil, 500)
	}

	response := fiber.Map{
		"data":  products,
		"page":  page,
		"limit": limit,
	}

	return helper.Success(ctx, "Succeed to GET data", response)
}

func UpdateProductById(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	userId := ctx.Locals("user_id").(int)

	nama_produk := ctx.FormValue("nama_produk")
	categoryId, _ := strconv.Atoi(ctx.FormValue("category_id"))
	harga_reseller := ctx.FormValue("harga_reseller")
	harga_konsumen := ctx.FormValue("harga_konsumen")
	stok, _ := strconv.Atoi(ctx.FormValue("stok"))
	deskripsi := ctx.FormValue("deskripsi")

	updateData := map[string]any{}

	if nama_produk != "" {
		updateData["nama_produk"] = nama_produk
	}
	if categoryId != 0 {
		updateData["id_category"] = categoryId
	}
	if harga_reseller != "" {
		updateData["harga_reseller"] = harga_reseller
	}
	if harga_konsumen != "" {
		updateData["harga_konsumen"] = harga_konsumen
	}
	if stok != 0 {
		updateData["stok"] = stok
	}
	if deskripsi != "" {
		updateData["deskripsi"] = deskripsi
	}

	if err := repository.UpdateProductRepo(database.Db, id, userId, updateData); err != nil {
		return helper.Fail(ctx, "Failed Save Photo", err.Error(), nil, 500)
	}

	form, err := ctx.MultipartForm()
	if err == nil && form != nil {
		files := form.File["photos"]
		for _, file := range files {
			filePath := fmt.Sprintf("uploads/%s", file.Filename)

			if err := ctx.SaveFile(file, filePath); err != nil {
				return helper.Fail(ctx, "Failed to save file", err.Error(), nil, 500)
			}

			foto := model.FotoProduk{
				Id_produk: id,
				Url:       filePath,
			}

			if err := repository.CreateProductRepository(database.Db, &foto); err != nil {
				return helper.Fail(ctx, "Failed Save Photo", err.Error(), nil, 500)
			}
		}
	}

	return helper.Success(ctx, "Succeed Update Data", nil)

}

func DeleteProductById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	userId := ctx.Locals("user_id").(int)

	if err != nil {
		return helper.Fail(ctx, "Id Not Valid", err.Error(), nil, 401)
	}

	if err := repository.DeleteProductByIDRepo(database.Db, id, userId); err != nil {
		return helper.Fail(ctx, "Failed To Delete Product", err.Error(), nil, 500)
	}

	return helper.Success(ctx, "Succeed To Delete Product", nil)

}
