package router_func

import (
	"fmt"
	"todolist/utils"

	"github.com/gin-gonic/gin"
)

// 映射数据表，只调用一次
func InitDatabase(c *gin.Context) {
	utils.Db_makeTable() //映射数据表，只调用一次
	fmt.Println("数据表初始化成功")
	c.String(200, "hello, 数据表初始化成功")
}

// 用于连接查询两个数据表，显示用户及其对应的事项记录
func FindUsersAddItems(c *gin.Context) {
	res := utils.Db_findUserItems()
	c.JSON(200, res)
}
