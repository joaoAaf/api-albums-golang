package mongodb

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

func ConnectionDB() *mongo.Collection {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI( /*os.Getenv("STRING_CONNECTION")*/ "mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}
	collection := client.Database("api").Collection("albums")
	return collection
}

func InsertData(obj any) *mongo.InsertOneResult {
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

func FindOne(id primitive.ObjectID) bson.M {
	collection := ConnectionDB()
	var result bson.M
	filter := bson.D{{Key: "_id", Value: id}}
	errFindOne := collection.FindOne(ctx, filter).Decode(&result)
	if errFindOne != nil {
		log.Println("Error", errFindOne)
		return nil
	}
	return result
}
