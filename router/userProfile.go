package router

import (
	"todolist/router_func"

	"github.com/gin-gonic/gin"
)

func InitUserProfile(r *gin.Engine) {
	userpro := r.Group("/userpro")

	//主页相关路由
	userpro.POST("/getitems", router_func.GetUserItems)
	userpro.POST("/additem", router_func.AddItem)
	userpro.POST("/changenickname", router_func.ChangeNickname)
	userpro.POST("/changeitemcontent", router_func.ChangeItemContent)
	userpro.POST("/deleteitem", router_func.DeleteItem)
	userpro.POST("/changestate", router_func.ChangeState)

	//网盘相关路由
	userpro.POST("/uploadfile", router_func.UploadFile)
}
