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

func GetOrders(c *gin.Context) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var orders []models.Order

	collection := database.GetCollection("orders")

	filter := bson.M{"disabled": false}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		logger.Log.Println("GetOrders: DB Find error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch orders"})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var o models.Order
		if err := cursor.Decode(&o); err != nil {
			logger.Log.Println("GetOrders: Decode error:", err)
			continue
		}
		orders = append(orders, o)
	}

	logger.Log.Printf("GetOrders: Success. Count=%d", len(orders))

	c.JSON(http.StatusOK, orders)
}
