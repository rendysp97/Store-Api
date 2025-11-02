package router

import (
	"store-api/controller/transaction"
	"store-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func TrxRouter(routes fiber.Router) {

	trx := routes.Group("/trx", middleware.AuthMiddleware())

	trx.Post("/", transaction.CreateTransaction)
	trx.Get("/:id", transaction.GetTransactionProductById)
	trx.Get("/", transaction.GetAllTransactions)

}
