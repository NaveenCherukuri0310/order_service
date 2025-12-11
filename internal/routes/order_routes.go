package routes

import (
	"order_service/internal/controllers"
	"order_service/internal/middleware"

	"github.com/gin-gonic/gin"
)

func OrderRoutes(r *gin.Engine) {

	order := r.Group("/orders")
	order.Use(middleware.AuthRequired())

	order.POST("", controllers.CreateOrder)
	order.GET("", controllers.GetOrders)
	order.GET("/:id", controllers.GetOrderByID)
	order.PUT("/:id", controllers.UpdateOrder)
}
