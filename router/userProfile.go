package router

import (
	"todolist/router_func"

	"github.com/gin-gonic/gin"
)

func InitUserProfile(r *gin.Engine) {
	userpro := r.Group("/userpro")
	userpro.POST("/getitems", router_func.GetUserItems)
	userpro.POST("/additem", router_func.AddItem)
}
