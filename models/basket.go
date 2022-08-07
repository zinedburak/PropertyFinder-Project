package models

type Basket struct {
	Id          int     `json:"id" gorm:"primaryKey"`
	CustomerId  int     `json:"customerId"`
	BasketTotal float64 `json:"basket_total"`
	IsPurchased bool    `json:"isPurchased"`
}
