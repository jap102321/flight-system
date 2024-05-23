package main

import (
	"context"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jap102321/flight-system/config"
	"github.com/jap102321/flight-system/driver"
	"github.com/joho/godotenv"
)

var app config.GoAppTools

func main() {
	InfoLogger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	ErrorLogger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)

	app.InfoLogger = *InfoLogger
	app.ErrorLogger = *ErrorLogger

	err := godotenv.Load()
	if err != nil {
		app.ErrorLogger.Fatal("No .env file available")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		app.ErrorLogger.Fatalln("mongodb_uri not found")
	}

	client := driver.ConnectToDb(uri)

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			app.ErrorLogger.Fatal(err)
			return
		}
	}()

	appRouter := gin.New()
	appRouter.GET("/", func(ctx *gin.Context) {
		app.InfoLogger.Println("Creating a scalable web application with Gin")
	})

	err = appRouter.Run()
	if err != nil {
		app.ErrorLogger.Fatal(err)
	}
}
