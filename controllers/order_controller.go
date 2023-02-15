package controllers

import (
	"fmt"
	"os/user"
	"strconv"

	"github.com/dikyayodihamzah/SynapsisChallenge/models"
	"github.com/gofiber/fiber/v2"
)

type Order_req struct {
	ItemsID []uint `json:"items_id"`
}

func Checkout(c *fiber.Ctx) error {
	var cart_items []models.CartItem
	var order models.Order
	var order_req Order_req

	customer, err := user.Current()
	if err != nil {
		return c.JSON(err.Error())
	}

	customer_id, err := strconv.ParseUint(customer.Uid, 10, 32)
	if err != nil {
		return c.JSON(err.Error())
	}

	for _, item_id := range order_req.ItemsID {
		id := fmt.Sprintf("%d", item_id)

		cart_item := models.DB.Preload("CartItems.Product").Preload("User").First(&order, "id = ? AND customer_id = ?", id, customer_id)

		if models.CartItem.OrderID != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		order.Total += cart_item.Product.Price * cart_item.Quantity
		cart_items = append(cart_items, cart_item)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Checkout success",
	})
}
