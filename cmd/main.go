//Go executable must have package main.
package main

import(
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"order_service/internal/database"
	"order_service/internal/routes"
	"github.com/joho/godotenv"
)

func main()  {
	//Load .env file
	if err:=godotenv.Load();err!=nil{
		log.Fatal("Error loading .env file")
	}
	log.Println("Loaded DB_NAME:", os.Getenv("DB_NAME"))

	//Check required env variables
	required:=[]string{"MONGO_URI","DB_NAME","PORT"}

	for _,key:=range required{
		if os.Getenv(key)==""{
			log.Fatalf("Environment variable %s is missing in .env",key)
		}
	}

	//Database connection
	database.ConnectDB()
	router:=gin.Default()
	router.GET("/ping", func(c *gin.Context){
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	routes.OrderRoutes(router)
	//This starts the web server
	port:= os.Getenv("PORT")
	router.Run(":"+port)
	
}