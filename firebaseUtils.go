package main

import (
	"context"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
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
