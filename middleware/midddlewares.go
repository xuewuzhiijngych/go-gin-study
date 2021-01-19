package midddleware

import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

//statsCost 统计请求耗时的中间件
func statsCost() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Set("name", "ych") // 可以通过c.Set在请求上下文中设置值，后续的处理函数能够取到该值
		// 调用该请求的剩余处理程序
		c.Next()
		// 不调用该请求的剩余处理程序
		c.Abort()
		// 计算耗时
		cost := time.Since(start)
		log.Println(cost)
	}
}
