package ports

import (
	"PropertyFinder/internal/models"
)

type CorePort interface {
	CalculateDiscount(basketProducts []models.BasketProduct, monthlyOrders []models.Order, allOrders []models.Order,
		givenAmountMonth, givenAmountBasket float64) []models.BasketProduct
	GetDiscount(discountedBasket []models.BasketProduct) (float64, float64, float64, float64)
}
