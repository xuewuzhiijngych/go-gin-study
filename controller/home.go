package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
	model "yyy/models"
)

func init() {
	setHomeRouter()
}

func setHomeRouter() {
	R.GET("/home", HomeHandler)
}

func HomeHandler(c *gin.Context) {

	var user model.User
	user.Name = "ych"
	user.Age = 12
	user.Birthday = int(time.Now().Unix())
	user.Email = "601266867@qq.com"

	model.UserAdd(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "HomeHandler",
	})
}
