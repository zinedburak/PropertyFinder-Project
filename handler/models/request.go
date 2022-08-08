package models

type AddToBasketRequest struct {
	CustomerId int `json:"customer_id"`
	ProductId  int `json:"product_id"`
}

type DeleteBasketRequest struct {
	CustomerId int `json:"customer_id"`
	ProductId  int `json:"product_id"`
}
