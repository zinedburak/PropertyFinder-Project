package campaign_improved

import "PropertyFinder/internal/models"

type MonthlyOrderDiscount struct {
	MonthlyOrderGivenAmount float64
	MonthlyOrders           []models.Order
}

func NewMonthlyOrderDiscount(monthlyOrderGivenAmount float64, monthlyOrders []models.Order) *MonthlyOrderDiscount {
	return &MonthlyOrderDiscount{
		MonthlyOrderGivenAmount: monthlyOrderGivenAmount,
		MonthlyOrders:           monthlyOrders,
	}
}

func (d MonthlyOrderDiscount) CalculateDiscount(basketProducts []models.BasketProduct) ([]models.BasketProduct, float64) {
	var discount float64
	if checkMonthlyOrderDiscount(d.MonthlyOrders, d.MonthlyOrderGivenAmount) {
		for index, basketProduct := range basketProducts {
			basketProducts[index].DiscountRate = 10
			basketProducts[index].DiscountPrice = basketProduct.TotalPrice * 0.1
			basketProducts[index].TotalWithDiscount = basketProduct.TotalPrice * 0.9
			discount += basketProducts[index].DiscountPrice
		}
	}
	return basketProducts, discount
}

func checkMonthlyOrderDiscount(monthlyOrders []models.Order, givenAmountMonthly float64) bool {
	var totalOrderPrice float64
	for _, order := range monthlyOrders {
		totalOrderPrice += order.TotalWithDiscount
	}
	return totalOrderPrice > givenAmountMonthly
}
