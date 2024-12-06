package main

import (
	"log"
	"net/http"
	"slotmachine/cache"
	"slotmachine/config"
	"slotmachine/handlers"
	"slotmachine/models"
)

func main() {
	mongoURI := config.GetEnv("MONGO_URI", "mongodb://localhost:27017")
	redisAddr := config.GetEnv("REDIS_ADDR", "localhost:6379")

	mongoClient, err := models.NewMongoClient(mongoURI)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v", err)
	}
	cacheClient := cache.NewCache(redisAddr)

	handler := handlers.ModeHandler{
		MongoClient: mongoClient,
		Cache:       cacheClient,
	}

	http.HandleFunc("/popular-modes", handler.PopularModesHandler)
	log.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
