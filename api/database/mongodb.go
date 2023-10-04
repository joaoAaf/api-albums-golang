package mongodb

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx, cancel = context.WithTimeout(context.Background(), 20*time.Second)

func ConnectionDB() *mongo.Collection {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("STRING_CONNECTION")))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}
	collection := client.Database("api").Collection("albums")
	return collection
}

func InsertData(obj interface{}) *mongo.InsertOneResult {
	collection := ConnectionDB()
	result, err := collection.InsertOne(ctx, obj)
	if err != nil {
		panic(err)
	}
	return result
}

func FindAll() []bson.M {
	var results []bson.M

	collection := ConnectionDB()

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		panic(err)
	}
	if err = cur.All(ctx, &results); err != nil {
		panic(err)
	}
	return results
}

func FindOne() {

}
