package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/wambugucoder/simple-go-service/configs"
	"github.com/wambugucoder/simple-go-service/controller"
	"log"
)

func main() {

	app := fiber.New(fiber.Config{
		Prefork:      true,
		ServerHeader: "fiber",
	})

	app.Use(cors.New())

	//app.Use(compress.Config{Level: compress.LevelBestSpeed})

	//DATABASE
	configs.Connect()

	//ROUTES
	controller.SetupControllers(app)

	log.Fatal(app.Listen(":3000"))

}
