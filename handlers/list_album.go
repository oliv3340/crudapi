package handlers

import (
	"crudapi/types"
	"net/http"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/iterator"
)

func ListAlbumsHandler(client *firestore.Client) func(c *gin.Context) {

	return func(c *gin.Context) {
		var result []types.Album
		iter := client.Collection(types.ALBUM_COLLECTION).Documents(c)
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
