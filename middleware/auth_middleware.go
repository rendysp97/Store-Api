package middleware

import (
	"os"
	"store-api/helper"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		tokenHeader := ctx.Get("Authorization")
		if len(tokenHeader) < 7 || tokenHeader[:7] != "Bearer " {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorization session"})

		}
		tokenHeader = tokenHeader[7:]

		claim := &helper.Claims{}
		token, err := jwt.ParseWithClaims(tokenHeader, claim, func(t *jwt.Token) (any, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return ctx.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid Token"})

		}

		ctx.Locals("notelp", claim.No_telp)
		ctx.Locals("is_admin", claim.Is_admin)
		ctx.Locals("user_id", claim.User_id)

		return ctx.Next()
	}
}
