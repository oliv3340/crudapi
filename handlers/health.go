package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheckerHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"ok": true})
	}
}
