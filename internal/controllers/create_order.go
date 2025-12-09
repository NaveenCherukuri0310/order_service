package controllers
import(
	"context"
	"log"
	"time"
	"net/http"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"order_service/internal/database"
	"order_service/internal/models"
)
func CreateOrder(c *gin.Context){
	
	//MongoDB requests must have a timeout to prevents the request from hanging forever
	ctx, cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	var order models.Order

	// Bind JSON body to Order struct
	if err:= c.BindJSON(&order);err!=nil{
		log.Println("CreateOrder: Invalid JSON body:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Assign a new MongoDB ObjectID
	order.ID=primitive.NewObjectID()
	collection := database.GetCollection("orders")

	result, err := collection.InsertOne(ctx, order)
	if err != nil {
		log.Println("CreateOrder: MongoDB InsertOne failed:", err)
    	c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert order"})
    	return
	}

	//success log
	log.Printf("CreateOrder: Order created with ID %v", result.InsertedID)
	c.JSON(http.StatusCreated, order)

}
