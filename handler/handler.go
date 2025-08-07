package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetHealth(c *gin.Context) {
	c.Header("Content-Type","application/json")
	c.JSON(http.StatusOK, gin.H{
		"message": "Health check is OK Ji!",
	})
}
