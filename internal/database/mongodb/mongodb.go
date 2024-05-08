package mongodb

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client *mongo.Client
	Users  *mongo.Collection
)

func Init(connectionString, dbName string) error {
	//connect to database
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.Background(), opts)
	if err != nil {
		return err
	}

	//create users collection
	Users = client.Database(dbName).Collection("Users")

	//check database connection
	err = client.Database("main").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()
	return err
}

func Close() error {
	//close database connection
	return client.Disconnect(context.Background())
}
