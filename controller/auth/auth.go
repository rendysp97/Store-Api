package auth

import (
	database "store-api/database/connection"
	"store-api/helper"
	"store-api/model"
	"store-api/repository"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Login(ctx *fiber.Ctx) error {

	var data model.User

	if err := ctx.BodyParser(&data); err != nil {

		return helper.Fail(ctx, "No Valid", "Invalid Json", nil, 400)

	}

	if data.Notelp == "" || data.Kata_sandi == "" {

		return helper.Fail(ctx, "No Valid", "Please Input Valid Phone And Password", nil, 401)

	}

	user, err := repository.GetUserByPhone(database.Db, data.Notelp)

	if err != nil {

		return helper.Fail(ctx, "Database error when checking user ", err.Error(), nil, 500)
	}

	if user == nil || bcrypt.CompareHashAndPassword([]byte(user.Kata_sandi), []byte(data.Kata_sandi)) != nil {
		return helper.Fail(ctx, "Failed To Post Data", []string{"Phone or Password Incorrect !"}, nil, 401)
	}

	token, err := helper.GenerateJwt(user.Notelp, user.Is_admin, user.Id)

	if err != nil {

		return helper.Fail(ctx, "No Valid", "Failed Generated Token", nil, 500)
	}

	var province, city any

	if user.Id_provinsi != nil {
		province, _ = helper.GetProvinceById(*user.Id_provinsi)
	}

	if user.Id_kota != nil {
		city, _ = helper.GetCityById(*user.Id_provinsi, *user.Id_kota)
	}

	resp := fiber.Map{
		"nama":          user.Nama,
		"no_telp":       user.Notelp,
		"tanggal_lahir": user.Tanggal_lahir,
		"tentang":       user.Nama,
		"pekerjaan":     user.Pekerjaan,
		"email":         user.Email,
		"id_provinsi":   province,
		"id_kota":       city,
		"token":         token,
	}

	return helper.Success(ctx, "Success To Post Data", resp)
}

func Register(ctx *fiber.Ctx) error {

	var data model.User

	if err := ctx.BodyParser(&data); err != nil {
		return helper.Fail(ctx, "Invalid Request", "Invalid Json", nil, 400)

	}

	if data.Notelp == "" || data.Kata_sandi == "" {
		return helper.Fail(ctx, "No Valid", "Please Input Valid Phone And Password", nil, 401)

	}

	user, err := repository.GetUserByPhoneAndEmail(database.Db, data.Notelp, data.Email)

	if err != nil {
		return helper.Fail(ctx, "Database error when checking user ", err.Error(), nil, 500)
	}

	if user != nil {

		if data.Notelp == user.Notelp && data.Email == user.Email {

			return helper.Fail(ctx, "Phone and Email Already Registered , Login Instead ", err, nil, 401)
		}

	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(data.Kata_sandi), bcrypt.DefaultCost)

	if err != nil {

		return helper.Fail(ctx, "Fail Hashing Password ", err, nil, 500)

	}

	data.Kata_sandi = string(hashed)

	if err := repository.CreateUser(database.Db, &data); err != nil {
		return helper.Fail(ctx, "Failed To Post Data ", err.Error(), nil, 500)
	}

	toko := model.Toko{
		Nama_toko: data.Nama + " Store",
		Id_user:   data.Id,
	}

	if err := repository.CreateStoreRepo(database.Db, &toko); err != nil {

		return helper.Fail(ctx, "Failed To Create Store ", err.Error(), nil, 500)

	}

	return helper.Success(
		ctx, "Success To Post Data", "Register Succedd")

}
