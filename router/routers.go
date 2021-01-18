package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/hello", HelloHandler)
	return r
}
