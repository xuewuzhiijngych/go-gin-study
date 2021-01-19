package router

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
	"path/filepath"
)

//loadTemplates 模板继承
func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()
	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}
	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}
	for _, include := range includes {
		layouCopy := make([]string, len(layouts))
		copy(layouCopy, layouts)
		files := append(layouCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}

// SetupRouter 设置路由
func SetupRouter() *gin.Engine {
	// 默认路由引擎
	r := gin.Default()
	// 自定义模板函数
	r.SetFuncMap(template.FuncMap{
		"safe": func(str string) template.HTML {
			return template.HTML(str)
		},
	})
	// 定义 templatesDir
	r.HTMLRender = loadTemplates("./templates")
	//静态文件处理
	r.Static("/static", "templates/static")
	// 模板渲染
	r.LoadHTMLGlob("templates/**/*")
	// 404
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404temp", nil)
	})

	r.GET("/hello", HelloHandler)
	return r
}

// 默认入口
func HelloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello Gin",
	})
}
