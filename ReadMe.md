# Basket Project

# How to Run The Application

- Make sure that you have docker installed
- Run the "docker-compose up --build" command inside the main folder that you clone the project
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
Adapters

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
add a product to his/her basket handler layer gets the basket information and the product  that needs to be added and 
send it to Service Layer. Service Layer checks these data and produce a discounted basket if any discount is eligible
for the current basket of the customer. Finally, it returns the discounted basket and shows it to the customer as a 
response.


# Development Cycle Of The Application

## Backend
The Backend side of the application has been developed with Go and for the database service I used the Postgres database

Building the application I used TDD approach
- Only unit test for the core logic was written for this application.
  - There are two functions in the core logic and main functions is CalculateDiscount. Since there are 4 different
scenarios for the discount I have implemented tests for each of these scenarios which are : No Discount, Discount Type A
, Discount Type B , Discount Type C. 
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
curl --location --request POST 'http://localhost:8000/api/delete-basket-item' \
--header 'Content-Type: application/json' \
--data-raw '{
    "customer_id" : 1,
    "product_id": 1
}'
```

### Complete Order 
```
curl --location --request POST 'http://localhost:8000/api/complete-order/1'
```
