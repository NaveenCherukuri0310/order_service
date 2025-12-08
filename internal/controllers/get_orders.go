//Function that handle API requests
package controllers
//importing required packages
import(
	//import controls request timeout
	"context"
	//import http status codes
	"net/http"
	"order_service/internal/database"
    "order_service/internal/models"
	//to set timeout duration
    "time"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
)
//handler func for GET/orders
//c *gin.Context is the request context.
func GetOrders(c *gin.Context) {
	//MongoDB requests must have a timeout to prevents the request from hanging forever
	ctx, cancel:= context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	//Opens the MongoDB collection "orders".
	cursor, err := database.GetCollection("orders").Find(ctx, bson.M{})
    if err != nil {
		//gin.H is a shortcut for creating JSON maps:
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
	//Creates an empty list of Order objects.
	var orders []models.Order
	//Moves through the cursor and stores all documents into orders
    if err := cursor.All(ctx, &orders)
	err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

    c.JSON(http.StatusOK, orders)
}

























