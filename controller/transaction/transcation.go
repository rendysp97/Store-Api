package transaction

import (
	database "store-api/database/connection"
	"store-api/helper"
	"store-api/model"
	"store-api/repository"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateTransaction(ctx *fiber.Ctx) error {
	var input model.CreateTransactionInput

	if err := ctx.BodyParser(&input); err != nil {
		return helper.Fail(ctx, "Invalid request body", err.Error(), nil, 400)
	}

	if input.AlamatKirim == 0 {
		return helper.Fail(ctx, "alamat_kirim tidak boleh kosong", nil, nil, 400)
	}

	if len(input.Products) == 0 {
		return helper.Fail(ctx, "detail_trx tidak boleh kosong", nil, nil, 400)
	}

	userID := ctx.Locals("user_id").(int)

	trxID, err := repository.CreateTransactionRepo(database.Db, userID, input)
	if err != nil {
		return helper.Fail(ctx, "Failed to create transaction", err.Error(), nil, 500)
	}

	return helper.Success(ctx, "Succeed to POST data", trxID)
}

func GetTransactionProductById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		return helper.Fail(ctx, "ID tidak valid", err.Error(), nil, 400)
	}

	userID := ctx.Locals("user_id").(int)

	trx, err := repository.GetTransactionProductByIdRepo(database.Db, id, userID)
	if err != nil {
		return helper.Fail(ctx, "Failed to GET data", err.Error(), nil, 500)
	}

	if trx == nil {
		return helper.Fail(ctx, "Failed to GET data", []string{"No Data Trx"}, nil, 404)
	}

	return helper.Success(ctx, "Succeed to GET data", trx)
}

func GetAllTransactions(ctx *fiber.Ctx) error {
	userID := ctx.Locals("user_id").(int)

	transactions, err := repository.GetAllTransactionsRepo(database.Db, userID)
	if err != nil {
		return helper.Fail(ctx, "Failed to GET data", err.Error(), nil, 500)
	}
	var dataList []model.Transaction
	for _, t := range transactions {
		dataList = append(dataList, model.Transaction{
			Id:                t.Id,
			Harga_total:       t.Harga_total,
			Kode_invoice:      t.Kode_invoice,
			Method_bayar:      t.Method_bayar,
			Alamat_pengiriman: t.Alamat_pengiriman,
			Detail:            t.Detail,
		})
	}

	return helper.Success(ctx, "Succeed to GET data", dataList)
}
