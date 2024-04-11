package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type MongoConn struct {
	Client *mongo.Client
}

func NewMongoClient(cnx context.Context, conn string) *MongoConn {

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

	return &MongoConn{Client: client}
}

func (c *MongoConn) GetCollection(db string, collectionName string) *mongo.Collection {
	return c.Client.Database(db).Collection(collectionName)
}
