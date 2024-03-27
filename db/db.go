package db

import (
	"context"
	"fmt"
	"log"
	"signalzero/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const ( // 172.17.0.1
	DSN        = "mongodb://127.0.0.1:27017/?directConnection=true&serverSelectionTimeoutMS=2000&appName=mongosh+2.2.1"
	DBName     = "signalzero"
	Collection = "users"
)

var userCollection *mongo.Collection

func init() {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(DSN))
	if err != nil {
		log.Fatalln("failed to connect db:", err)
	}

	userCollection = client.Database(DBName).Collection(Collection)
}

func InsertOne(ctx context.Context, args models.User) error {
	res, err := userCollection.InsertOne(ctx, args)
	if err != nil {
		log.Println("failed to insert:", err)
		return err
	}

	fmt.Println("inserted id :", res.InsertedID)
	return nil
}

func FetchUsers(ctx context.Context, query string) ([]models.User, error) {

	filter := bson.D{}
	if query == "" {

		filter = bson.D{}

	} else {
		filter = bson.D{{"username", primitive.Regex{Pattern: query, Options: ""}}}
	}

	cur, err := userCollection.Find(ctx, filter)
	if err != nil {
		log.Println("failed to find:", err)
		return nil, err
	}

	defer cur.Close(ctx)
	users := []models.User{}

	for cur.Next(ctx) {
		result := models.User{}

		err = cur.Decode(&result)
		if err != nil {
			log.Println("failed to decode:", err)
			return nil, err
		}

		users = append(users, result)
	}

	return users, nil
}
