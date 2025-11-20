package rest

import "github.com/gofiber/fiber/v2"


func ErrorResponse(ctx *fiber.Ctx, statusCode int, err error) error {
	return ctx.Status(statusCode).JSON(err.Error())
}


func OkResponse(ctx *fiber.Ctx, msg string, data any) error {
	return ctx.Status(200).JSON(&fiber.Map{
		"message": msg,
		"data": data,
	})
}