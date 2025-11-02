package address

import (
	database "store-api/database/connection"
	"store-api/helper"
	"store-api/model"
	"store-api/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetMyAlamat(ctx *fiber.Ctx) error {
	id := ctx.Locals("user_id").(int)

	address, err := repository.GetAddressRepo(database.Db, id)

	if err != nil {
		return helper.Fail(ctx, "Failed To Get Address", err.Error(), nil, 500)
	}

	if len(address) == 0 {
		return helper.Success(ctx, "No Address Found", []model.Alamat{})
	}

	return helper.Success(ctx, "Success Get User Data", address)
}

func GetAlamatById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	if err != nil {
		return helper.Fail(ctx, "Id Not Valid", err.Error(), nil, 401)
	}

	userId, _ := ctx.Locals("user_id").(int)

	address, _ := repository.GetAlamatByIdRepo(database.Db, id, userId)

	if address == nil {
		return helper.Fail(ctx, "Unauthorizate ", "Delete Data Or Unauthorizated Session", nil, 403)
	}

	return helper.Success(ctx, "Success To Get Data", address)

}

func CreateAlamat(ctx *fiber.Ctx) error {

	var id = ctx.Locals("user_id").(int)

	var address model.Alamat

	if err := ctx.BodyParser(&address); err != nil {
		return helper.Fail(ctx, "Invalid Request", "Invalid Json", nil, 400)
	}

	address.Id_user = id

	if address.Judul == "" || address.Nama_penerima == "" || address.No_telp == "" || address.Detail_Alamat == "" {
		return helper.Fail(ctx, "Please Input Valid Address", "Invalid Json", nil, 400)
	}

	id, err := repository.CreateAddressRepo(database.Db, &address)

	if err != nil {

		return helper.Fail(ctx, "Database error when checking user ", err.Error(), nil, 500)
	}

	return helper.Success(ctx, "Succeed To Post Data", id)
}

func UpdateAlamat(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))

	var userId = ctx.Locals("user_id").(int)

	var address model.Alamat

	if err != nil {
		return helper.Fail(ctx, "Id Not Valid", err.Error(), nil, 401)
	}

	if err := ctx.BodyParser(&address); err != nil {
		return helper.Fail(ctx, "Invalid Request", "Invalid Json", nil, 400)
	}

	updateData := map[string]any{
		"judul":         address.Judul,
		"nama_penerima": address.Nama_penerima,
		"no_telp":       address.No_telp,
		"detail_alamat": address.Detail_Alamat,
	}

	res := repository.UpdateAddressRepo(database.Db, id, userId, updateData)

	if res == 0 {
		return helper.Fail(ctx, "Failed To Get Data", []string{"Record Not Found"}, nil, 400)
	}

	return helper.Success(ctx, "Succeed to UPDATE data", "Update Alamat Succeed")
}

func DeleteAlamat(ctx *fiber.Ctx) error {
	id, _ := strconv.Atoi(ctx.Params("id"))

	userId := ctx.Locals("user_id").(int)

	res := repository.DeleteAlamatByIdRepo(database.Db, id, userId)

	if res == 0 {
		return helper.Fail(ctx, "Failed To Delete Data", []string{"Record Not Found"}, nil, 400)
	}

	return helper.Success(ctx, "Succeed to Delete ", "Delete Data Succeed")

}
