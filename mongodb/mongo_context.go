package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Context struct {
	Client *mongo.Client
}

func NewMongoClient(conn string) *Context {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(conn).
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Panic(err)
	}

	return &Context{Client: client}
}
