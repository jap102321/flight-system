package query

import (
	"github.com/jap102321/flight-system/config"
	"go.mongodb.org/mongo-driver/mongo"
)

type GoAppDB struct {
	App config.GoAppTools
	DB  *mongo.Client
}

//	func NewGoAppDB(app config.GoAppTools, db *mongo.Client) database.DBRepo {
//		return &GoAppDB{
//			App: app,
//			DB:  db,
//		}
//	}
func (g *GoAppDB) InsertUser() {
	return
}
