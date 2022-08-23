package ports

import "github.com/gofiber/fiber/v2"

type BasketHandlerPort interface {
	RegisterRoutes(app *fiber.App)
	ShowBasket(c *fiber.Ctx) error
	AddToBasket(c *fiber.Ctx) error
	DeleteBasketItem(c *fiber.Ctx) error
	CompleteOrder(c *fiber.Ctx) error
	ListProducts(c *fiber.Ctx) error
}
