package ports

import "PropertyFinder/service/models"

type ServicePort interface {
	AddToBasket(customerId, productId int) (models.AddedBasketResponse, error)
	ShowBasket(customerId int) (models.ShowBasketResponse, error)
	DeleteBasketItem(customerId, productId int) (models.DeleteBasketResponse, error)
	CompleteOrder(customerId int) (models.CompleteOrderResponse, error)
	ListProducts() (models.ListProductsResponse, error)
}
