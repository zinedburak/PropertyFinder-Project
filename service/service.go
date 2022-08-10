package service

import (
	"PropertyFinder/ports"
	"PropertyFinder/service/models"
	"errors"
	"log"
)

type Service struct {
	repository ports.DbPort
	core       ports.CorePort
}

func NewAdapter(repository ports.DbPort, core ports.CorePort) *Service {
	return &Service{repository: repository, core: core}
}

func (s Service) AddToBasket(customerId, productId int) (models.AddedBasketResponse, error) {
	err := s.repository.AddToBasket(customerId, productId)
	if err != nil {
		log.Printf("Add basket service layer error")
		return models.AddedBasketResponse{}, err
	}
	basket, err := s.repository.GetBasket(customerId)
	if err != nil {
		log.Printf("Get basket service layer error")
		return models.AddedBasketResponse{}, err
	}
	monthlyOrders, err := s.repository.GetMonthlyOrders(customerId)
	if err != nil {
		log.Printf("Get Monthly basket service layer error")
		return models.AddedBasketResponse{}, err
	}
	allOrders, err := s.repository.GetAllOrders(customerId)
	if err != nil {
		log.Printf("Get All basket service layer error")
		return models.AddedBasketResponse{}, err
	}
	monthlyLimit, basketLimit := s.repository.GetLimits()
	discountedBasket := s.core.CalculateDiscount(basket, monthlyOrders,
		allOrders, monthlyLimit, basketLimit)
	err = s.repository.UpdateBasketProducts(discountedBasket)
	if err != nil {
		log.Printf("Update basket service layer error")
		return models.AddedBasketResponse{}, err
	}

	response := models.AddedBasketResponse{
		AddedProductId: productId,
		Message:        "Successfully Added Item To Your Basket",
		StatusCode:     200,
	}
	return response, nil
}
func (s Service) ShowBasket(customerId int) (models.ShowBasketResponse, error) {
	basket, err := s.repository.GetBasket(customerId)
	if err != nil {
		log.Printf("error while getting the basket of the customer %d the error : %v", customerId, err)
		return models.ShowBasketResponse{}, err
	}
	if len(basket) == 0 {
		err = errors.New("you dont have any item in your basket please add an item")
		return models.ShowBasketResponse{}, err
	}

	total, totalDiscount, totalWithDiscount, discountRate := s.core.GetDiscount(basket)
	response := models.ShowBasketResponse{
		StatusCode:        200,
		Total:             total,
		TotalDiscount:     totalDiscount,
		TotalWithDiscount: totalWithDiscount,
		DiscountRate:      discountRate,
		BasketProducts:    basket,
	}
	return response, nil
}
func (s Service) DeleteBasketItem(customerId, productId int) (models.DeleteBasketResponse, error) {
	err := s.repository.DeleteBasketItem(customerId, productId)
	if err != nil {
		log.Printf("error while deleting item to basket for the customer %d, the error: %v",
			customerId, err)
		return models.DeleteBasketResponse{}, err
	}

	basket, err := s.repository.GetBasket(customerId)
	if err != nil {
		log.Printf("error while getting basket of the customer with id : %d  err : %v", customerId, err)
		return models.DeleteBasketResponse{}, err
	}
	monthlyOrders, err := s.repository.GetMonthlyOrders(customerId)
	if err != nil {
		log.Printf("error while getting monthly orders of the customer %d error : %v", customerId, err)
		return models.DeleteBasketResponse{}, err
	}
	allOrders, err := s.repository.GetAllOrders(customerId)
	if err != nil {
		log.Printf("error while getting all orders of the customer %d error : %v", customerId, err)
		return models.DeleteBasketResponse{}, err
	}
	monthlyLimit, totalLimit := s.repository.GetLimits()
	discountedBasket := s.core.CalculateDiscount(basket, monthlyOrders,
		allOrders, monthlyLimit, totalLimit)
	err = s.repository.UpdateBasketProducts(discountedBasket)
	if err != nil {
		log.Printf("error while updating customers basket with the discounted basket customer id : %d error : %v",
			customerId, err)
		return models.DeleteBasketResponse{}, err
	}
	response := models.DeleteBasketResponse{
		DeletedProductsId: productId,
		Message:           "Successfully Deleted Item From Your Basket",
		StatusCode:        200,
	}
	return response, nil

}
func (s Service) CompleteOrder(customerId int) (models.CompleteOrderResponse, error) {
	basket, err := s.repository.GetBasket(customerId)
	if err != nil {
		log.Printf("error while getting the basket of the customer %d the error : %v", customerId, err)
		return models.CompleteOrderResponse{}, err
	}
	total, totalDiscount, totalWithDiscount, discountRate := s.core.GetDiscount(basket)
	err = s.repository.CompleteOrder(customerId, total, totalDiscount, discountRate, totalWithDiscount)
	if err != nil {
		log.Printf("error while completing order for custormer : %d error : %v",
			customerId, err)
		return models.CompleteOrderResponse{}, err
	}
	response := models.CompleteOrderResponse{
		StatusCode:        200,
		Total:             total,
		TotalDiscount:     totalDiscount,
		TotalWithDiscount: totalWithDiscount,
		DiscountRate:      discountRate,
	}
	return response, nil
}
func (s Service) ListProducts() (models.ListProductsResponse, error) {
	products, err := s.repository.ListProducts()
	if err != nil {
		log.Printf("error while getting products from database error := %v", err)
		return models.ListProductsResponse{}, err
	}
	response := models.ListProductsResponse{
		StatusCode: 200,
		Products:   products,
	}
	return response, nil
}
