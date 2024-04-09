package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Context struct {
	Client *mongo.Client
	cnx    context.Context
}

func NewMongoClient(conn string, cnx context.Context) *Context {

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)

	clientOptions := options.Client().
		ApplyURI(conn).
		SetServerAPIOptions(serverAPI)

	client, err := mongo.Connect(cnx, clientOptions)

	if err != nil {
		log.Panic(err)
	}

	err = client.Ping(cnx, nil)

	if err != nil {
		log.Panic(err)
	}

	return &Context{Client: client}
}
