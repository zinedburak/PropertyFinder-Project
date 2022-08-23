package models

import (
	"PropertyFinder/internal/models"
)

type DeleteBasketResponse struct {
	DeletedProductsId int    `json:"deleted_products_id"`
	Message           string `json:"message"`
	StatusCode        int    `json:"status_code"`
}

type AddedBasketResponse struct {
	AddedProductId int    `json:"added_product_id"`
	Message        string `json:"message"`
	StatusCode     int    `json:"status_code"`
}

type ShowBasketResponse struct {
	StatusCode        int                    `json:"status_code"`
	Total             float64                `json:"total"`
	TotalDiscount     float64                `json:"total_discount"`
	TotalWithDiscount float64                `json:"total_with_discount"`
	DiscountRate      float64                `json:"discount_rate"`
	BasketProducts    []models.BasketProduct `json:"basket_products"`
}

type CompleteOrderResponse struct {
	StatusCode        int     `json:"status_code"`
	Total             float64 `json:"total"`
	TotalDiscount     float64 `json:"total_discount"`
	TotalWithDiscount float64 `json:"total_with_discount"`
	DiscountRate      float64 `json:"discount_rate"`
}

type ListProductsResponse struct {
	StatusCode int              `json:"status_code"`
	Products   []models.Product `json:"basket_products"`
}
