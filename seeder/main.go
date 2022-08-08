package main

import (
	"PropertyFinder/models"
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

func main() {
	seedDatabase()
	os.Exit(0)
}

func seedDatabase() {
	dsn := "host=db user=postgres password=postgres dbname=propertyFinder port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("DB connection failure : %v", err)
		log.Fatalf("The Dsn is : %v", dsn)
	}
	if err = db.First(&models.Product{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		product := models.Product{
			Id:         1,
			Price:      100,
			VATPercent: 1,
			VATPrice:   1,
			TotalPrice: 101,
			Stock:      1500,
		}
		db.Create(&product)
		product.Id = 2
		product.VATPercent = 8
		product.VATPrice = 8
		product.TotalPrice = 108
		db.Create(&product)
		product.Id = 3
		product.VATPercent = 18
		product.VATPrice = 18
		product.TotalPrice = 118
		db.Create(&product)

	}
	if err = db.First(&models.Customer{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		customer := models.Customer{
			Id:      1,
			Name:    "Burak",
			Surname: "Deniz",
		}
		db.Create(&customer)
		customer.Id = 2
		customer.Name = "Sevde"
		customer.Surname = "Kucukvar"
		db.Create(&customer)

	}
	if err = db.First(&models.Limit{}).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		monthlyLimit := models.Limit{
			Id:         1,
			Name:       "monthly",
			LimitValue: 50000,
		}
		db.Create(&monthlyLimit)
		basketLimit := models.Limit{
			Id:         2,
			Name:       "basket",
			LimitValue: 500,
		}
		db.Create(&basketLimit)
	}
	log.Println("Database seed is complete")
}
