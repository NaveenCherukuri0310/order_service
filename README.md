# Order Service API

A production ready Order Management API built using Go Gin and MongoDB.
This service demonstrates clean architecture JWT authentication centralized logging Swagger documentation and a scalable order model with multiple items per order.

---

## Features

* Create a new order
* Get all orders
* Get order by UUID
* Update an existing order
* Multiple items per order
* Unit price and total price per item
* UUID based public order ID
* Internal MongoDB ObjectID hidden from API
* Automatic created and updated timestamps
* Disabled flag supported in data model
* JWT authentication for all order routes
* Centralized file based logging
* Swagger UI for API documentation
* Environment variable validation at startup

---

## Project Structure

```
order_service
│
├── cmd
│   └── main.go
│
├── docs
│   └── orders.yaml        # OpenAPI specification
│
├── internal
│   ├── controllers
│   │     create_order.go
│   │     get_orders.go
│   │     get_order_by_id.go
│   │     update_order.go
│   │     ping_controller.go
│   │
│   ├── routes
│   │     order_routes.go
│   │     ping_routes.go
│   │
│   ├── middleware
│   │     auth.go
│   │     logger.go
│   │
│   ├── utils
│   │     jwt.go
│   │
│   ├── logger
│   │     logger.go
│   │
│   ├── database
│   │     connection.go
│   │
│   ├── models
│         order.go
│
├── sample.env
├── go.mod
├── go.sum
├── .gitignore
└── README.md
```

---

## Requirements

* Go installed
* MongoDB Atlas or local MongoDB
* Postman Thunder Client or browser
* Valid JWT token for protected routes

---

## Setup Instructions

### Clone repository

```
git clone <your repository url>
cd order_service
```

### Create environment file

```
cp sample.env .env
```

Edit `.env`:

```
MONGO_URI=mongodb+srv://...
DB_NAME=orderdb
PORT=8080
JWT_SECRET=your_secret_here
```

### Install dependencies

```
go mod tidy
```

### Run the server

```
go run cmd/main.go
```

Expected output:

```
Loaded DB_NAME: orderdb
Connected to MongoDB
Server running on port: 8080
```

---

## Authentication

All order APIs require a JWT token.

Header format:

```
Authorization: Bearer <jwt_token>
```

### Generate test token (development only)

```
GET /test-token
```

Response:

```
{
  "token": "eyJhbGciOiJIUzI1NiIs..."
}
```

---

## API Endpoints

### Public

#### Health check

```
GET /ping
```

#### Development token

```
GET /test-token
```

---

### Protected Order APIs

All require Authorization header.

#### Create order

```
POST /orders
```

Example body:

```
{
  "status": "new",
  "items": [
    {
      "product_name": "iPhone 15",
      "quantity": 3,
      "unit_price": 999
    },
    {
      "product_name": "MacBook Air",
      "quantity": 1,
      "unit_price": 1299
    }
  ]
}
```

Server calculates total_price created_at updated_at and UUID.

---

#### Get all orders

```
GET /orders
```

Returns only orders where disabled is false.

---

#### Get order by ID

```
GET /orders/<uuid>
```

---

#### Update order

```
PUT /orders/<uuid>
```

Example body:

```
{
  "status": "updated",
  "items": [
    {
      "product_name": "iPhone 15 Pro",
      "quantity": 2,
      "unit_price": 1199
    }
  ]
}
```

Server recalculates total_price and updates updated_at.

---

## Swagger API Documentation

Swagger UI is enabled for development.

### Swagger UI

```
http://localhost:8080/swagger/index.html
```

### Raw OpenAPI spec

```
http://localhost:8080/docs/orders.yaml
```

Swagger supports JWT authorization.

Click **Authorize** and paste:

```
Bearer <jwt_token>
```

---

## Logging

All logs are written to a file:

```
app.log
```

Includes:

* HTTP request logs
* Controller success and error logs
* Authentication failures
* Database errors

The log file is ignored in Git.

---

## Order Model Overview

```
OrderItem
  product_name
  quantity
  unit_price
  total_price

Order
  order_id (UUID)
  items []
  status
  created_at
  updated_at
  disabled (hidden from API)
```




