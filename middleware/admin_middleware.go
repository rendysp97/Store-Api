package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func AdminRoute() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		admin, exist := ctx.Locals("is_admin").(bool)

		if !exist || !admin {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "(Admin Only) Login with admin account"})
		}

		return ctx.Next()

	}

}
