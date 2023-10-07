package router

import (
	"todolist/router_func"

	"github.com/gin-gonic/gin"
)

func InitApi(r *gin.Engine) {
	api := r.Group("/lkz")
	api.GET("/init", router_func.InitDatabase)
	// v1 := api.Group("/v1")
}
