package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/prometheus/common/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"

	"github.com/spf13/viper"
)

// Connection gets connection of mongodb database
func Connection() (db *mongo.Client) {
	host := viper.Get("mongodb_host")
	port := viper.Get("mongodb_port")
	dsn := fmt.Sprintf("mongodb://%v:%v", host, port)

	log.Warnf(dsn)

	db, err := mongo.NewClient(options.Client().ApplyURI(dsn))
	if err != nil {
		log.Error(err)
	}

	// Connect the mongo client to the MongoDB server
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = db.Connect(ctx)

	// Ping MongoDB
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	if err = db.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("could not ping to mongo db service: %v\n", err)
		return
	}

	return db
}
