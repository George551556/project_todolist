package router

//身份认证相关路由
import (
	"todolist/router_func"

	"github.com/gin-gonic/gin"
)

func InitAuthen(r *gin.Engine) {
	authen := r.Group("/authen")
	authen.POST("/login", router_func.Login)
	authen.POST("/signup", router_func.Signup)

}
