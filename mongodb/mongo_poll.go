package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"time"
)

type MongoPool struct {
	pool        chan *mongo.Client
	timeout     time.Duration
	uri         string
	connections int
	poolSize    int
}

func (mp *MongoPool) getContextTimeOut(cnx context.Context) context.Context {
	ctx, _ := context.WithTimeout(cnx, mp.timeout)
	return ctx
}

func (mp *MongoPool) createToChan(cnx context.Context) {
	client, err := mongo.Connect(mp.getContextTimeOut(cnx),
		options.Client().ApplyURI(mp.uri).SetServerAPIOptions(options.ServerAPI(options.ServerAPIVersion1)))

	if err != nil {
		log.Fatalf("Create the Pool failed，err=%v", err)
	}

	mp.pool <- client
	mp.connections++
}

func (mp *MongoPool) CloseConnection(cnx context.Context, conn *mongo.Client) error {
	select {
	case mp.pool <- conn:
		return nil
	default:
		if err := conn.Disconnect(cnx); err != nil {
			log.Fatalf("Close the Pool failed，err=%v", err)
			return err
		}
		mp.connections--
		return nil
	}
}

func (mp *MongoPool) GetConnection(cnx context.Context) (*mongo.Client, error) {
	for {
		select {
		case conn := <-mp.pool:
			err := conn.Ping(mp.getContextTimeOut(cnx), readpref.Primary())
			if err != nil {
				log.Fatalf("err=%v", err)
				return nil, err
			}
			return conn, nil
		default:
			if mp.connections < mp.poolSize {
				mp.createToChan(cnx)
			}
		}
	}
}

func GetCollection(conn *mongo.Client, dbname, collection string) *mongo.Collection {
	return conn.Database(dbname).Collection(collection)
}
