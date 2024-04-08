package handlers

import (
	"crudapi/types"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

// DeleteAlbumByID godoc
// @Summary      Delete an account
// @Description  delete an album
// @Tags         album
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Account ID"
// @Success      200  {object}  types.Album
// @Router       /albums/{id} [delete]
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
