package campaign_improved

import "PropertyFinder/internal/models"

type FourthOrderDiscount struct {
	FourthOrderGivenAmount float64
	AllOrders              []models.Order
}

func NewFourthOrderDiscount(givenAmount float64, allOrders []models.Order) *FourthOrderDiscount {
	return &FourthOrderDiscount{
		FourthOrderGivenAmount: givenAmount,
		AllOrders:              allOrders,
	}
}

func (d FourthOrderDiscount) CalculateDiscount(basketProducts []models.BasketProduct) ([]models.BasketProduct, float64) {
	var basketPriceTotal float64
	var totalDiscount float64
	for _, basketProduct := range basketProducts {
		basketPriceTotal += basketProduct.TotalPrice
	}
	if checkFourthOrderDiscount(d.AllOrders, basketPriceTotal, d.FourthOrderGivenAmount) {
		for index, basketProduct := range basketProducts {
			if basketProduct.VATPercent == 8 {
				discount := basketProduct.TotalPrice * 0.1
				totalDiscount += discount

				basketProducts[index].TotalWithDiscount = basketProduct.TotalPrice - discount
				basketProducts[index].DiscountRate = 10
				basketProducts[index].DiscountPrice = discount
			} else if basketProduct.VATPercent == 18 {
				discount := basketProduct.TotalPrice * 0.15
				totalDiscount += discount

				basketProducts[index].TotalWithDiscount = basketProduct.TotalPrice - discount
				basketProducts[index].DiscountRate = 15
				basketProducts[index].DiscountPrice = discount
			}
		}
	}
	return basketProducts, totalDiscount

}

func checkFourthOrderDiscount(allOrders []models.Order, basketPriceTotal, fourthOrderGivenAmount float64) bool {
	if basketPriceTotal < fourthOrderGivenAmount {
		return false
	}
	eligibleOrderCount := 1
	for _, order := range allOrders {
		if order.TotalWithDiscount > fourthOrderGivenAmount {
			eligibleOrderCount += 1
		}
	}

	return eligibleOrderCount%4 == 0
}
