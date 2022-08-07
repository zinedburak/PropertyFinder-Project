package models

type BasketProduct struct {
	Id                int
	BasketId          int
	ProductId         int
	DiscountRate      int
	DiscountPrice     float64
	TotalPrice        float64
	TotalWithDiscount float64
	VATPercent        int
	VATPrice          float64
}
