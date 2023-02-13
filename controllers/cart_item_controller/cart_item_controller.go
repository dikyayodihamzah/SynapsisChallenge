package cart_item_controller

import (
	"net/http"

	"github.com/dikyayodihamzah/SynapsisChallenge/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Index(c *fiber.Ctx) error {
	var cart_item []models.CartItem
	models.DB.Find(&cart_item)

	return c.Status(fiber.StatusOK).JSON(cart_item)
}

func Show(c *fiber.Ctx) error {

	id := c.Params("id")
	var cart_item models.CartItem
	
	if err := models.DB.First(&cart_item, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Item not found",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(cart_item)
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
	
	id := c.Params("id")
	var cart_item models.CartItem

	if err := c.BodyParser(&cart_item); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}
	
	if models.DB.First(&cart_item, id).Error; err != nil {
		
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Item tidak ditemukan",
			})
		}

		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		}) 
	}
	
	
}

func Delete(c *fiber.Ctx) error {
	return nil
}