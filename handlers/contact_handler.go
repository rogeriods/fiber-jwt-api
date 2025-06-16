package handlers

import (
	"rogeriods/fiber-jwt-api/configs"
	"rogeriods/fiber-jwt-api/models"

	"github.com/gofiber/fiber/v2"
)

// Create new contact
func CreateContact(c *fiber.Ctx) error {
	userID := c.Locals("userID")

	contact := new(models.Contact)
	if err := c.BodyParser(contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	contact.UserID = uint(userID.(float64)) // convert any to uint
	if err := configs.DB.Create(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not create contact"})
	}

	return c.JSON(contact)
}

// Get contact by user logged
func GetContacts(c *fiber.Ctx) error {
	userID := c.Locals("userID")

	var contacts []models.Contact
	if err := configs.DB.Where("user_id = ?", userID).Find(&contacts).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Could not fetch contacts"})
	}

	return c.JSON(contacts)
}

// Get contact by user logged
func GetContactById(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	id := c.Params("id")

	var contact models.Contact
	if err := configs.DB.Where("id = ? AND user_id = ?", id, userID).First(&contact).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Contact not found"})
	}

	return c.JSON(contact)
}

// Update contact
func UpdateContact(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	id := c.Params("id")

	var contact models.Contact
	if err := configs.DB.Where("id = ? AND user_id = ?", id, userID).First(&contact).Error; err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Contact not found"})
	}

	if err := c.BodyParser(&contact); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	configs.DB.Save(&contact)
	return c.JSON(contact)
}

// Delete contact by id and user_id
func DeleteContact(c *fiber.Ctx) error {
	userID := c.Locals("userID")
	id := c.Params("id")

	if err := configs.DB.Where("id = ? AND user_id = ?", id, userID).Delete(&models.Contact{}).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Delete failed"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
