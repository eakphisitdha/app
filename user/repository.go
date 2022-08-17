package user

import (
	"errors"
	"fmt"

	//MongoDB
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var Collection *mongo.Collection
var ConnectionURI = "mongodb://192.168.1.83:1"
var Client, Err = mongo.NewClient(options.Client().ApplyURI(ConnectionURI))

func ConnectDB() {
	//connectionURI := "mongodb://<host_ip>:<host_port>"

	if Err != nil {
		log.Fatal(Err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	Err = Client.Connect(ctx)
	if Err != nil {
		log.Fatal(Err)
	}

	Err = Client.Ping(ctx, readpref.Primary())
	if Err != nil {
		log.Fatal(Err)
	}

	fmt.Println("Connected")
}

func AddUser(newUser Profile) error {
	coll := Client.Database("user").Collection("user")
	doc := bson.D{{"Username", newUser.Username}, {"Phone", newUser.Phone}}
	_, err := coll.InsertOne(context.TODO(), doc)
	if err != nil {
		return err
	}
	//fmt.Printf("Inserted document with _id: %v\n", result.InsertedID)
	return nil
}

func IsExits(phone string) bool {
	coll := Client.Database("user").Collection("user")
	result := coll.FindOne(context.TODO(), bson.D{{"Phone", phone}})
	if errors.Is(result.Err(), mongo.ErrNoDocuments) {
		return false
	}
	return true
}
