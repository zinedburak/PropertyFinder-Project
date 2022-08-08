# Basket Project

# How to Run The Application

- Make sure that you have docker installed
- Run the ```docker-compose up --build``` command inside the main folder that you clone the project
- Run the ```go run seeder/main.go``` command for seeding the database with products customers and limits
- The application will be built and run on the 8000 port in your local host with the above command
- Example curls will be given on the following sections

# Architecture

The architecture used in this project is Hexagonal Architecture. The hexagonal architecture provides the flexibility to
change any technology that you are using in your application easily by providing ports and adaptors components also by
separating layers of your application the architecture reduces coupling.

# How was The Architecture Implemented

- There are 3 Ports in this application
    - Core Port
    - DB Port
    - Handler Port

Ports provide a general structure for the layers using interfaces. Application layers connect to these ports with their
Adaptors

- There are 3 Layers in this application
    - Repository Layer
    - Service Layer
    - Handler Layer

### Repository Layer

When an instance of this layer is created we automatically connect to the DB that is running in the docker and with the
help of gorm we migrate our models to database as tables. Alongside the connection and migration of the DB we also have
all the Databases function that reads from or writes to DB in this layer.

### Service Layer

This layer holds the core business logic of the application which consists of the discount logic.

### Handler Layer

This layer uses instances of Repository port and Service port that are connected by Repository Layer and Service Layer
respectively. Handler layer is the bridge between Repository and Service Layers. For example when a customer wants to
add a product to his/her basket handler layer gets the basket information and the product that needs to be added and
send it to Service Layer. Service Layer checks these data and produce a discounted basket if any discount is eligible
for the current basket of the customer. Finally, it returns the discounted basket and shows it to the customer as a
response.

# Development Cycle Of The Application

## Backend

The Backend side of the application has been developed with Go and for the database service I used the Postgres database

Building the application I used TDD approach

- Only unit test for the core logic was written for this application.
    - There are two functions in the core logic and main functions is CalculateDiscount. Since there are 4 different
      scenarios for the discount I have implemented tests for each of these scenarios which are : No Discount, Discount
      Type A , Discount Type B , Discount Type C.
- After unit test was done I implemented the core logic for CalculateDiscount it uses helper functions to calculate
  Discount A and Discount B. With service layer of the application was done.
- Completing the service layer provided me the information that I will need from the database. I wrote database
  functionalities according to this. Used gorm library for all database transactions
- Last I implemented Handler Layer using fiber library for responding the customers requests and building a bridge
  between Repository Layer and Service Layer

# Example Requests

### List Products

```
  curl --location --request GET 'http://localhost:8000/api/list-products'
```

### Show Basket

```
  curl --location --request GET 'http://localhost:8000/api/show-basket/1'
```

### Add Basket Item

```
curl --location --request POST 'http://localhost:8000/api/add-to-basket' \
--header 'Content-Type: application/json' \
--data-raw '{
    "customer_id" : 1,
    "product_id": 1
}'
```

### Delete Basket Item

```
curl --location --request DELETE 'http://localhost:8000/api/delete-basket-item' \
--header 'Content-Type: application/json' \
--data-raw '{
    "customer_id" : 1,
    "product_id": 3
}'
```

### Complete Order

```
curl --location --request POST 'http://localhost:8000/api/complete-order/1'
```

## Example Responses

### List Products

```json
{
    "status_code": 200,
    "basket_products": [
        {
            "Id": 1,
            "Price": 100,
            "VATPercent": 8,
            "VATPrice": 8,
            "TotalPrice": 108,
            "Stock": 1465
        },
        {
            "Id": 3,
            "Price": 300,
            "VATPercent": 18,
            "VATPrice": 54,
            "TotalPrice": 354,
            "Stock": 1500
        },
        {
            "Id": 4,
            "Price": 500,
            "VATPercent": 1,
            "VATPrice": 5,
            "TotalPrice": 505,
            "Stock": 1500
        },
        {
            "Id": 2,
            "Price": 200,
            "VATPercent": 8,
            "VATPrice": 16,
            "TotalPrice": 216,
            "Stock": 1457
        }
    ]
}
```

### Show Basket Discount Type A

