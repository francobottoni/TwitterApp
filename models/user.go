package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	LastName  string             `bson:"lastName" json:"lastName,omitempty"`
	Birthday  time.Time          `bson:"birthDay" json:"birthDay,omitempty"`
	Email     string             `bson:"email" json:"email,omitempty"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar"`
	Banner    string             `bson:"banner" json:"banner"`
	Biography string             `bson:"biography" json:"biography"`
	Location  string             `bson:"location" json:"location"`
	WebSite   string             `bson:"webSite" json:"webSite"`
}
