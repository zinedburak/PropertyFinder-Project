package campaign_improved

import (
	"PropertyFinder/internal/models"
)

type DiscountCalculator struct {
	discountRules []DiscountRules
}

type DiscountRules interface {
	CalculateDiscount(basketProducts []models.BasketProduct) ([]models.BasketProduct, float64)
}

func NewDiscountCalculator(discounts []DiscountRules) *DiscountCalculator {
	return &DiscountCalculator{discountRules: discounts}
}

func (dc DiscountCalculator) CalculateDiscount(basketProducts []models.BasketProduct) ([]models.BasketProduct, float64) {
	var finalDiscount float64
	var finalBasket []models.BasketProduct
	for _, discountRule := range dc.discountRules {
		discountedBasket, discount := discountRule.CalculateDiscount(basketProducts)
		if discount > finalDiscount {
			finalDiscount = discount
			finalBasket = discountedBasket
		}
	}
	if finalDiscount == 0 {
		finalBasket = resetDiscount(basketProducts)
	}
	return finalBasket, finalDiscount
}

func resetDiscount(basketProducts []models.BasketProduct) []models.BasketProduct {
	for index := range basketProducts {
		basketProducts[index].TotalWithDiscount = basketProducts[index].TotalPrice
		basketProducts[index].DiscountPrice = 0
		basketProducts[index].DiscountRate = 0
	}
	return basketProducts
}
