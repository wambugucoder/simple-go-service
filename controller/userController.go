package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/wambugucoder/simple-go-service/services"
)

func SetupControllers(app *fiber.App) {
	api := app.Group("/api/v1")
	api.Post("/add-user", services.AddUser)
}
