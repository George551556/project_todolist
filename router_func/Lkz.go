package router_func

import "github.com/gin-gonic/gin"

func Lkz(c *gin.Context) {
	c.String(200, "hello, 分组路由，调用的路由为\"Lkz\"")
}
