package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	setHomeRouter()
}

func setHomeRouter() {
	R.GET("/home", HomeHandler)
}

func HomeHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "HomeHandler",
	})
}
