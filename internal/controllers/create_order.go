package controllers

import (
	"context"
	"net/http"
	"time"

	"order_service/internal/database"
	"order_service/internal/logger"
	"order_service/internal/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateOrder(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var order models.Order

	// Parse JSON
	if err := c.BindJSON(&order); err != nil {
		logger.Log.Println("CreateOrder: Invalid JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Core fields
	order.MongoID = primitive.NewObjectID()
	order.OrderID = uuid.New().String()
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.Disabled = false

	// Auto-calc total price per item
	for i := range order.Items {
		order.Items[i].TotalPrice = order.Items[i].UnitPrice * float64(order.Items[i].Quantity)
	}

	collection := database.GetCollection("orders")

	_, err := collection.InsertOne(ctx, order)
	if err != nil {
		logger.Log.Println("CreateOrder: Insert failed:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
		return
	}

	logger.Log.Printf("CreateOrder: Success. OrderID=%s", order.OrderID)

	c.JSON(http.StatusCreated, order)
}
