package store

import (
	"fmt"
	database "store-api/database/connection"
	"store-api/helper"
	"store-api/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetMyToko(ctx *fiber.Ctx) error {

	var id = ctx.Locals("user_id").(int)

	toko, err := repository.GetMyTokoRepo(database.Db, id)

	if err != nil {
		helper.Fail(ctx, "User Not Found", err.Error(), nil, 401)
	}

	if toko == nil {

		helper.Fail(ctx, "Store Not Found For This User", err.Error(), nil, 401)
	}

	return helper.Success(ctx, "Success Get User Data", toko)
}

func UpdateProfileToko(ctx *fiber.Ctx) error {

	id, _ := strconv.Atoi(ctx.Params("id"))

	namaToko := ctx.FormValue("nama_toko")

	file, err := ctx.FormFile("url_foto")
	var filePath string

	if err == nil && file != nil {
		filePath = fmt.Sprintf("./uploads/%s", file.Filename)

		if err := ctx.SaveFile(file, filePath); err != nil {
			return helper.Fail(ctx, "Failed To Save data", err.Error(), nil, 500)
		}
	}

	updateData := map[string]any{
		"nama_toko": namaToko,
	}
	if filePath != "" {
		updateData["url_foto"] = filePath
	}

	if err := repository.UpdateStoreByIDRepo(database.Db, id, updateData); err != nil {

		return helper.Fail(ctx, "Failed To Update data", err.Error(), nil, 500)
	}

	return helper.Success(ctx, "Succeed to UPDATE data", "Update Toko Succeed")
}

func GetTokoByIDRepo(ctx *fiber.Ctx) error {

	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return helper.Fail(ctx, "Invalid ID", []string{"ID Tidak Valid"}, nil, 404)
	}

	toko, err := repository.GetTokoByIDRepo(database.Db, id)

	if err != nil {
		return helper.Fail(ctx, "Failed To Get Data", []string{err.Error()}, nil, 500)
	}

	if toko == nil {
		return helper.Fail(ctx, "Failed To Get Data", []string{"Toko Tidak Ditemukan"}, nil, 404)
	}

	return helper.Success(ctx, "Succeed To Get Data", toko)

}

func GetAllToko(ctx *fiber.Ctx) error {
	limit, _ := strconv.Atoi(ctx.Query("limit", "10"))
	page, _ := strconv.Atoi(ctx.Query("page", "1"))

	name := ctx.Query("name", "")

	tokoList, _, err := repository.GetAllTokoRepo(database.Db, limit, page, name)

	if err != nil {
		return helper.Fail(ctx, "Failed To Get Data", err.Error(), nil, 500)
	}

	return helper.Success(ctx, "Succeed Get Data", fiber.Map{
		"Page":  page,
		"Limit": limit,
		"data":  tokoList,
	})

}
