## Basket Project

## How to Run The Application

- Make sure that you have docker installed
- Run the "docker-compose up --build" command inside the main folder that you clone the project
- The application will be built and run on the 8000 port in your local host with the above command
- Example curls will be given on the following sections

## Architecture

The architecture used in this project is Hexagonal Architecture. The hexagonal architecture provides the flexibility to
change any technology that you are using in your application easily by providing ports and adaptors components also by
separating layers of your application the architecture reduces coupling.

## How was The Architecture Implemented

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