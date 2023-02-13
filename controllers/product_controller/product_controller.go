package product_controller

import (
	"github.com/dikyayodihamzah/SynapsisChallenge/models"
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {

	var products []models.Product
	models.DB.Find(&products)

	return c.Status(fiber.StatusOK).JSON(products)
}

func Show(c *fiber.Ctx) error {
	id := c.Params("id")
	var product models.Product
	models.DB.First(&product, id)

	return c.Status(fiber.StatusOK).JSON(product)
}