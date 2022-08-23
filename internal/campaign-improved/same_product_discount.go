package campaign_improved

import "PropertyFinder/internal/models"

type SameProductDiscount struct {
}

func (d SameProductDiscount) CalculateDiscount(basketProducts []models.BasketProduct) ([]models.BasketProduct, float64) {
	var discount float64
	itemCount := make(map[int]int)

	for index, basketProduct := range basketProducts {
		itemCount[basketProduct.ProductId] += 1
		if itemCount[basketProduct.ProductId] > 3 {
			basketProducts[index].DiscountRate = 8
			basketProducts[index].TotalWithDiscount = basketProduct.TotalPrice * 0.92
			basketProducts[index].DiscountPrice = basketProduct.TotalPrice * 0.08
			discount += basketProducts[index].DiscountPrice
		}
	}
	return basketProducts, discount
}
