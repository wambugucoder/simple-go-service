package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wambugucoder/simple-go-service/services"
)

func SetupControllers(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/add-user", services.AddUser)
	api.Get("/auth/:provider", services.BeginGoogleOauth)
	api.Get("/auth/:provider/callback", services.CompleteOauth)
	api.Get("/hello", func(ctx *fiber.Ctx) error {
		return ctx.SendString("hello")
	})

}
