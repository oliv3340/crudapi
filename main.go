package main

import (
	"context"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx := context.Background()
	client, err := initFirebaseClient(ctx)
	if err != nil {
		log.Printf("error creating firestore client: %v", err)
		return
	}
	collection := client.Collection("albums").NewDoc()
	log.Printf("set value in collection")
	result, err := collection.Set(ctx, map[string]interface{}{
		"ID": "1", "Title": "Blue Train", "Artist": "John Coltrane", "Price": 56.99,
	})

	if err != nil {
		log.Printf("error creating default value: %v", err)
	}

	log.Printf("Result: %v", result)
}
