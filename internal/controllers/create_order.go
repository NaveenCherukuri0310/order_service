package controllers
import(
	"context"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"order_service/internal/database"
	"order_service/internal/models"
)
func CreateOrder(c *gin.Context){
	ctx, cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	var order models.Order
	// Bind JSON body to Order struct
	if err:= c.BindJSON(&order);err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Assign a new MongoDB ObjectID
	order.ID=primitive.NewObjectID()
	collection := database.GetCollection("orders")

	_, err := collection.InsertOne(ctx, order)
	if err != nil {
    	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert order"})
    	return
	}

	c.JSON(http.StatusCreated, order)

}
