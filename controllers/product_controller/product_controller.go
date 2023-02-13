package product_controller

import (
	"net/http"

	"github.com/dikyayodihamzah/SynapsisChallenge/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {

	var products []models.Product
	models.DB.Find(&products)

	return c.Status(fiber.StatusOK).JSON(products)
}

func Show(c *fiber.Ctx) error {
	
	id := c.Params("id")
	var product models.Product
	
	if err := models.DB.First(&product, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Product not found",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(product)
}