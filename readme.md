# Mini Ecommerce API

Simple REST API 

## Installation With Docker
1. Simple Installation only run command below and need install docker
```bash
docker-compose up
```
## Installation run local machine
1. Create database Postgresql with name `postgres` and port `localhost:5432`
2. install go migrate
```bash
go install -tags ‘postgres’ github.com/golang-migrate/migrate/v4/cmd/migrate@latest
```
3. run migrate Database
```bash
migrate -database "postgres://postgres:secrect@localhost:5432/postgres?sslmode=disable" -path /migrations 
```
4. just run with command
```bash
go run main.go
```

## List Endpoint and Test Case
Import postman collection 
```bash
User-Service.postman_collection.json
Product-Service.postman_collection.json
Warehouse-Service.postman_collection.json
Shop-Service.postman_collection.json
Order-Service.postman_collection.json
```


## Structure Service

### controllers

controllers folder hosts all the structs under controllers namespace, controllers are the handler of all requests coming in, to the router, its doing just that, business logic and data access layer should be done separately.

controller struct implement services through their interface, no direct services implementation should be done in controller, this is done to maintain decoupled systems. The implementation will be injected during the compiled time.


### adapter

adapter folder host all structs under infrasctructures namespace, infrasctructures consists of setup for the system to connect to external data source, it is used to host things like database connection configurations, MySQL, MariaDB, MongoDB, DynamoDB.

### models

models folder hosts all structs under models namespace, model is a struct reflecting our data object from / to database. models should only define data structs, no other functionalities should be included here.

### repositories

repositories folder hosts all structs under repositories namespace, repositories is where the implementation of data access layer. All queries and data operation from / to database should happen here, and the implementor should be agnostic of what is the database engine is used, how the queries is done, all they care is they can pull the data according to the interface they are implementing.

### services

services folder hosts all structs under services namespace, services is where the business logic lies on, it handles controller request and fetch data from data layer it needs and run their logic to satisfy what controller expect the service to return.

controller might implement many services interface to satisfy the request needs, and controller should be agnostic of how services implements their logic, all they care is that they should be able to pull the result they need according to the interface they implements.


### Logger

Standard logger output that displays, _reqId_, _request header_, _request body_, _response body_, _response code_, _execution time request_, etc. The output from this logger is JSON.


### Middleware

Middleware is intercepts incoming requests and outgoing responses and allows you to modify them. Middleware can be used to perform a wide range of tasks such as authentication, logging, compression, and caching.

### main.go

main.go is the entry point of our system, here lies the router bindings it triggers ChiRouter singleton and call InitRouter to bind the router.

### route.go

route.go is where we binds controllers to appropriate route to handle desired http request. By default we are using Echo router as it is a light weight router and not bloated with unnecessary unwanted features.

