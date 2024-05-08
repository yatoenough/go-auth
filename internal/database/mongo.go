package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Init(connectionString string) error {
  serverAPI := options.ServerAPI(options.ServerAPIVersion1)
  opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

  localClient, err := mongo.Connect(context.Background(), opts)
  if err != nil {
    return err
  }
  
  client = localClient

  err = client.Database("main").RunCommand(context.TODO(), bson.D{{Key: "ping", Value: 1}}).Err()
  return err
}

func CLose() error{
    return client.Disconnect(context.Background())
}
