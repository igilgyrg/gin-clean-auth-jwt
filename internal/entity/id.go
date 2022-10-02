package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ID string

func NewID() ID {
	return ID(primitive.NewObjectID().Hex())
}
