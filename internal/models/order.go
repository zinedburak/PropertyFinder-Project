package models

import "time"

type Order struct {
	Id                int
	CustomerId        int
	BasketId          int
	Discount          float64
	DiscountRate      float64
	TotalWithDiscount float64
	Total             float64
	CreatedAt         time.Time
}
