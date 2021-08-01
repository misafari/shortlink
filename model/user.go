package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserId    int32              `bson:"user_id"`
	Name      string             `bson:"name"`
	Surname   string             `bson:"surname"`
	CreatedAt time.Time          `bson:"created_at"`
}
