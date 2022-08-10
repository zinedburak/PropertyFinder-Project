package handler

import (
	"PropertyFinder/handler/models"
	"PropertyFinder/ports"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Handler struct {
	repository ports.DbPort
	core       ports.CorePort
	service    ports.ServicePort
}

func NewAdapter(repository ports.DbPort, core ports.CorePort, service ports.ServicePort) *Handler {
	return &Handler{repository: repository, core: core, service: service}
}

func (h Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/api/show-basket/:customerId", h.ShowBasket)
	app.Get("/api/list-products", h.ListProducts)
	app.Post("/api/add-to-basket", h.AddToBasket)
	app.Delete("/api/delete-basket-item", h.DeleteBasketItem)
	app.Post("/api/complete-order/:customerId", h.CompleteOrder)
}

func (h Handler) AddToBasket(c *fiber.Ctx) error {
	var request models.AddToBasketRequest
	err := c.BodyParser(&request)

	if err != nil || request.ProductId == 0 || request.CustomerId == 0 {
		if err != nil {
			log.Printf("error while parsing the request : %v", err)
		}
		c.Status(fiber.StatusBadRequest)
		return fiber.ErrBadRequest
	}

	response, err := h.service.AddToBasket(request.CustomerId, request.ProductId)

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return fiber.ErrInternalServerError
	}
	return c.JSON(response)
}

func (h Handler) ShowBasket(c *fiber.Ctx) error {
	customerId, err := c.ParamsInt("customerId")
	if err != nil {
		return err
	}
	response, err := h.service.ShowBasket(customerId)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return fiber.ErrInternalServerError
	}
	return c.JSON(response)
}

func (h Handler) DeleteBasketItem(c *fiber.Ctx) error {
	var request models.DeleteBasketRequest
	err := c.BodyParser(&request)
	if err != nil || request.ProductId == 0 || request.CustomerId == 0 {
		if err != nil {
			log.Printf("error while parsing the request : %v", err)
		}
		c.Status(fiber.StatusBadRequest)
		return fiber.ErrBadRequest
	}
	response, err := h.service.DeleteBasketItem(request.CustomerId, request.ProductId)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return fiber.ErrInternalServerError
	}
	return c.JSON(response)
}

func (h Handler) CompleteOrder(c *fiber.Ctx) error {
	customerId, err := c.ParamsInt("customerId")
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return fiber.ErrBadRequest
	}
	response, err := h.service.CompleteOrder(customerId)
	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h Handler) ListProducts(c *fiber.Ctx) error {
	response, err := h.service.ListProducts()
	if err != nil {
		return fiber.ErrInternalServerError
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
