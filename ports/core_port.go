package ports

import "PropertyFinder/models"

type CorePort interface {
	CalculateDiscount(basketProducts []models.BasketProduct, monthlyOrders []models.Order, allOrders []models.Order,
		givenAmountMonth, givenAmountBasket float64) []models.BasketProduct
	GetDiscount(discountedBasket []models.BasketProduct) (float64, float64, float64, float64)
}
