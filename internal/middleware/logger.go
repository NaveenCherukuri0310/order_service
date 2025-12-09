package middleware
import(
	"log"
	"time"
	"github.com/gin-gonic/gin"
)
func Logger() gin.HandlerFunc{
	return func(c *gin.Context){

		//records the current time the request enters middleware.
		start := time.Now()

		//Continue to next handler
		c.Next()
		duration:= time.Since(start)

		log.Printf("%s %s | Status: %d | Time: %v",
			c.Request.Method,
			c.Request.URL.Path,
			c.Writer.Status(),
			duration,
		)
	}
}