package product_controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/dikyayodihamzah/SynapsisChallenge/models"
)

func Index(c *fiber.Ctx) error {
	var products []models.Product
	models.DB.Find(&products)

	return c.Status(fiber.StatusOK).JSON(products)
}