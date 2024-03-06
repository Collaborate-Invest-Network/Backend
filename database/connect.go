package database

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const dbName = "cin"
const mongoURI = "mongodb://localhost:27017/" + dbName

func ConnectMongo() error {
	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal("Error connecting to mongodb", err)
	}

	db := client.Database(dbName)

	Mi = MongoInstance{
		Client: client,
		Db:     db,
	}

	log.Printf("Mongodb connected (%v)", dbName)

	return nil
}
