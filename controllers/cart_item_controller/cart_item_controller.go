package cart_item_controller

import (
	"github.com/dikyayodihamzah/SynapsisChallenge/models"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	var cart_item []models.CartItem
	models.DB.Find(&cart_item)

	return c.Status(fiber.StatusOK).JSON(cart_item)
}

func Show(c *fiber.Ctx) error {
	return nil
}

func Create(c *fiber.Ctx) error {

	var cart_item models.CartItem

	if err := c.BodyParser(&cart_item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	if err := models.DB.Create(&cart_item).Error; err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message" : err.Error(),
		})
	}

	return c.JSON(cart_item)
}

func Update(c *fiber.Ctx) error {
	return nil
}

func Delete(c *fiber.Ctx) error {
	return nil
}