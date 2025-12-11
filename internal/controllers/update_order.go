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

func UpdateOrder(c *gin.Context) {

	orderID := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var updated models.Order

	if err := c.BindJSON(&updated); err != nil {
		logger.Log.Println("UpdateOrder: Invalid JSON:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid body"})
		return
	}

	// Recalculate total prices
	for i := range updated.Items {
		updated.Items[i].TotalPrice = updated.Items[i].UnitPrice * float64(updated.Items[i].Quantity)
	}

	update := bson.M{
		"$set": bson.M{
			"items":      updated.Items,
			"status":     updated.Status,
			"updated_at": time.Now(),
		},
	}

	collection := database.GetCollection("orders")

	result, err := collection.UpdateOne(ctx, bson.M{
		"order_id": orderID,
		"disabled": false,
	}, update)

	if err != nil {
		logger.Log.Println("UpdateOrder: DB Update error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update order"})
		return
	}

	if result.MatchedCount == 0 {
		logger.Log.Printf("UpdateOrder: OrderID %s not found", orderID)
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		return
	}

	logger.Log.Printf("UpdateOrder: Success. OrderID=%s", orderID)

	c.JSON(http.StatusOK, gin.H{"message": "Order updated successfully"})
}