```json
{
    "status_code": 200,
    "total": 967,
    "total_discount": 63.900000000000006,
    "total_with_discount": 903.1,
    "discount_rate": 6.608066184074457,
    "basket_products": [
        {
            "Id": 95,
            "BasketId": 16,
            "ProductId": 4,
            "DiscountRate": 0,
            "DiscountPrice": 0,
            "TotalPrice": 505,
            "TotalWithDiscount": 505,
            "VATPercent": 1,
            "VATPrice": 5
        },
        {
            "Id": 94,
            "BasketId": 16,
            "ProductId": 3,
            "DiscountRate": 15,
            "DiscountPrice": 53.1,
            "TotalPrice": 354,
            "TotalWithDiscount": 300.9,
            "VATPercent": 18,
            "VATPrice": 54
        },
        {
            "Id": 91,
            "BasketId": 16,
            "ProductId": 1,
            "DiscountRate": 10,
            "DiscountPrice": 10.8,
            "TotalPrice": 108,
            "TotalWithDiscount": 97.2,
            "VATPercent": 8,
            "VATPrice": 8
        }
    ]
}
```

### Show Basket Discount Type B

```json
{
    "status_code": 200,
    "total": 432,
    "total_discount": 8.64,
    "total_with_discount": 423.36,
    "discount_rate": 2,
    "basket_products": [
        {
            "Id": 90,
            "BasketId": 16,
            "ProductId": 1,
            "DiscountRate": 0,
            "DiscountPrice": 0,
            "TotalPrice": 108,
            "TotalWithDiscount": 108,
            "VATPercent": 8,
            "VATPrice": 8
        },
        {
            "Id": 86,
            "BasketId": 16,
            "ProductId": 1,
            "DiscountRate": 0,
            "DiscountPrice": 0,
            "TotalPrice": 108,
            "TotalWithDiscount": 108,
            "VATPercent": 8,
            "VATPrice": 8
        },
        {
            "Id": 89,
            "BasketId": 16,
            "ProductId": 1,
            "DiscountRate": 0,
            "DiscountPrice": 0,
            "TotalPrice": 108,
            "TotalWithDiscount": 108,
            "VATPercent": 8,
            "VATPrice": 8
        },
        {
            "Id": 91,
            "BasketId": 16,
            "ProductId": 1,
            "DiscountRate": 8,
            "DiscountPrice": 8.64,
            "TotalPrice": 108,
            "TotalWithDiscount": 99.36,
            "VATPercent": 8,
            "VATPrice": 8
        }
    ]
} 
```

### Show Basket Discount Type C

```json
{
    "status_code": 200,
    "total": 1472,
    "total_discount": 147.2,
    "total_with_discount": 1324.8,
    "discount_rate": 10,
    "basket_products": [
        {
            "Id": 94,
            "BasketId": 16,
            "ProductId": 3,
            "DiscountRate": 10,
            "DiscountPrice": 35.4,
            "TotalPrice": 354,
            "TotalWithDiscount": 318.6,
            "VATPercent": 18,
            "VATPrice": 54
        },
        {
            "Id": 91,
            "BasketId": 16,
            "ProductId": 1,
            "DiscountRate": 10,
            "DiscountPrice": 10.8,
            "TotalPrice": 108,
            "TotalWithDiscount": 97.2,
            "VATPercent": 8,
            "VATPrice": 8
        },
        {
            "Id": 96,
            "BasketId": 16,
            "ProductId": 4,
            "DiscountRate": 10,
            "DiscountPrice": 50.5,
            "TotalPrice": 505,
            "TotalWithDiscount": 454.5,
            "VATPercent": 1,
            "VATPrice": 5
        },
        {
            "Id": 95,
            "BasketId": 16,
            "ProductId": 4,
            "DiscountRate": 10,
            "DiscountPrice": 50.5,
            "TotalPrice": 505,
            "TotalWithDiscount": 454.5,
            "VATPercent": 1,
            "VATPrice": 5
        }
    ]
}
```

### Complete Order

```json
{
    "status_code": 200,
    "total": 1472,
    "total_discount": 147.2,
    "total_with_discount": 1324.8,
    "discount_rate": 10
}
```

### Add Basket Item

```json
{
    "added_product_id": 4,
    "message": "Successfully Added Item To Your Basket",
    "status_code": 200
}
```

### Delete Basket Item

```json
{
    "deleted_products_id": 4,
    "message": "Successfully Deleted Item From Your Basket",
    "status_code": 200
}
```
