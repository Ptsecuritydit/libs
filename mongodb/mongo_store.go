package mongodb

import (
	"context"
	"errors"
	"fmt"
	"github.com/Ptsecuritydit/libs/mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type MongoStore struct {
	Session *Context
}

func (receiver *MongoStore) GetUidPersonFromKey(key string, value string, dataBase string, tab string) (string, error) {

	var personKey models.Person

	filter := bson.D{{key, value}}
	collection := receiver.Session.Client.Database(dataBase).Collection(tab)

	err := collection.FindOne(context.TODO(), filter).Decode(&personKey)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", err
		}
		log.Panic(err)
	}
	return personKey.PersonId, nil
}

func (receiver *MongoStore) InsertOrUpdateItem(personKey models.Person, dataBase string, tab string) {

	collection := receiver.Session.Client.Database(dataBase).Collection(tab)

	filter := bson.D{{"person_id", personKey.PersonId}}
	update := bson.M{
		"$set": bson.M{
			"login_id":    personKey.DomainId,
			"email":       personKey.Email,
			"telegram":    personKey.Telegram,
			"phone":       personKey.Phone,
			"telegram_id": personKey.TelegramId,
		},
	}

	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		log.Panic(err)
	}

	if updateResult.MatchedCount == 0 {
		_, err := collection.InsertOne(context.TODO(), personKey)
		if err != nil {
			log.Panic(err)
		}
	}
}

func (receiver *MongoStore) UpdateItem(key string, value string, dataBase string, tab string) error {
	collection := receiver.Session.Client.Database(dataBase).Collection(tab)
	filter := bson.D{{key, value}}
	update := bson.M{
		"$set": bson.M{
			"telegram":    key,
			"telegram_id": key,
		},
	}
	updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Panic(err)
	}
	if updateResult.MatchedCount == 0 {
		return fmt.Errorf("key not found: %v", value)
	}
	return nil
}
