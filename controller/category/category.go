package category

import (
	"fmt"
	database "store-api/database/connection"
	"store-api/helper"
	"store-api/model"
	"store-api/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetAllCategory(ctx *fiber.Ctx) error {

	data, err := repository.GetAllCategoryRepo(database.Db)

	if err != nil {
		return helper.Fail(ctx, "Failed To Get Categories", err.Error(), nil, 500)
	}

	if len(data) == 0 {
		return helper.Success(ctx, "No Categories Found", nil)
	}

	return helper.Success(ctx, "Succeed To  Get Data ", data)

}

func GetCategoryById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return helper.Fail(ctx, "Id Not Valid", err.Error(), nil, 401)
	}

	cat, _ := repository.GetCategoryByIdRepo(database.Db, id)

	if cat == nil {
		return helper.Fail(ctx, "Failed To Get Data", []string{"No Data Category"}, nil, 404)
	}

	return helper.Success(ctx, "Success To Get Data", cat)

}

func CreateCategory(ctx *fiber.Ctx) error {

	var category model.Category

	if err := ctx.BodyParser(&category); err != nil {
		return helper.Fail(ctx, "Invalid Request", "Invalid Json", nil, 400)
	}

	id, err := repository.CreateCategoryRepo(database.Db, &category)

	if err != nil {

		return helper.Fail(ctx, "Database error when checking user ", err.Error(), nil, 500)
	}

	return helper.Success(ctx, "Succeed To Post Data", id)
}
func UpdateCategoryById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.Fail(ctx, "Id Not Valid", err.Error(), nil, 401)
	}

	var cat model.Category
	if err := ctx.BodyParser(&cat); err != nil {
		return helper.Fail(ctx, "Invalid Request", "Invalid Json", nil, 400)
	}

	updatedData := map[string]any{
		"nama": cat.Nama,
	}

	if err := repository.UpdateCategoryByIdRepo(database.Db, id, updatedData); err != nil {

		if err.Error() == fmt.Sprintf("no category found with id %d", id) {
			return helper.Fail(ctx, "Failed To Update Data", err.Error(), nil, 404)
		}

		return helper.Fail(ctx, "Failed To Update Data", err.Error(), nil, 500)
	}

	return helper.Success(ctx, "Succeed To Update Data", nil)
}

func DeleteProductByID(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return helper.Fail(ctx, "Id Not Valid", err.Error(), nil, 401)
	}

	res, _ := repository.DeleteCategoryByIDRepo(database.Db, id)

	if res == 0 {
		return helper.Fail(ctx, "Failed To Delete Data", []string{"Record Not Found"}, nil, 400)
	}

	return helper.Success(ctx, "Succeed To Delete Data", nil)
}
