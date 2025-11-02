package router

import (
	"store-api/controller/auth"

	"github.com/gofiber/fiber/v2"
)

func AuthRouter(router fiber.Router) {

	router.Post("/auth/register", auth.Register)

	router.Post("/auth/login", auth.Login)

}
