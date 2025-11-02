package server

import (
	"store-api/router"

	"github.com/gofiber/fiber/v2"
)

func StartServer() *fiber.App {
	routes := fiber.New()

	routes.Static("/uploads", "./uploads")

	router.AuthRouter(routes)
	router.TokoRoutes(routes)
	router.AddressRoutes(routes)
	router.UserRoutes(routes)
	router.CategoryRoutes(routes)
	router.ProductRoutes(routes)
	router.TrxRouter(routes)

	return routes
}
