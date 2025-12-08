package controllers
import(
	"context"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"order_service/internal/database"
	"order_service/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
)
func GetOrderByID(c *gin.Context){
	//context
	ctx,cancel:= context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	//Extract the ID from URL
	id:= c.Param("id")

	//convert the id string into Mongo ObjectID
	objID, err:= primitive.ObjectIDFromHex(id)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid order ID"})
		return
	}	
	var order models.Order
	collection := database.GetCollection("orders")
	
	//Query filter
	filter:= bson.M{"_id": objID}
	err= collection.FindOne(ctx, filter).Decode(&order)
	if err!=nil{
		c.JSON(http.StatusNotFound,gin.H{"error":"Order Not Found"})
		return
	}
	c.JSON(http.StatusOK, order)
}
