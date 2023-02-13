package main

import (
	"github.com/dikyayodihamzah/SynapsisChallenge/controllers/cart_item_controller"
	"github.com/dikyayodihamzah/SynapsisChallenge/controllers/product_controller"
	"github.com/dikyayodihamzah/SynapsisChallenge/models"
	"github.com/gofiber/fiber/v2"
)

func main() {
	models.ConnectDatabase()

	app := fiber.New()

	api := app.Group("/api")
	product := api.Group("/products")
	cart_item := api.Group("/cart-items")

	//products route
	product.Get("/", product_controller.Index)

	//cart items route
	cart_item.Get("/", cart_item_controller.FindAll)
	cart_item.Get("/:id", cart_item_controller.Show)
	cart_item.Post("/", cart_item_controller.Create)
	cart_item.Put("/:id", cart_item_controller.Update)
	cart_item.Delete("/:id", cart_item_controller.Delete)


	app.Listen(":3000")
}