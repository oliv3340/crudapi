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
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"crudapi/docs"
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

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  None

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

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

	// programmatically set swagger info
	docs.SwaggerInfo.Title = "Swagger Album API"
	docs.SwaggerInfo.Description = "This is a sample server Album server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost"
	docs.SwaggerInfo.BasePath = "/v1"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("health", handlers.HealthCheckerHandler())
		album := v1.Group("/albums")
		{
			album.POST("", handlers.CreateAlbumHandler(client))
			album.GET("", handlers.ListAlbumsHandler(client))
			album.GET(":id", handlers.GetAlbumByIDHandler(client))
			album.PATCH(":id", handlers.UpdateAlbumByIDHandler(client))
			album.DELETE(":id", handlers.DeleteAlbumByIDHandler(client))

		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	if err := router.Run("localhost:8085"); err != nil {
		log.Fatalln(err)
	}
}
