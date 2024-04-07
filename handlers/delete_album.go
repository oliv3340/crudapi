package handlers

import (
	"crudapi/types"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

func DeleteAlbumByIDHandler(client *firestore.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		id := c.Param("id")
		_, err := client.Collection(types.ALBUM_COLLECTION).Doc(id).Delete(c)
		if err != nil {
			c.JSON(http.StatusInternalServerError, "")
			return
		}
		c.JSON(http.StatusOK, "")

	}
}
