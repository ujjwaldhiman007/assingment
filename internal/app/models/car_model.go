package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Car struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	
	// Completed bool               `json:"completed"`
	   Name string             `json:"name" bson:"name"`
	
}
