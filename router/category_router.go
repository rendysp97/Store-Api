package router

import (
	"store-api/controller/category"
	"store-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func CategoryRoutes(routes fiber.Router) {

	cat := routes.Group("/category", middleware.AuthMiddleware(), middleware.AdminRoute())

	cat.Post("/", category.CreateCategory)
	cat.Get("/", category.GetAllCategory)
	cat.Get("/:id", category.GetCategoryById)
	cat.Put("/:id", category.UpdateCategoryById)
	cat.Delete("/:id", category.DeleteProductByID)

}
