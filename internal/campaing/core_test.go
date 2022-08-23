package campaing

import (
	"PropertyFinder/internal/models"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCore_CalculateDiscount(t *testing.T) {
	core := NewAdapter()
	var initialMonthlyOrder []models.Order
	var initialAllOrders []models.Order
	for i := 0; i < 2; i++ {
		order := models.Order{
			Id:                i,
			CustomerId:        1,
			BasketId:          i,
			Discount:          0,
			DiscountRate:      0,
			TotalWithDiscount: 1000,
			Total:             1000,
			CreatedAt:         time.Time{},
		}
		initialMonthlyOrder = append(initialMonthlyOrder, order)
		initialAllOrders = append(initialAllOrders, order)
	}

	var initialBasketProducts []models.BasketProduct
	for i := 0; i < 2; i++ {
		basketProduct := models.BasketProduct{
			Id:                4,
			BasketId:          3,
			ProductId:         1,
			DiscountRate:      0,
			DiscountPrice:     0,
			TotalPrice:        1080,
			TotalWithDiscount: 1080,
			VATPercent:        i*7 + 1,
			VATPrice:          80,
		}
		initialBasketProducts = append(initialBasketProducts, basketProduct)
	}

	givenAmountMonthly := 5000.0
	givenAmountBasket := 500.0

	t.Run("Given there should not be a discount", func(t *testing.T) {
		discountedBasketProducts := core.CalculateDiscount(initialBasketProducts, initialMonthlyOrder,
			initialAllOrders, givenAmountMonthly, givenAmountBasket)
		for _, basketProduct := range discountedBasketProducts {
			assert.Equal(t, basketProduct.TotalPrice, basketProduct.TotalWithDiscount)
			assert.EqualValues(t, 0, basketProduct.DiscountPrice)
			assert.EqualValues(t, 0, basketProduct.DiscountRate)
		}
	})

	t.Run("Given Type A discount should be applied", func(t *testing.T) {
		thirdMonthlyOrder := models.Order{
			TotalWithDiscount: 1000,
			Total:             1000,
			CreatedAt:         time.Time{},
		}
		basketProduct := models.BasketProduct{
			ProductId:         10,
			DiscountRate:      0,
			DiscountPrice:     0,
			TotalPrice:        1800,
			TotalWithDiscount: 1800,
			VATPercent:        18,
			VATPrice:          324,
		}
		typeAAllOrders := append(initialAllOrders, thirdMonthlyOrder)

		typeABasketProducts := append(initialBasketProducts, basketProduct)
		discountedBasketProducts := core.CalculateDiscount(typeABasketProducts, initialMonthlyOrder,
			typeAAllOrders, givenAmountMonthly, givenAmountBasket)

		for _, discountedProduct := range discountedBasketProducts {
			if discountedProduct.VATPercent == 8 {
				assert.EqualValues(t, discountedProduct.TotalPrice*0.10, discountedProduct.DiscountPrice)
				assert.EqualValues(t, 10, discountedProduct.DiscountRate)
				assert.EqualValues(t, discountedProduct.TotalPrice*0.9, discountedProduct.TotalWithDiscount)
			} else if discountedProduct.VATPercent == 18 {
				assert.EqualValues(t, discountedProduct.TotalPrice*0.15, discountedProduct.DiscountPrice)
				assert.EqualValues(t, 15, discountedProduct.DiscountRate)
				assert.EqualValues(t, discountedProduct.TotalPrice*0.85, discountedProduct.TotalWithDiscount)
			} else if discountedProduct.VATPercent == 1 {
				assert.EqualValues(t, 0, basketProduct.DiscountPrice)
				assert.EqualValues(t, 0, basketProduct.DiscountRate)
				assert.EqualValues(t, basketProduct.TotalPrice, basketProduct.TotalWithDiscount)
			}
		}

	})

	t.Run("Given Type B discount should be applied", func(t *testing.T) {
		thirdItem := models.BasketProduct{
			Id:                6,
			BasketId:          3,
			ProductId:         1,
			DiscountRate:      0,
			DiscountPrice:     0,
			TotalPrice:        1080,
			TotalWithDiscount: 1080,
			VATPercent:        8,
			VATPrice:          80,
		}
		discountedItem := models.BasketProduct{
			Id:                5,
			BasketId:          3,
			ProductId:         1,
			DiscountRate:      0,
			DiscountPrice:     0,
			TotalPrice:        1080,
			TotalWithDiscount: 1080,
			VATPercent:        8,
			VATPrice:          80,
		}

		//initialMonthlyOrder = []models.Order{}
		//initialAllOrders = []models.Order{}

		typeBBasketProducts := append(initialBasketProducts, thirdItem)
		typeBBasketProducts = append(typeBBasketProducts, discountedItem)
		discountedBasket := core.CalculateDiscount(typeBBasketProducts, initialMonthlyOrder,
			initialAllOrders, givenAmountMonthly, givenAmountBasket)
		//fmt.Println(discountedBasket)
		lastItem := discountedBasket[len(discountedBasket)-1]
		assert.EqualValues(t, thirdItem.TotalPrice*0.08, lastItem.DiscountPrice)
		assert.EqualValues(t, 8, lastItem.DiscountRate)
		assert.EqualValues(t, thirdItem.TotalPrice*0.92, lastItem.TotalWithDiscount)

	})
	t.Run("Given Type C discount should be applied", func(t *testing.T) {
		lastOrder := models.Order{
			Id:                0,
			CustomerId:        1,
			BasketId:          3,
			Discount:          0,
			DiscountRate:      0,
			TotalWithDiscount: 3001,
			Total:             3001,
			CreatedAt:         time.Time{},
		}
		typeCMonthlyOrder := append(initialMonthlyOrder, lastOrder)
		discountedBasket := core.CalculateDiscount(initialBasketProducts, typeCMonthlyOrder,
			initialAllOrders, givenAmountMonthly, givenAmountBasket)
		for _, discountedProduct := range discountedBasket {
			assert.EqualValues(t, 10, discountedProduct.DiscountRate)
			assert.EqualValues(t, discountedProduct.TotalPrice*0.1, discountedProduct.DiscountPrice)
			assert.EqualValues(t, discountedProduct.TotalPrice*0.9, discountedProduct.TotalWithDiscount)

		}

	})

}

func TestCore_GetDiscount(t *testing.T) {
	var discountedBasket []models.BasketProduct
	core := NewAdapter()
	for i := 0; i < 10; i++ {
		basketProduct := models.BasketProduct{
			DiscountRate:      10,
			DiscountPrice:     100,
			TotalPrice:        1000,
			TotalWithDiscount: 900,
		}
		discountedBasket = append(discountedBasket, basketProduct)
	}
	total, totalDiscount, totalWithDiscount, discountRate := core.GetDiscount(discountedBasket)

	assert.EqualValues(t, 10000, total)
	assert.EqualValues(t, 1000, totalDiscount)
	assert.EqualValues(t, 9000, totalWithDiscount)
	assert.EqualValues(t, 10, discountRate)

}
