package repository

import (
	"PropertyFinder/internal/models"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"time"
)

type Adapter struct {
	db *gorm.DB
}

func NewAdapter(dsn string) (*Adapter, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("DB connection failure : %v", err)
		log.Fatalf("The Dsn is : %v", dsn)
	}

	err = db.AutoMigrate(models.Product{}, models.Customer{}, models.Basket{},
		models.BasketProduct{}, models.Order{}, models.Limit{})
	if err != nil {
		return nil, err
	}
	return &Adapter{db: db}, nil
}

// ListProducts Functionality
func (a Adapter) ListProducts() ([]models.Product, error) {
	var products []models.Product
	err := a.db.Find(&products).Error
	if err != nil {
		err = errors.New("error while listing products ")
		return nil, err
	}
	return products, nil
}

// AddToBasket functionality
func (a Adapter) AddToBasket(customerId int, productId int) error {
	var customer models.Customer

	a.db.Where("id = ?", customerId).Find(&customer)
	if customer.Id == 0 {
		err := errors.New("There is no customer with id the given id ")
		return err
	}

	var basket models.Basket
	a.db.Where("customer_id = ? and is_purchased = ?", customerId, false).Find(&basket)

	if basket.Id == 0 {
		basket.CustomerId = customerId
		basket.IsPurchased = false
		a.db.Create(&basket)
	}

	var product models.Product
	a.db.Where("id = ?and stock > ?", productId, 1).Find(&product)

	if product.Stock < 1 {
		err := errors.New("there is not enough stock for you to purchase this item")
		return err
	}

	var basketProduct models.BasketProduct
	basketProduct.BasketId = basket.Id
	basketProduct.ProductId = productId
	basketProduct.TotalPrice = product.TotalPrice
	basketProduct.TotalWithDiscount = product.TotalPrice
	basketProduct.VATPercent = product.VATPercent
	basketProduct.VATPrice = product.VATPrice

	err := a.db.Create(&basketProduct).Error
	if err != nil {
		return err
	}

	return nil
}

// GetBasket Functionality
func (a Adapter) GetBasket(customerId int) ([]models.BasketProduct, error) {

	var basketProducts []models.BasketProduct

	// select *
	//		from basket_products bp
	// 		inner joint baskets b on b.id = bp.basket_id
	// 		where b.customer_id = 10 and b.is_purchased = false

	err := a.db.Model(&basketProducts).Joins("JOIN baskets ON basket_products.basket_id = baskets.id").
		Where("baskets.customer_id = ? AND baskets.is_purchased = ?", customerId, false).
		Find(&basketProducts).Error

	if err != nil {
		err = fmt.Errorf("there was an error getting the customers basket data: %v", err)
		return nil, err
	}
	return basketProducts, nil
}

// DeleteBasketItem Functionality
func (a Adapter) DeleteBasketItem(customerId, productId int) error {
	var basketProduct models.BasketProduct

	err := a.db.Model(&basketProduct).
		Joins("JOIN baskets ON basket_products.basket_id = baskets.id").
		Where("baskets.customer_id = ? AND baskets.is_purchased = ? AND basket_products.product_id = ?",
			customerId, false, productId).Limit(1).Find(&basketProduct).Error

	if basketProduct.Id == 0 {
		err = errors.New("could not find the item that you want to delete in your basket")
		return err
	}
	if err != nil {
		return err
	}
	a.db.Delete(&basketProduct)
	return nil
}

// CompleteOrder Functionality
func (a Adapter) CompleteOrder(customerId int, total, totalDiscount, discountRate, totalWithDiscount float64) error {
	var activeBasket models.Basket

	err := a.db.Where("customer_id = ? AND is_purchased = ?", customerId, false).Find(&activeBasket).Error
	if err != nil {
		return err
	}

	if activeBasket.Id == 0 {
		err = errors.New("could not find a basket that you have not purchased please make sure that you haven " +
			"active basket")
		return err
	}

	var basketProducts []models.BasketProduct
	a.db.Where("basket_id = ?", activeBasket.Id).Find(&basketProducts)

	if len(basketProducts) == 0 {
		err = errors.New("could not find any item on your current basket please add items to your basket")
		return err
	}

	order := models.Order{
		CustomerId:        customerId,
		BasketId:          activeBasket.Id,
		Discount:          totalDiscount,
		DiscountRate:      discountRate,
		TotalWithDiscount: totalWithDiscount,
		Total:             total,
	}
	err = a.db.Create(&order).Error
	if err != nil {
		return err
	}

	for _, basketProduct := range basketProducts {
		var product models.Product
		a.db.Where("id = ?", basketProduct.ProductId).Find(&product)
		product.Stock -= 1
		err = a.db.Save(&product).Error
		if err != nil {
			return err
		}
	}
	activeBasket.IsPurchased = true
	a.db.Save(&activeBasket)
	return nil
}

func (a Adapter) UpdateBasketProducts(basketProducts []models.BasketProduct) error {
	for _, product := range basketProducts {
		err := a.db.Save(&product).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func (a Adapter) GetMonthlyOrders(customerId int) ([]models.Order, error) {
	monthAgo := time.Now().AddDate(0, -1, 1)
	var monthlyOrders []models.Order
	err := a.db.Where("created_at > ? AND customer_id = ?", monthAgo, customerId).
		Find(&monthlyOrders).Error
	if err != nil {
		return nil, err
	}
	return monthlyOrders, nil
}

func (a Adapter) GetAllOrders(customerId int) ([]models.Order, error) {
	var allOrders []models.Order
	err := a.db.Where("customer_id = ?", customerId).Find(&allOrders).Error
	if err != nil {
		return nil, err
	}
	return allOrders, nil
}

func (a Adapter) GetLimits() (float64, float64) {
	var limits []models.Limit
	var monthlyLimit, totalLimit float64
	a.db.Find(&limits)
	for _, limit := range limits {
		if limit.Name == "monthly" {
			monthlyLimit = limit.LimitValue
		} else if limit.Name == "basket" {
			totalLimit = limit.LimitValue
		}
	}
	return monthlyLimit, totalLimit
}
