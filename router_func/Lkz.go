package router_func

import (
	"fmt"
	"todolist/utils"

	"github.com/gin-gonic/gin"
)

func InitDatabase(c *gin.Context) {
	utils.Db_makeTable() //映射数据表，只调用一次
	fmt.Println("数据表初始化成功")
	c.String(200, "hello, 数据表初始化成功")
}
