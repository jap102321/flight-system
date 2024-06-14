package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jap102321/flight-system/config"
	"github.com/jap102321/flight-system/db"
)

var app config.GoAppTools

func main() {
	db.InitDatabase()
	appRouter := gin.Default()
	config := cors.Config{
        AllowOrigins:     []string{"http://localhost:3000"}, // Reemplaza con el origen de tu frontend
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        ExposeHeaders:    []string{"Content-Length"},
        AllowCredentials: true,
        MaxAge:           12 * time.Hour,
    }
	appRouter.Use(cors.New(config))
	
	Routes(appRouter)
	err := appRouter.Run()
	if err != nil {
		app.ErrorLogger.Fatal(err)
	}
}
