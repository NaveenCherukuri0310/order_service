package controllers

import (
	"context"
	"net/http"
	"time"

	"order_service/internal/database"
	"order_service/internal/logger"
	"order_service/internal/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetOrderByID(c *gin.Context) {

	orderID := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var order models.Order
	collection := database.GetCollection("orders")

	filter := bson.M{"order_id": orderID, "disabled": false}

	err := collection.FindOne(ctx, filter).Decode(&order)
	if err != nil {
		logger.Log.Printf("GetOrderByID: OrderID %s not found: %v", orderID, err)
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	logger.Log.Printf("GetOrderByID: Success. OrderID=%s", orderID)

	c.JSON(http.StatusOK, order)
}
