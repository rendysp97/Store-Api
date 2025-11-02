package router

import (
	address "store-api/controller/User/Address"
	"store-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func AddressRoutes(routes fiber.Router) {

	Addres := routes.Group("/user", middleware.AuthMiddleware())

	Addres.Post("/alamat", address.CreateAlamat)
	Addres.Get("/alamat", address.GetMyAlamat)
	Addres.Get("/alamat/:id", address.GetAlamatById)
	Addres.Put("/alamat/:id", address.UpdateAlamat)
	Addres.Delete("/alamat/:id", address.DeleteAlamat)

}
