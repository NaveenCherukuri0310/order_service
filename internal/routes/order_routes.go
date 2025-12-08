package routes

import (
"github.com/gin-gonic/gin"
"order_service/internal/controllers"
)

func OrderRoutes(router *gin.Engine) {

// All order routes go inside this group
orders := router.Group("/orders")

orders.GET("/", controllers.GetOrders)
orders.POST("/",controllers.CreateOrder)
orders.GET("/:id",controllers.GetOrderByID)


}