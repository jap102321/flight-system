package driver

import (
	"context"
	"time"

	"github.com/jap102321/flight-system/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var app config.GoAppTools

func ConnectToDb(URI string) *mongo.Client {
	ctx, cancelCtx := context.WithTimeout(context.Background(), time.Second)
	defer cancelCtx()

	client, err := mongo.Connect(ctx, options.Client(), options.Client().ApplyURI(URI))
	if err != nil {
		app.ErrorLogger.Panicln(err)
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		app.ErrorLogger.Fatalln(err)
	}

	return client
}
