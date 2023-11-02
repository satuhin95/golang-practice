package main

import (
	"log"
	"react-go-fiber-jwt/database"
	"react-go-fiber-jwt/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)
func init(){
	database.ConnectDB()
}
func main() {
    app := fiber.New()

	// app.Use(cors.New(cors.Config{AllowCredentials: true}))
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowOrigins: "http://localhost:3000", // Set the specific origin here
	}))

	routes.SetupRoutes(app)

 

    log.Fatal(app.Listen(":8000"))
}