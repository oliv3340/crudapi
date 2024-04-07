package handlers

import (
	"crudapi/types"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func GetAlbumByIDHandler(client *firestore.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		dsnap, err := client.Collection(types.ALBUM_COLLECTION).Doc(id).Get(c)
		if err != nil {

			c.JSON(http.StatusInternalServerError, "")
		}
		doc := dsnap.Data()
		album := &types.Album{
			ID:     id,
			Title:  doc["title"].(string),
			Artist: doc["artist"].(string),
			Price:  doc["price"].(float64),
		}
		c.JSON(http.StatusOK, album)
	}
}
