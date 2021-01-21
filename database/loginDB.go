package database

import (
	"context"
	"time"

	"github.com/francobottoni/TwitterApp/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Insert a new User in db
func InsertUser(u models.User) (string, bool, error, models.User) {

	//Here set Timeout
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConection.Database("twittor")
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err, models.User{}
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)
	return ObjID.String(), true, nil, u
}
