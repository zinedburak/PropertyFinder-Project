package ports

import "PropertyFinder/models"

type DbPort interface {
	ListProducts() ([]models.Product, error)
	GetBasket(customerId int) ([]models.BasketProduct, error)
	DeleteBasketItem(customerId, productId int) error
	AddToBasket(customerId, productId int) error
	CompleteOrder(customerId int, total, totalDiscount, discountRate, totalWithDiscount float64) error
	GetMonthlyOrders(customerId int) ([]models.Order, error)
	GetAllOrders(customerId int) ([]models.Order, error)
	UpdateBasketProducts(basketProducts []models.BasketProduct) error
	GetLimits() (float64, float64)
}
