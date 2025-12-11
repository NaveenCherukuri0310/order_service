# Order Service API

A production ready Order Management API built using Go, Gin, and MongoDB.
This service demonstrates clean architecture, environment variable usage, custom logging, and full CRUD operations.

## Features

* Create a new order
* Get all orders
* Get order by ID
* Update an order
* Delete an order
* Custom request logger middleware
* Logging implemented in CreateOrder handler
* Central logger available for all files (optional to use)
* MongoDB connection with timeout
* Clean folder structure
* Environment variable validation
* Dedicated ping route for health check

## Project Structure

```
order_service
│
├── cmd
│   └── main.go
│
├── internal
│   ├── controllers
│   │     ├── ping_controller.go
│   │     ├── create_order.go  ← logging added here
│   │     ├── get_orders.go
│   │     ├── get_order_by_id.go
│   │     └── update_order.go
│   │
│   ├── routes
│   │     ├── ping_routes.go
│   │     └── order_routes.go
│   │
│   ├── database
│   │     └── connection.go
│   │
│   ├── models
│   │     └── order.go
│
├── middleware
│   └── logger.go   
│
│
├── sample.env
├── .gitignore
├── go.mod
├── go.sum
└── README.md
```

## Request Logging

### Request Logger Middleware

Every HTTP request is logged automatically through:

```
middleware/logger.go
```

Logs include
* Method
* Path
* Status
* Duration

This runs for *every route* without needing to modify controllers.

### Application Logging

You can log from any file using:

```go
logger.Log.Println("message here")
```

### Current Implementation

At this moment, internal logging is added only inside:

```
create_order.go
```

Other controllers do not use logging yet, but support is ready.

## Requirements

* Go installed
* MongoDB (local or Atlas)
* Postman or Thunder Client

## Setup Instructions

### 1. Clone

```
git clone <repo-url>
cd order_service
```

### 2. Create `.env`

```
cp sample.env .env
```

Edit:

```
MONGO_URI=mongodb://localhost:27017
DB_NAME=orderdb
PORT=8080
```

### 3. Install dependencies

```
go mod tidy
```

### 4. Run

```
go run cmd/main.go
```

## API Endpoints

### Health Check

```
GET /ping
```

### Create Order

```
POST /orders
```

Body:

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

## Environment Variables

```
MONGO_URI=
DB_NAME=
PORT=
```


