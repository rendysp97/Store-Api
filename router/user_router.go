package router

import (
	user "store-api/controller/User"
	"store-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(routes fiber.Router) {

	users := routes.Group("/user", middleware.AuthMiddleware())

	users.Get("/", user.GetMyProfile)
	users.Put("/", user.UpdateProfile)

}
