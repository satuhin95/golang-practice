package routes

import (
	"react-go-fiber-jwt/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App){
	api := app.Group("api")
	api.Get("/", controller.Hello)
	api.Post("/register", controller.Register)
	api.Post("/login", controller.Login)
	api.Get("/user", controller.User)
	api.Get("/logout", controller.Logout)
}