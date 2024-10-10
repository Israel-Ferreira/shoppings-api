package databases

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client

func CreateMongoConn(host, port, username, password string) error {
	mongoConnStr := fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnStr))

	if err != nil {
		return err
	}

	if err = client.Ping(ctx, nil); err != nil {
		return err
	}

	Client = client

	return err
}

func GetCollection(dbConn *mongo.Client, database, collectionName string) *mongo.Collection {
	collection := dbConn.Database(database).Collection(collectionName)
	return collection
}
