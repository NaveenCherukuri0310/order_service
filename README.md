## Order Service API

A production ready Order Management API built using Go, Gin, and MongoDB.
This service demonstrates clean architecture, environment variable usage, and full CRUD operations.

## Features

* Create a new order
* Get all orders
* Get order by ID
* Update an order
* Delete an order
* Uses .env configuration
* MongoDB connection with timeout
* Clean folder structure
* Ping route separated into its own controller

## Project Structure

```
order_service
│
├── cmd
│   └── main.go
│
├── internal
│   ├── controllers
│   │     ├──ping_controller.go
│   │     ├──create_order.go
│   │     ├──get_orders.go
│   │     ├──get_order_by_id.go
│   │     └──update_order.go
│   │     
│   │
│   ├── routes
│   │     ├──ping_routes.go
│   │     └── order_routes.go
│   │
│   ├── database
│   │     └── onnection.go
│   │
│   ├── models
│         └── order.go
│
├── sample.env
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Requirements

* Go installed
* MongoDB running locally or using MongoDB Atlas
* Postman or Thunder Client for API testing

## Setup Instructions

### Clone the repository

```
git clone <your repository url>
cd order_service
```

### Create your environment file

```
cp sample.env .env
```

Now open the new .env file and set values

```
MONGO_URI=mongodb://localhost:27017
DB_NAME=orderdb
PORT=8080
```
#### Important Notes

* Do not commit .env to version control
* Do not share real database credentials in README or anywhere in the repository
* Always use sample.env to communicate required variable names
* The application validates environment variables at startup to prevent misconfiguration

### Install dependencies

```
go mod tidy
```

### Run the server

```
go run cmd/main.go
```

Expected output

```
Connected to MongoDB
Listening and serving HTTP on :8080
```

## API Endpoints

### Health Check

```
GET /ping
```

Response

```
{
  "message": "pong"
}
```

### Create Order

```
POST /orders
```

Body example

```
{
  "item": "Laptop",
  "quantity": 2,
  "price": 1299.99,
  "status": "new"
}
```

### Get All Orders

```
GET /orders
```

### Get Order by ID

```
GET /orders/:id
```

### Update Order

```
PUT /orders/:id
```

Body example

```
{
  "item": "Laptop Pro",
  "quantity": 3,
  "price": 1599.50,
  "status": "updated"
}
```


## Environment Variables

Use `.env` in the project root.
Never commit `.env` to Git.
Use `sample.env` as a template.

Content of sample.env

```
MONGO_URI=
DB_NAME=
PORT=
```


## License

This project is for learning and development use.


