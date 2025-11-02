package user

import (
	"fmt"
	database "store-api/database/connection"
	"store-api/helper"
	"store-api/model"
	"store-api/repository"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func GetMyProfile(ctx *fiber.Ctx) error {
	id := ctx.Locals("user_id").(int)

	user, err := repository.GetMyProfileRepo(database.Db, id)

	if err != nil {
		return helper.Fail(ctx, "Failed To Get Address", err.Error(), nil, 500)
	}

	return helper.Success(ctx, "Success Get Profile", fiber.Map{
		"id":            user.Id,
		"nama":          user.Nama,
		"no_telp":       user.Notelp,
		"email":         user.Email,
		"tanggal_lahir": user.Tanggal_lahir,
		"pekerjaan":     user.Pekerjaan,
		"id_Provinsi":   user.Id_provinsi,
		"id_kota":       user.Id_kota,
	})

}

func UpdateProfile(ctx *fiber.Ctx) error {
	userId := ctx.Locals("user_id").(int) // ✅ Ambil dari token, bukan URL

	var user model.User

	if err := ctx.BodyParser(&user); err != nil {
		return helper.Fail(ctx, "Invalid Request", "Invalid JSON", nil, 400)
	}

	updateData := map[string]any{}

	if user.Id_provinsi != nil {
		updateData["id_provinsi"] = *user.Id_provinsi
	}
	if user.Id_kota != nil {
		updateData["id_kota"] = *user.Id_kota
	}
	if user.Nama != "" {
		updateData["nama"] = user.Nama
	}
	if user.Notelp != "" {
		updateData["notelp"] = user.Notelp
	}
	if user.Email != "" {
		updateData["email"] = user.Email
	}
	if user.Pekerjaan != nil {
		updateData["pekerjaan"] = user.Pekerjaan
	}
	if user.Tanggal_lahir != nil {
		updateData["tanggal_lahir"] = user.Tanggal_lahir
	}
	if user.Kata_sandi != "" {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Kata_sandi), bcrypt.DefaultCost)
		updateData["kata_sandi"] = string(hashed)
	}

	if len(updateData) == 0 {
		return helper.Fail(ctx, "No Data Provided", "Nothing to update", nil, 400)
	}

	if err := repository.ProfileUpdateRepo(database.Db, userId, updateData); err != nil {
		if err.Error() == fmt.Sprintf("no user found with id %d", userId) {
			return helper.Fail(ctx, "User Not Found", err.Error(), nil, 404)
		}
		return helper.Fail(ctx, "Failed To Update Data", err.Error(), nil, 500)
	}

	return helper.Success(ctx, "Succeed To Update Data", "Update Profile Success")
}
