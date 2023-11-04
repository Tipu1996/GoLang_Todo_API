package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	Item      string         `json:"item" bson:"item"`
	Completed bool           `json:"completed" bson:"completed"`
}
