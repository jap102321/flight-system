package query

import "go.mongodb.org/mongo-driver/mongo"

func Flight(db *mongo.Client, collection string) *mongo.Collection {
	var user = db.Database("flightsystem").Collection(collection)
	return user
}
