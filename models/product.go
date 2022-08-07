package models

type Product struct {
	Id         int
	Price      float64
	VATPercent int
	VATPrice   float64
	TotalPrice float64
	Stock      int
}
