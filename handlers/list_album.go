package handlers

import (
	"crudapi/types"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

// ListAlbums godoc
// @Summary      List Albums
// @Description  get all albums
// @Tags         albums
// @Accept       json
// @Produce      json
// @Param        title    query     string  false  "name search by title"  Format(string)
// @Success      200  {array}   types.Album
// @Router       /albums [get]
func ListAlbumsHandler(client *firestore.Client) func(c *gin.Context) {

	return func(c *gin.Context) {
		var result []types.Album
		title := c.Request.URL.Query().Get("title")
		query := client.Collection(types.ALBUM_COLLECTION)
		iter := query.Documents(c)
		if title != "" {
			iter = query.Where("title", "==", title).Documents(c)
		}
		for {
			doc, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				c.JSON(http.StatusInternalServerError, "interal error")
				return
			}
			result = append(result,
				types.Album{
					ID:     doc.Ref.ID,
					Artist: doc.Data()["artist"].(string),
					Title:  doc.Data()["title"].(string),
					Price:  doc.Data()["price"].(float64),
				})

		}
		c.JSON(http.StatusOK, result)
	}
}
