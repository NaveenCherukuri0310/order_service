package routes

import (
	"order_service/internal/controllers"

	"github.com/gin-gonic/gin"
)

func PingRoutes(r *gin.Engine) {
	r.GET("/ping", controllers.Ping)
	r.GET("/test-token", controllers.GetTestToken)
}
