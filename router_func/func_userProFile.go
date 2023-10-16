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
	if utils.Db_createOneItem(userid, content) {
		fmt.Println("添加事项:", userid, "内容：", content)
		c.Status(200)
	} else {
		c.Status(400)
	}

}

// 路由函数：根据传入的userid以及content修改用户的昵称
func ChangeNickname(c *gin.Context) {
	temp_userid := c.PostForm("userid")
	userid, err := strconv.Atoi(temp_userid)
	if err != nil {
		fmt.Println(err)
	}
	new_nickname := c.PostForm("content")
	if utils.Db_change(userid, new_nickname) {
		c.Status(200)
	} else {
		c.Status(400)
	}
}

// 路由函数：根据传入的itemId以及content修改事项内容
func ChangeItemContent(c *gin.Context) {
	temp_itemId := c.PostForm("itemId")
	itemId, err := strconv.Atoi(temp_itemId)
	if err != nil {
		fmt.Println(err)
	}
	new_content := c.PostForm("new_content")
	if utils.Db_changeItemContent(itemId, new_content) {
		c.Status(200)
	} else {
		c.Status(400)
	}
}

// 路由函数：根据传入的itemId删除item
func DeleteItem(c *gin.Context) {
	temp_itemId := c.PostForm("itemId")
	itemId, err := strconv.Atoi(temp_itemId)
	if err != nil {
		fmt.Println(err)
	}
	if utils.Db_deleteItem(itemId) {
		c.Status(200)
	} else {
		c.Status(400)
	}
}

// 路由函数：根据传入的itemId完成item
func ChangeState(c *gin.Context) {
	temp_itemId := c.PostForm("itemId")
	itemId, err := strconv.Atoi(temp_itemId)
	if err != nil {
		fmt.Println(err)
	}
	if utils.Db_changeItemState(itemId) {
		c.Status(200)
	} else {
		c.Status(400)
	}
}

// 路由函数：根据传入的文件以及用户id存储文件并存储文件信息
func UploadFile(c *gin.Context) {
	temp_userid := c.PostForm("userid")
	userid, err := strconv.Atoi(temp_userid)
	if err != nil {
		fmt.Println(err)
	}
	file, file_err := c.FormFile("file")

	if file_err != nil {
		fmt.Println("上传文件为空")
		c.JSON(400, gin.H{"status": 400, "msg": "请选择文件"})
	}
	fmt.Println("文件保存信息：", userid, file.Filename)

	//文件存储在本地
	rootDir := "files/"
	file_site := rootDir + temp_userid + "/" //文件的完整存放地址
	fileStorage := file_site + file.Filename //文件保存最终路径
	c.SaveUploadedFile(file, fileStorage)

	//文件及用户信息保存在数据库
	result := utils.Db_storeFiles(userid, file.Filename, file_site)
	//返回信息
	c.JSON(200, gin.H{"status": 200, "msg": result})
}

// 路由函数：根据传入的userid和file_id返回下载对应的文件
func DownloadFile(c *gin.Context) {
	temp_userid := c.PostForm("userid")
	temp_file_id := c.PostForm("file_id")
	userid, err := strconv.Atoi(temp_userid)
	if err != nil {
		fmt.Println(err)
	}
	file_id, err1 := strconv.Atoi(temp_file_id)
	if err1 != nil {
		fmt.Println(err1)
	}

	file_info := utils.Db_getFileInfo(userid, file_id)
	// dir := "files/8/"
	// filename := "数.txt"
	dir := file_info.FileSite
	filename := file_info.FileName

	// c.Header("Content-Disposition", "attachment: filename="+filename)
	c.File(dir + filename)
}

// 路由函数：根据传入的userid和file_id删除文件信息和文件
func DeleteFile(c *gin.Context) {
	temp_userid := c.PostForm("userid")
	temp_file_id := c.PostForm("file_id")
	userid, err := strconv.Atoi(temp_userid)
	if err != nil {
		fmt.Println(err)
	}
	file_id, err1 := strconv.Atoi(temp_file_id)
	if err1 != nil {
		fmt.Println(err1)
	}

	//删除数据库信息 tag为TRUE或FALSE表示删除是否成功；temp是被删除的信息的结构实例，可以为后面删除实际文件提供信息
	tag, temp := utils.Db_deleteFileInfo(userid, file_id)
	//删除实际文件

	if tag {
		fmt.Println("删除一条文件信息：", temp)
		c.Status(200)
	} else {
		fmt.Println("文件删除失败...")
		c.Status(400)
	}
}

// 根据传入的userid返回该用户拥有的文件信息
func GetFileItems(c *gin.Context) {
	temp_userid := c.PostForm("userid")
	userid, err := strconv.Atoi(temp_userid)
	if err != nil {
		fmt.Println(err)
	}
	results := utils.Db_findFiles(userid)
	c.JSON(200, gin.H{"items": results})
}
