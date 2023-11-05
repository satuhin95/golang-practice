package clientinfoRouter

import (
	"fiber-project/controller"

	"github.com/gofiber/fiber/v2"
)

func SetupClientinfoRouters(router fiber.Router){
	clientinfo := router.Group("clientinfo")


	clientinfo.Post("/", controller.CreateClientInfo)
	clientinfo.Get("/", controller.GetClientInfos)
	clientinfo.Get("/:cinfoId", controller.GetClientInfo)
	clientinfo.Delete("/:cinfoId", controller.DeleteClientInfo)

}