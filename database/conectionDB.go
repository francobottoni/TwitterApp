package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConection = ConectDB()
var clientOptions = options.Client().ApplyURI("mongodb+srv://twitter:trades123@cluster0.jplgb.mongodb.net/twittor?retryWrites=true&w=majority")

func ConectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}

	//Validate conection run OK
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conection ok with your DB!")
	return client
}

func CheckConection() bool {
	err := MongoConection.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}

	return true
}
