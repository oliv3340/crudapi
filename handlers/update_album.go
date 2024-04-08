package handlers

import (
	"crudapi/types"
	"log"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

// UpdateAlbumByID godoc
// @Summary      Update an account
// @Description  update an album by id
// @Tags         album
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  types.Album
// @Router       /albums/{id} [patch]
func UpdateAlbumByIDHandler(client *firestore.Client) func(c *gin.Context) {
	return func(c *gin.Context) {
		var album types.Album
		album.ID = c.Param("id")

		if err := c.BindJSON(&album); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		_, err := client.Collection(types.ALBUM_COLLECTION).
			Doc(album.ID).Set(c, map[string]interface{}{
			"artist": album.Artist,
			"title":  album.Title,
			"price":  album.Price,
		}, firestore.MergeAll)

		if err != nil {
			log.Printf("error creating document: %s", err)
			c.JSON(http.StatusInternalServerError, "")
			return
		}
		c.JSON(http.StatusOK, album)

	}
}
