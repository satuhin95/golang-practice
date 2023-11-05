package main

import (
	"fiber-project/database"
	"fiber-project/router"

	"github.com/gofiber/fiber/v2"
	  "github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())
	database.ConnectDB()

	router.SetupRouters(app)

	app.Get("/", func(c *fiber.Ctx) error {
		err := c.SendString("And the API is UP!")
		return err
	})


	app.Listen(":8000")
}