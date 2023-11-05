package router

import (
	clientinfoRouter "fiber-project/router/clientinfo"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupRouters(app *fiber.App){
	api := app.Group("/api", logger.New())

	clientinfoRouter.SetupClientinfoRouters(api)
}