package controller

import (
	"fiber-project/database"
	"fiber-project/model"
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func GetClientInfos(c *fiber.Ctx) error {
	db := database.DB
	var clientinfo []model.Clientinfo
	time.Sleep(time.Millisecond * 1000)
	// find all clientinfo in the database
	db.Find(&clientinfo)

	// If no clientinfo is present return an error
	if len(clientinfo) == 0 {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No clientinfo present", "data": nil})
	}

	// Else return clientinfo
	return c.JSON(fiber.Map{"status": "success", "message": "clientinfo Found", "data": clientinfo})
}

func CreateClientInfo(c *fiber.Ctx) error {
	db := database.DB
	clientinfo := new(model.Clientinfo)

	// Store the body in the clientinfo and return error if encountered
	err := c.BodyParser(clientinfo)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Review your input", "data": err})
	}
	// Add a uuid to the clientinfo
	clientinfo.ID = uuid.New()
	// Create the clientinfo and return error if encountered
	err = db.Create(&clientinfo).Error
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"status": "error", "message": "Could not create clientinfo", "data": err})
	}

	// Return the created clientinfo
	return c.JSON(fiber.Map{"status": "success", "message": "Created Note", "data": clientinfo})
}
func GetClientInfo(c *fiber.Ctx) error {
	db := database.DB
	var clientinfo model.Clientinfo
	time.Sleep(time.Millisecond * 1000)
	// Read the param clientinfoid
	id := c.Params("cinfoId")

	// Find the clientinfo with the given Id
	db.Find(&clientinfo, "id = ?", id)

	// If no such clientinfo present return an error
	if clientinfo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No clientinfo present", "data": nil})
	}

	// Return the clientinfo with the Id
	return c.JSON(fiber.Map{"status": "success", "message": "Clientinfo Found", "data": clientinfo})
}

func DeleteClientInfo(c *fiber.Ctx) error {
	db := database.DB
	var clientinfo model.Clientinfo

	// Read the param cinfoId
	id := c.Params("cinfoId")

	// Find the note with the given Id
	db.Find(&clientinfo, "id = ?", id)

	// If no such note present return an error
	if clientinfo.ID == uuid.Nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "No clientinfo present", "data": nil})
	}

	// Delete the clientinfo and return error if encountered
	err := db.Delete(&clientinfo, "id = ?", id).Error

	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status": "error", "message": "Failed to delete clientinfo", "data": nil})
	}

	// Return success message
	return c.JSON(fiber.Map{"status": "success", "message": "Deleted clientinfo"})
}