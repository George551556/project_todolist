package router

import (
	"todolist/router_func"

	"github.com/gin-gonic/gin"
)

func InitUserProfile(r *gin.Engine) {
	userpro := r.Group("/userpro")
	userpro.POST("/getitems", router_func.GetUserItems)
	userpro.POST("/additem", router_func.AddItem)
	userpro.POST("/changenickname", router_func.ChangeNickname)
	userpro.POST("/changeitemcontent", router_func.ChangeItemContent)
	userpro.POST("/deleteitem", router_func.DeleteItem)
	userpro.POST("/changestate", router_func.ChangeState)
}
