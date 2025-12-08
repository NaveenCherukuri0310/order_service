package routes
import(
	"order_service/internal/controllers"
	"github.com/gin-gonic/gin"
)
func PingRoutes(router *gin.Engine){
	router.GET("/ping", controllers.Ping)
}