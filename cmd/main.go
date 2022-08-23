package main

import (
	"PropertyFinder/internal/campaing"
	"PropertyFinder/internal/handler"
	"PropertyFinder/internal/repository"
	"PropertyFinder/internal/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	// ports

	dsn := "host=db user=postgres password=postgres dbname=propertyFinder port=5432"

	repositoryLayer, _ := repository.NewAdapter(dsn)

	coreLayer := campaing.NewAdapter()

	serviceLayer := service.NewAdapter(repositoryLayer, coreLayer)

	api := handler.NewAdapter(serviceLayer)

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
