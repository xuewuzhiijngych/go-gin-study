package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	model "yyy/models"
)

func init() {
	setIndexRouter()
}

func setIndexRouter() {
	R.GET("/index", IndexHandler)
}

func IndexHandler(c *gin.Context) {
	var articles []model.Article
	model.Db.Find(&articles)

	fmt.Printf("$#s", articles)

	c.HTML(http.StatusOK, "index", gin.H{
		"articles": articles,
	})
}
