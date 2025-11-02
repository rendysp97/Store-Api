package router

import (
	"store-api/controller/product"
	"store-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func ProductRoutes(routes fiber.Router) {

	prod := routes.Group("/product", middleware.AuthMiddleware())

	prod.Post("/", product.CreateProduct)
	prod.Get("/:id", product.GetProductById)
	prod.Get("/", product.GetAllProduct)
	prod.Put("/:id", product.UpdateProductById)
	prod.Delete("/:id", product.DeleteProductById)

}
