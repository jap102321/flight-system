package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jap102321/flight-system/config"
	"github.com/jap102321/flight-system/db"
)

var app config.GoAppTools

func main() {
	db.InitDatabase()
	appRouter := gin.Default()
	Routes(appRouter)

	err := appRouter.Run()
	if err != nil {
		app.ErrorLogger.Fatal(err)
	}
}
