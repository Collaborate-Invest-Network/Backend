package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id", json:"_id"`
	FirstName string             `bson:"firstname", json:"firstname"`
	LastName  string             `bson:"lastname", json:"lastname"`
	Username  string             `bson:"username", json:"username"`
	Password  string             `bson:"password", json:"-"`
	Email     string             `bson:"email", json:"email"`
	Birthday  time.Time          `bson:"birthday", json:"birthday"`
	Address   string             `bson:"address", json:"address"`
	Phone     string             `bson:"phone", json:"phone"`
	IsActive  bool               `bson:"isAvtive", json:"isActive"`
	CreatedAt time.Time          `bson:"createdAt", json:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt", json:"updatedAt"`
}
