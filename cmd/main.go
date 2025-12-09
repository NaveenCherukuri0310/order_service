//Go executable must have package main.
package main

import(
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"order_service/internal/database"
	"order_service/internal/routes"
	"order_service/internal/middleware"
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
	
	// Gin engine
	router := gin.New()      
	router.Use(gin.Recovery())
	
	//Custom Logger
	router.Use(middleware.Logger())

	routes.PingRoutes(router)
	routes.OrderRoutes(router)

	//This starts the web server
	port:= os.Getenv("PORT")
	log.Println("Server running on port:", port)
	router.Run(":"+port)
	
}