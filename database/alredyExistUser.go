package database

import (
	"context"
	"time"

	"github.com/francobottoni/TwitterApp/models"
	"go.mongodb.org/mongo-driver/bson"
)

//Validate if user alredy exist in the db
func IfUserAlreadyExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoConection.Database("Twittor")
	col := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := col.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
