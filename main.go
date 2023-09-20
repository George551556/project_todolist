package main

import (
	"mygin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//加载模板目录下模板文件
	r.LoadHTMLGlob("templates/*")
	r.StaticFile("/static/file1", "static/数.txt")

	r.GET("/", rtHTML)
	r.GET("/dingxiang", rRedirect)
	r.GET("/json", rtJSON)

	//导入文件写的路由
	router.InitApi(r)

	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

/*以下是一些示例*/
// 直接响应json
func rtJSON(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "lkz",
		"age":  23,
	})
}

// 响应html
func rtHTML(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{"time": "疯疯"})
}

// 重定向
func rRedirect(c *gin.Context) {
	c.Redirect(301, "/html1")
}
