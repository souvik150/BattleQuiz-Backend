package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"github.com/souvik150/BattleQuiz-Backend/config"
	"github.com/souvik150/BattleQuiz-Backend/database"
	"github.com/souvik150/BattleQuiz-Backend/middlewares"
	"github.com/souvik150/BattleQuiz-Backend/routes"
)

func main() {
	r := gin.New()
	r.Use(middlewares.Logger())
	r.Use(gin.Recovery())

	config.LoadConfig()
	log.Println("Configuration loaded successfully")


	// Connect to the database
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	} else {
		log.Println("Connected to the database successfully")
	}

	// Connect to Redis
	if err := database.ConnectRedis(); err != nil {
		log.Fatalf("Failed to connect to Redis: %v", err)
	} else {
		log.Println("Connected to Redis successfully")
	}

	routes.RegisterRoutes(r)
	log.Println("Routes registered successfully")

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	} else {
		log.Println("Server started successfully on port 8080")
	}
}
