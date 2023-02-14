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
	
	if models.DB.Where("id = ?", id).Updates(&cart_item).RowsAffected == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Can not update data",
		})
	}
	
	return c.JSON(fiber.Map{
		"message": "Data succesfully updated",
	})
}

func Delete(c *fiber.Ctx) error {

	id := c.Params("id")
	var cart_item models.CartItem

	if models.DB.Delete(&cart_item, id).RowsAffected == 0 {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Data not found",
		})
	}

	return c.JSON(fiber.Map{
		"message": "Data succesfully deleted",
	})

}