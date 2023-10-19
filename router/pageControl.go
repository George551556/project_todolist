package router

import (
	"todolist/router_func"

	"github.com/gin-gonic/gin"
)

func InitControl(r *gin.Engine) {
	control := r.Group("/control")
	control.GET("/tologinpage", router_func.Tologinpage)
	control.GET("/tosignuppage", router_func.Tosignuppage)
	control.GET("/tousermainpage", router_func.Tousermainpage)
	control.GET("/totestpage", router_func.Totestpage)
	control.GET("/touserpanpage", router_func.ToUserPan)
	control.GET("/tochatpage", router_func.ToChanPage)
}
