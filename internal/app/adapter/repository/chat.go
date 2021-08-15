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

// Chat is the repository of domain.Chat
type Chat struct {
	collection *mongo.Collection
}

func NewChatRepo(db *mongo.Client) Chat {
	dbName := viper.GetString("mongodb_dbname")
	collection := db.Database(dbName).Collection("chat")
	return Chat{collection: collection}
}

// Create create new chat
func (u Chat) Create(chat *model.Chat) error {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	_, err := u.collection.InsertOne(ctx, chat)
	if err != nil {
		log.Error(err)
		return errors.New("can't create new chat")
	}

	return nil
}

// Get get chat by filter
func (u Chat) Get(f *model.Chat) (*model.Chat, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	out := new(model.Chat)

	err := u.collection.FindOne(ctx, f).Decode(&out)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("chat not found")
		}
		log.Error(err)
		return nil, errors.New("failed get chat")
	}

	return out, nil
}

// Search get chats by users filter
func (u Chat) Search(f *model.Chat) ([]*model.Chat, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)

	var results []*model.Chat

	filter := bson.D{{"users", f.UserIDs[0]}, {"users", f.UserIDs[1]}}

	cur, err := u.collection.Find(ctx, filter)
	if err != nil {
		return nil, errors.New("failed get chats")
	}

	for cur.Next(ctx) {
		var elem model.Chat
		err := cur.Decode(&elem)
		if err != nil {
			return nil, errors.New("failed get chats")
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		return nil, errors.New("failed get chats")
	}

	cur.Close(ctx)

	return results, nil
}
