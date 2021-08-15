package repository

import (
	"context"
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/orensimple/trade-chat-app/internal/app/adapter/mongodb/model"
	"github.com/prometheus/common/log"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
)

// Message is the repository of domain.Message
type Message struct {
	collection *mongo.Collection
}

func NewMessageRepo(db *mongo.Client) Message {
	dbName := viper.GetString("mongodb_dbname")
	collection := db.Database(dbName).Collection("message")
	return Message{collection: collection}
}

// Create create new message
func (u Message) Create(chat *model.Message) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := u.collection.InsertOne(ctx, chat)
	if err != nil {
		log.Error(err)
		return errors.New("can't create new message")
	}

	return nil
}

// Get get message by filter
func (u Message) Get(f *model.Message) (*model.Message, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	out := new(model.Message)

	err := u.collection.FindOne(ctx, bson.D{{"_id", f.ID}}).Decode(&out)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("message not found")
		}
		log.Error(err)
		return nil, errors.New("failed get message")
	}

	return out, nil
}

// Search get messages by users filter
func (u Message) Search(f *model.Message) ([]*model.Message, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var results []*model.Message

	cur, err := u.collection.Find(ctx, f)
	if err != nil {
		return nil, errors.New("failed get messages")
	}

	for cur.Next(ctx) {
		var elem model.Message
		err := cur.Decode(&elem)
		if err != nil {
			return nil, errors.New("failed get messages")
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, errors.New("failed get messages")
	}

	cur.Close(ctx)

	return results, nil
}
