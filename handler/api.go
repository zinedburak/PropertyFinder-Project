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
}

func NewAdapter(repository ports.DbPort, core ports.CorePort) *Handler {
	return &Handler{repository: repository, core: core}
}

func (h Handler) RegisterRoutes(app *fiber.App) {
	app.Get("/api/show-basket/:customerId", h.ShowBasket)
	app.Get("/api/list-products", h.ListProducts)
	app.Post("/api/add-to-basket", h.AddToBasket)
	app.Post("/api/delete-basket-item", h.DeleteBasketItem)
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

	err = h.repository.AddToBasket(request.CustomerId, request.ProductId)
	if err != nil {
		log.Printf("error while adding item to basket for the customer %d, the error: %v", request.CustomerId, err)
		c.Status(fiber.StatusInternalServerError)
		return fiber.ErrInternalServerError
	}
	basket, err := h.repository.GetBasket(request.CustomerId)
	if err != nil {
		log.Printf("error while getting basket of the customer with id : %d  err : %v", request.CustomerId, err)
		c.Status(fiber.StatusInternalServerError)
		return fiber.ErrInternalServerError
	}

	monthlyOrders, err := h.repository.GetMonthlyOrders(request.CustomerId)
	if err != nil {
		log.Printf("error while getting monthly orders of the customer %d error : %v", request.CustomerId, err)
		c.Status(fiber.StatusInternalServerError)
		return fiber.ErrInternalServerError
	}
	allOrders, err := h.repository.GetAllOrders(request.CustomerId)
	if err != nil {
		log.Printf("error while getting all orders of the customer %d error : %v", request.CustomerId, err)
		c.Status(fiber.StatusInternalServerError)
		return fiber.ErrInternalServerError
	}
	monthlyLimit, basketLimit := h.repository.GetLimits()
	discountedBasket := h.core.CalculateDiscount(basket, monthlyOrders,
		allOrders, monthlyLimit, basketLimit)
	err = h.repository.UpdateBasketProducts(discountedBasket)
	if err != nil {
		log.Printf("error while updating customers basket with the discounted basket customer id : %d error : %v",
			request.CustomerId, err)
		c.Status(fiber.StatusInternalServerError)
		return fiber.ErrInternalServerError
	}

	response := models.AddedBasketResponse{
		AddedProductId: request.ProductId,
		Message:        "Successfully Added Item To Your Basket",
		StatusCode:     fiber.StatusOK,
	}
	return c.JSON(response)
}

func (h Handler) ShowBasket(c *fiber.Ctx) error {
	customerId, err := c.ParamsInt("customerId")
	if err != nil {
		return err
	}
	basket, err := h.repository.GetBasket(customerId)
	if err != nil {
		log.Printf("error while getting the basket of the customer %d the error : %v", customerId, err)
		c.Status(fiber.StatusInternalServerError)
		return err
	}
	total, totalDiscount, totalWithDiscount, discountRate := h.core.GetDiscount(basket)
	response := models.ShowBasketResponse{
		StatusCode:        200,
		Total:             total,
		TotalDiscount:     totalDiscount,
		TotalWithDiscount: totalWithDiscount,
		DiscountRate:      discountRate,
		BasketProducts:    basket,
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

	err = h.repository.DeleteBasketItem(request.CustomerId, request.ProductId)
	if err != nil {
		log.Printf("error while deleting item to basket for the customer %d, the error: %v",
			request.CustomerId, err)
		c.Status(fiber.StatusInternalServerError)
		return fiber.ErrInternalServerError
	}

	basket, err := h.repository.GetBasket(request.CustomerId)
	if err != nil {
		log.Printf("error while getting basket of the customer with id : %d  err : %v", request.CustomerId, err)
		c.Status(fiber.StatusInternalServerError)
		return fiber.ErrInternalServerError
	}

	monthlyOrders, err := h.repository.GetMonthlyOrders(request.CustomerId)
	if err != nil {
		log.Printf("error while getting monthly orders of the customer %d error : %v", request.CustomerId, err)
		c.Status(fiber.StatusInternalServerError)
		return fiber.ErrInternalServerError
	}

	allOrders, err := h.repository.GetAllOrders(request.CustomerId)
	if err != nil {
		log.Printf("error while getting all orders of the customer %d error : %v", request.CustomerId, err)
		c.Status(fiber.StatusInternalServerError)
		return fiber.ErrInternalServerError
	}

	monthlyLimit, totalLimit := h.repository.GetLimits()
	discountedBasket := h.core.CalculateDiscount(basket, monthlyOrders,
		allOrders, monthlyLimit, totalLimit)

	err = h.repository.UpdateBasketProducts(discountedBasket)
	if err != nil {
		log.Printf("error while updating customers basket with the discounted basket customer id : %d error : %v",
			request.CustomerId, err)
		c.Status(fiber.StatusInternalServerError)
		return fiber.ErrInternalServerError
	}

	response := models.DeleteBasketResponse{
		DeletedProductsId: request.ProductId,
		Message:           "Successfully Deleted Item From Your Basket",
		StatusCode:        fiber.StatusOK,
	}
	return c.JSON(response)
}

func (h Handler) CompleteOrder(c *fiber.Ctx) error {
	customerId, err := c.ParamsInt("customerId")
	if err != nil {
		c.Status(fiber.StatusBadRequest)
		return fiber.ErrBadRequest
	}
	basket, err := h.repository.GetBasket(customerId)

	total, totalDiscount, totalWithDiscount, discountRate := h.core.GetDiscount(basket)

	err = h.repository.CompleteOrder(customerId, total, totalDiscount, discountRate, totalWithDiscount)
	if err != nil {
		log.Printf("error while completing order for custormer : %d error : %v",
			customerId, err)
		c.Status(fiber.StatusInternalServerError)
		return fiber.ErrInternalServerError
	}
	response := models.CompleteOrderResponse{
		StatusCode:        200,
		Total:             total,
		TotalDiscount:     totalDiscount,
		TotalWithDiscount: totalWithDiscount,
		DiscountRate:      discountRate,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}

func (h Handler) ListProducts(c *fiber.Ctx) error {
	products, err := h.repository.ListProducts()
	if err != nil {
		log.Printf("error while getting products from database error := %v", err)
		return fiber.ErrInternalServerError
	}
	response := models.ListProductsResponse{
		StatusCode:     200,
		BasketProducts: products,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
