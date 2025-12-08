package controllers
import(
	"net/http"
	"context"
	"time"
	"order_service/internal/database"
	"order_service/internal/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)
func UpdateOrder(c *gin.Context){
	
	//MongoDB requests must have a timeout to prevents the request from hanging forever
	ctx, cancel:= context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	//Extract the ID from URL
	id:=c.Param("id")

	//convert the id string into Mongo ObjectID
	objID,err:= primitive.ObjectIDFromHex(id)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":"Invalid ID"})
		return
	}
	var updatedData models.Order
	
	//Bind JSON body
	if err := c.BindJSON(&updatedData); err != nil {
    	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    	return
	}
	collection:=database.GetCollection("orders")
	update:=bson.M{
		"$set": bson.M{
			"item":     updatedData.Item,
        	"quantity": updatedData.Quantity,
        	"price":    updatedData.Price,
        	"status":   updatedData.Status,
		},
	}
	_,err =collection.UpdateByID(ctx,objID,update)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update"})
		return
	}
	
	updatedData.ID=objID
	c.JSON(http.StatusOK, updatedData)

}