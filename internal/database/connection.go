package database
import(
	"context"
	"log"
	"os"
	"time"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
var Client *mongo.Client
func ConnectDB() *mongo.Client{
	//Read from .env(MONGO_URI)
	uri:= os.Getenv("MONGO_URI")
	if uri==""{
		log.Fatal("Mising MONGO_URI")
	}
	ctx, cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()

	client, err:=mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err!= nil{
		log.Fatal("Failed to connect MongoDB:",err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
    	log.Fatal("MongoDB ping failed:", err)
	}

	log.Println("Connected to MongoDB")
	Client=client
	return client
}
func GetCollection(collectionName string) *mongo.Collection {
	//read DB_NAME from .env
	dbName:=os.Getenv("DB_Name")
	if dbName==""{
		log.Fatal("Missing DB_NAME in .env file")
	}
	return Client.Database(dbName).Collection(collectionName)
}