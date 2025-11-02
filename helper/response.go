package helper

import "github.com/gofiber/fiber/v2"

type Response struct {
	Status  bool
	Message string
	Errors  any
	Data    any
}

func Success(ctx *fiber.Ctx, message string, data any) error {
	return ctx.Status(fiber.StatusOK).JSON(
		Response{
			Status:  true,
			Message: message,
			Errors:  nil,
			Data:    data,
		},
	)
}

func Fail(ctx *fiber.Ctx, message string, errors, data any, code int) error {
	return ctx.Status(code).JSON(Response{
		Status:  false,
		Message: message,
		Errors:  errors,
		Data:    data,
	})
}
