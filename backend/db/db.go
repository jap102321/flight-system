package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jap102321/flight-system/config"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client mongo.Client
var app config.GoAppTools
var DB *mongo.Database

func ConnectToDb(URI string) (*mongo.Client, error) {
    ctx, cancelCtx := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancelCtx()

    clientOptions := options.Client().ApplyURI(URI)
    client, err := mongo.Connect(ctx, clientOptions)
    if err != nil {
        return nil, err
    }

    err = client.Ping(ctx, nil)
    if err != nil {
        return nil, err
    }

    return client, nil
}

func InitDatabase() {
	err := godotenv.Load()
	if err != nil {
		app.ErrorLogger.Fatal("No .env file available")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		app.ErrorLogger.Fatalln("mongodb_uri not found")
	}

	client, err  := ConnectToDb(uri)

	if err != nil{
		panic(err)
	}


	DB = client.Database("flightsystem")
	log.Println("Connected to MongoDB!")
}