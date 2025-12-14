// Go executable must have package main.
package main

import (
	"log"
	"order_service/internal/database"
	"order_service/internal/logger"
	"order_service/internal/middleware"
	"order_service/internal/routes"
	"os"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Init()
	//Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Loaded DB_NAME:", os.Getenv("DB_NAME"))

	//Check required env variables
	required := []string{"MONGO_URI", "DB_NAME", "PORT"}

	for _, key := range required {
		if os.Getenv(key) == "" {
			log.Fatalf("Environment variable %s is missing in .env", key)
		}
	}

	//Database connection
	database.ConnectDB()

	// Gin engine
	router := gin.New()
	router.Use(gin.Recovery())

	//Custom Logger
	router.Use(middleware.Logger())

	// Serve OpenAPI spec
	router.StaticFile("/docs/orders.yaml", "./docs/orders.yaml")

	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(
		swaggerFiles.Handler,
		ginSwagger.URL("/docs/orders.yaml"),
	))

	routes.PingRoutes(router)
	routes.OrderRoutes(router)

	//This starts the web server
	port := os.Getenv("PORT")
	log.Println("Server running on port:", port)
	router.Run(":" + port)

}
