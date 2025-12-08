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
	ctx, cancel:= context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	id:=c.Param("id")
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
	result,err :=collection.UpdateByID(ctx,objID,update)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update"})
		return
	}
	if result.MatchedCount==0{
		c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
        return
	}
	updatedData.ID=objID
	c.JSON(http.StatusOK, updatedData)

}