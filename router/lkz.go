package router

import (
	"todolist/router_func"

	"github.com/gin-gonic/gin"
)

// 此处放置我的后门路由

func InitApi(r *gin.Engine) {
	api := r.Group("/lkz")
	api.GET("/init", router_func.InitDatabase)
	api.GET("/find", router_func.FindUsersAddItems)
	// v1 := api.Group("/v1")
}
