package router_func

import (
	"fmt"
	"strconv"
	"todolist/utils"

	"github.com/gin-gonic/gin"
)

// 路由函数：根据id查询并返回该用户的所有事项数据
func GetUserItems(c *gin.Context) {
	temp_userid := c.PostForm("userid")
	userid, err := strconv.Atoi(temp_userid)
	if err != nil {
		fmt.Println(err)
	}
	items := utils.Db_findItems(userid)
	//对时间格式进行解析优化
	// for i := range items {
	// 	temp_time := items[i].CreateTime
	// 	this_time, err1 := time.Parse(time.RFC3339, temp_time)
	// 	if err1 != nil {
	// 		fmt.Println("one of the Createtime parse failed...")
	// 	}
	// 	res_time := this_time.Format("2006-01-02 15:04")
	// 	items[i].CreateTime = res_time
	// }
	//返回数据
	c.JSON(200, gin.H{
		"items": items,
	})
}

// 路由函数：根据传入的userid以及content添加新的事项
func AddItem(c *gin.Context) {
	temp_userid := c.PostForm("userid")
	userid, err := strconv.Atoi(temp_userid)
	if err != nil {
		fmt.Println(err)
	}
	content := c.PostForm("content")
	utils.Db_createOneItem(userid, content)
	fmt.Println("添加事项:", userid, "内容：", content)
	c.Status(200)
}
