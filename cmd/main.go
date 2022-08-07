package main

import (
	"PropertyFinder/handler"
	"PropertyFinder/ports"
	"PropertyFinder/repository"
	"PropertyFinder/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// ports

	dsn := "host=db user=postgres password=postgres dbname=propertyFinder port=5432"
	var repositoryLayer ports.DbPort
	repositoryLayer, _ = repository.NewAdapter(dsn)

	var core ports.CorePort
	core = service.NewAdapter()

	var api ports.BasketHandlerPort
	api = handler.NewAdapter(repositoryLayer, core)

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	api.RegisterRoutes(app)

	fmt.Println("We have set the routes up")

	err := app.Listen(":8000")
	if err != nil {
		fmt.Printf("We have encoutered an error while running the application err : %v\n", err)
	}

	//var coreLayer core.CorePort
	//coreLayer := core.NewAdapter()
	//
	//coreLayer.CalculateDiscount()

	//customers, _ := repository.GetCustomers()
	//fmt.Println("=====Customers=====")
	//fmt.Println(customers)
	//
	//productList,_ := repository.ListProducts()
	//fmt.Println("=====Products====")
	//fmt.Println(productList)
	//
	//basketProducts,_ := repository.GetBasket(1)
	//fmt.Println("====Basket Products====")
	//fmt.Println(basketProducts)

	//err := repository.DeleteBasketItem(1,2)
	//if err != nil{
	//	fmt.Println(err)
	//}
	////
	//fmt.Println("====Basket Products After Delete====")
	//basketProducts,_ = repository.GetBasket(1)
	//fmt.Println(basketProducts)

	//err := repository.AddToBasket(1,2)
	//if err != nil{
	//	fmt.Println(err)
	//}
	//
	//basketProducts,_ = repository.GetBasket(1)
	//fmt.Println(basketProducts)
	//fmt.Println("====Basket Products After Delete====")
}
