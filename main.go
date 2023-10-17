package main

import (
	"go-fiber-jwt/initial/database"
	"go-fiber-jwt/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	database.DBconn()
	routes.Setup(app)

	app.Listen(":3000")
}
