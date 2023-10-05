package router_func

import (
	"fmt"
	"net/http"
	"strconv"
	"todolist/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	temp_username := c.PostForm("username")
	password := c.PostForm("password") //获取输入的账号密码
	username, err := strconv.Atoi(temp_username)
	if err != nil {
		fmt.Println("输入违法")
	}
	fmt.Println("登录请求信息：", username, password)
	authResult, id, nickname := utils.Db_Login_middle(username, password)
	if authResult {
		// c.Status(200) //可以重定向到test页面
		c.JSON(http.StatusOK, gin.H{
			"id":       id,
			"nickname": nickname,
			"status":   http.StatusOK,
		})
	} else {
		// c.Status(400) //如果验证失败可以直接返回400
		c.JSON(400, gin.H{"status": 400})
	}
}

func Signup(c *gin.Context) {

}
