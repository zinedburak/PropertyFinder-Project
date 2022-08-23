package campaing

import (
	"PropertyFinder/internal/models"
	"log"
	"math"
)

type Core struct {
}

func NewAdapter() *Core {
	return &Core{}
}

func (c Core) CalculateDiscount(basketProducts []models.BasketProduct, monthlyOrders []models.Order, allOrders []models.Order,
	givenAmountMonth, givenAmountBasket float64) []models.BasketProduct {

	var basketPriceTotal float64

	var discountB float64

	itemCount := make(map[int]int)

	for index, basketProduct := range basketProducts {
		basketPriceTotal += basketProduct.TotalPrice
		itemCount[basketProduct.ProductId] += 1

		// Discount type b check
		if itemCount[basketProduct.ProductId] > 3 {
			basketProducts[index].DiscountRate = 8
			basketProducts[index].TotalWithDiscount = basketProduct.TotalPrice * 0.92
			basketProducts[index].DiscountPrice = basketProduct.TotalPrice * 0.08
			discountB += basketProducts[index].DiscountPrice
		} else {
			basketProducts[index].DiscountRate = 0
			basketProducts[index].TotalWithDiscount = basketProduct.TotalPrice
			basketProducts[index].DiscountPrice = 0
		}
	}
	discountA, basketProductA := getDiscountTotalTypeA(allOrders, basketProducts, basketPriceTotal, givenAmountBasket)
	discountC, basketProductC := getDiscountTotalTypeC(monthlyOrders, basketProducts, basketPriceTotal, givenAmountMonth)
	maxDiscount := math.Max(discountA, math.Max(discountB, discountC))
	if maxDiscount == 0 {
		log.Printf("We are not applying any discount ")
		basketProducts = resetDiscount(basketProducts)
		return basketProducts
	}
	if maxDiscount == discountA {
		log.Printf("We are applying discount type A the total discount is : %f", discountA)
		return basketProductA
	} else if maxDiscount == discountB {
		log.Printf("We are applying discount type B the total discount is : %f", discountB)
		return basketProducts
	}

	log.Printf("We are applying discount type C the total discount is : %f", discountC)
	return basketProductC

}
func (c Core) GetDiscount(discountedBasket []models.BasketProduct) (float64, float64, float64, float64) {
	var totalDiscount float64
	var discountRate float64
	var total float64

	for _, product := range discountedBasket {
		totalDiscount += product.DiscountPrice
		total += product.TotalPrice
	}
	discountRate = totalDiscount / total * 100

	return total, totalDiscount, total - totalDiscount, discountRate

}

func getDiscountTotalTypeA(allOrders []models.Order, basketProducts []models.BasketProduct, basketPriceTotal, givenAmountBasket float64) (float64, []models.BasketProduct) {
	var totalDiscount float64
	if checkDiscountTypeA(allOrders, basketPriceTotal, givenAmountBasket) {
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
	return totalDiscount, basketProducts
}

func getDiscountTotalTypeC(monthlyOrders []models.Order, basketProducts []models.BasketProduct, basketPriceTotal, givenAmountMonthly float64) (float64, []models.BasketProduct) {
	var discount float64
	if checkDiscountTypeC(monthlyOrders, givenAmountMonthly) {
		discount = basketPriceTotal * 0.1
		for index, basketProduct := range basketProducts {
			basketProducts[index].DiscountRate = 10
			basketProducts[index].DiscountPrice = basketProduct.TotalPrice * 0.1
			basketProducts[index].TotalWithDiscount = basketProduct.TotalPrice * 0.9
		}
	}
	return discount, basketProducts
}

func checkDiscountTypeC(monthlyOrders []models.Order, givenAmountMonthly float64) bool {
	var totalOrderPrice float64
	for _, order := range monthlyOrders {
		totalOrderPrice += order.TotalWithDiscount
	}
	return totalOrderPrice > givenAmountMonthly
}

func checkDiscountTypeA(allOrders []models.Order, basketPriceTotal float64, givenAmountBasket float64) bool {
	if basketPriceTotal < givenAmountBasket {
		return false
	}
	eligibleOrderCount := 1
	for _, order := range allOrders {
		if order.TotalWithDiscount > givenAmountBasket {
			eligibleOrderCount += 1
		}
	}

	return eligibleOrderCount%4 == 0
}

func resetDiscount(basketProducts []models.BasketProduct) []models.BasketProduct {
	for index := range basketProducts {
		basketProducts[index].TotalWithDiscount = basketProducts[index].TotalPrice
		basketProducts[index].DiscountPrice = 0
		basketProducts[index].DiscountRate = 0
	}
	return basketProducts
}
