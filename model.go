package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Users struct {
	ID   primitive.ObjectID `json:"id,omitempty"`
	Name string             `json:"name,omitempty" validate:"required"`
	Age  string             `json:"age,omitempty" validate:"required"`
}
