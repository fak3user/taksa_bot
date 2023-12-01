package db

import (
	"context"
	"log"
	"taksa/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func GetClientOptions() *options.ClientOptions {
	envs, err := utils.GetEnvs()
	if err != nil {
		log.Fatal(err)
	}

	dburi := envs.DbUri

	serverAPIOptions := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().
		ApplyURI(dburi).
		SetServerAPIOptions(serverAPIOptions)

	return clientOptions
}

func GetCollection(collection string) *mongo.Collection {
	client := GetMongoClient()

	return client.Database("taksa").Collection(collection)
}

func Init() {
	clientOptions := GetClientOptions()

	newClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	} else {
		client = newClient
	}
}

func GetMongoClient() mongo.Client {
	return *client
}
