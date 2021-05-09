package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/session"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/shareed2k/goth_fiber"
	"github.com/wambugucoder/simple-go-service/controller"
	"github.com/wambugucoder/simple-go-service/services"
	"log"
	"time"
)

func main() {

	app := fiber.New(fiber.Config{
		Prefork:      false,
		ServerHeader: "fiber",
	})

	app.Use(cors.New())

	//USE THIS FOR OAUTH
	services.GoogleAuth()
	// optional config
	config := session.Config{
		Expiration:   30 * time.Minute, // default: 2 * time.Hour
		CookieSecure: true,             // default: false
	}

	// create session handler
	sessions := session.New(config)

	goth_fiber.SessionStore = sessions

	//app.Use(compress.Config{Level: compress.LevelBestSpeed})

	//DATABASE
	//configs.Connect()

	//ROUTES
	controller.SetupControllers(app)

	log.Fatal(app.Listen(":3000"))

}
