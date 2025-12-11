package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Single product inside an order
type OrderItem struct {
	ProductName string  `bson:"product_name" json:"product_name"`
	Quantity    int     `bson:"quantity" json:"quantity"`
	UnitPrice   float64 `bson:"unit_price" json:"unit_price"`
	TotalPrice  float64 `bson:"total_price" json:"total_price"`
}

// Main order structure
type Order struct {
	MongoID primitive.ObjectID `bson:"_id,omitempty" json:"-"`
	OrderID string             `bson:"order_id" json:"order_id"`

	Items  []OrderItem `bson:"items" json:"items"`
	Status string      `bson:"status" json:"status"`

	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	Disabled  bool      `bson:"disabled" json:"-"`
}
