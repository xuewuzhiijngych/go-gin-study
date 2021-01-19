package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	setIndexRouter()
}

func setIndexRouter() {
	R.GET("/index", IndexHandler)
}

func IndexHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "IndexHandler",
	})
}
