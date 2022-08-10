package main

import (
	"PropertyFinder/campaing"
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

	var coreLayer ports.CorePort

	coreLayer = campaing.NewAdapter()

	var serviceLayer ports.ServicePort
	serviceLayer = service.NewAdapter(repositoryLayer, coreLayer)

	var api ports.BasketHandlerPort
	api = handler.NewAdapter(serviceLayer)

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

}
