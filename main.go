package main

import (
	"context"
	"crudapi/handlers"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func initFirebaseClient(ctx context.Context) (*firestore.Client, error) {
	conf := &firebase.Config{
		ProjectID: os.Getenv("GCLOUD_PROJECT_ID"),
	}

	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		log.Printf("error, initializing app : %v", err)
		return nil, err
	}
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Printf("error, initializing client : %v", err)
		return nil, err
	}

	return client, err
}

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

	router := gin.Default()

	router.GET("/api/health", handlers.HealthCheckerHandler())

	router.POST("/album", handlers.CreateAlbumHandler(client))
	router.GET("/albums", handlers.ListAlbumsHandler(client))
	router.GET("/album/:id", handlers.GetAlbumByIDHandler(client))
	router.PATCH("/album/:id", handlers.UpdateAlbumByIDHandler(client))
	router.DELETE("/album/:id", handlers.DeleteAlbumByIDHandler(client))

	router.Run("localhost:8081")
}
