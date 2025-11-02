package router

import (
	"store-api/controller/store"
	"store-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func TokoRoutes(routes fiber.Router) {

	toko := routes.Group("/toko", middleware.AuthMiddleware())

	toko.Get("/my", store.GetMyToko)
	toko.Put("/:id", store.UpdateProfileToko)
	toko.Get("/:id", store.GetTokoByIDRepo)
	toko.Get("/", store.GetAllToko)

}
