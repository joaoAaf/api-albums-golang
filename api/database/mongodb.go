package mongodb

import (
	"api/config"
	album "api/model"
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var ctx = context.Background()

func connectionDB() *mongo.Collection {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.StringConnectionMongo()))
	if err != nil {
		panic(err)
	}
	if err := client.Ping(ctx, nil); err != nil {
		panic(err)
	}
	collection := client.Database("api").Collection("albums")
	return collection
}

func InsertData(newAlbum album.Album) *mongo.InsertOneResult {
	collection := connectionDB()
	result, err := collection.InsertOne(ctx, newAlbum)
	if err != nil {
		panic(err)
	}
	return result
}

func FindAll() []bson.M {
	var results []bson.M

	collection := connectionDB()

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
	collection := connectionDB()
	var result bson.M
	filter := bson.D{{Key: "_id", Value: id}}
	errFindOne := collection.FindOne(ctx, filter).Decode(&result)
	if errFindOne != nil {
		log.Println("Error", errFindOne)
	}
	return result
}

func UpdateData(id primitive.ObjectID, newAlbum album.Album) bson.M {
	collection := connectionDB()
	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: newAlbum}}
	var result bson.M
	errUpdate := collection.FindOneAndUpdate(ctx, filter, update).Decode(&result)
	if errUpdate != nil {
		log.Println("Error", errUpdate)
	}
	return result
}
