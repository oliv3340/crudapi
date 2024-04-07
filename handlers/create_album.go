package handlers

import (
	"crudapi/types"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func CreateAlbumHandler(client *firestore.Client) func(c *gin.Context) {

	return func(c *gin.Context) {
		var album types.Album
		if err := c.BindJSON(&album); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		ref := client.Collection(types.ALBUM_COLLECTION).NewDoc()
		_, err := ref.Set(c, map[string]interface{}{
			"artist": album.Artist,
			"title":  album.Title,
			"price":  album.Price,
		})
		if err != nil {
			log.Printf("error creating document: %s", err)
			c.JSON(http.StatusInternalServerError, "")
			return
		}
		album.ID = ref.ID
		c.JSON(http.StatusCreated, album)
	}
}
