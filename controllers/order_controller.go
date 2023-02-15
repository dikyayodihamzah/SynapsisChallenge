package controllers

import (
	"net/http"

	"github.com/dikyayodihamzah/SynapsisChallenge/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func Checkout(c *fiber.Ctx) error {
	return nil
}

func Init(c *fiber.Ctx) error {
	order := models.Order{
		CustomerID: 1,
		
	}

	if err := c.BodyParser(&order); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

}