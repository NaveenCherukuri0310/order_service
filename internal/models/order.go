package models
import "go.mongodb.org/mongo-driver/bson/primitive"

type Order struct {
    ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Item     string             `json:"item" bson:"item"`
    Quantity int                `json:"quantity" bson:"quantity"`
    Price    float64            `json:"price" bson:"price"`
    Status   string             `json:"status" bson:"status"`
}
