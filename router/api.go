package router

import (
	"todolist/router_func"

	"github.com/gin-gonic/gin"
)

func InitApi(r *gin.Engine) {
	api := r.Group("/api")
	api.GET("/lkz", router_func.Lkz)
	// v1 := api.Group("/v1")
}
