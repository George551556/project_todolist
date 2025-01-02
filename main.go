package main

import (
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"time"
	"todolist/router"
	"todolist/utils"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// map记录所有ws客户端
var clients = make(map[int]*websocket.Conn)
var ch = make(chan []byte, 50)
var peopleNums int
var onLineNums int //记录当前已连接人数

// 升级http
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {

	r := gin.Default()
	//加载模板目录下模板文件
	// 使用 CORS 中间件处理跨域请求
	r.Use(cors.Default())
	r.LoadHTMLGlob("templates/*")
	r.Static("static", "static/")
	r.Static("express", "express/")
	utils.Db_makeTable() //初始化数据库连接选项

	peopleNums = 0
	onLineNums = 0

	r.GET("/", rtHTML)
	r.POST("/dingxiang", rRedirect)
	r.GET("/json", rtJSON)
	r.GET("/testdb", utils.TestDB)
	r.GET("/str", str)
	r.POST("strp", str_p)
	r.GET("str_f", str_fortune)

	r.GET("/ws", handleWebSocket)
	go broadcastMsg()
	//导入文件写的路由
	router.InitApi(r)
	router.InitControl(r)
	router.InitAuthen(r)
	router.InitUserProfile(r)

	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

/*以下是一些示例*/
// 直接响应json
func rtJSON(c *gin.Context) {
	c.JSON(200, gin.H{
		"name": "lkz",
		"age":  23,
	})
}

// 响应html
func rtHTML(c *gin.Context) {
	//获取get请求的参数
	// fmt.Println(c.Query("user")) //使用变量名获取参数值

	c.HTML(200, "index.html", gin.H{"time": "疯疯"})
}

// 重定向
func rRedirect(c *gin.Context) {
	//获取post请求的参数值
	// message := c.PostForm("user")   //方法1
	forms, err := c.MultipartForm() //方法2， 接收所有的form参数 注意：使用这个方法要在form表单属性中设置enctype

	fmt.Println("表单多：", forms, err)
	fmt.Println("表单多：", forms)
	// fmt.Println("post信息为：", message)
	c.HTML(200, "test.html", gin.H{"message": "message"})
}
func str(c *gin.Context) {
	current_time := time.Now()
	thistime := current_time.Format("2006-01-02 15:04:05")
	ret := "给你返回一个动态值就当前时间：" + thistime
	c.String(200, ret)
}
func str_p(c *gin.Context) {
	temp := c.PostForm("source")
	ret_msg := "你发送的信息为：" + temp
	fmt.Println("收到消息：", temp)
	c.String(200, ret_msg)
}

// 返回一句古诗
func str_fortune(c *gin.Context) {
	var ret string
	fmt.Println("fortune test")
	instruction := "fortune"
	cmd := exec.Command(instruction)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		ret = "诗词获取失败"
	}
	fmt.Println(string(output))
	ret = string(output)
	c.JSON(200, gin.H{
		"fortune": ret,
	})
}

func handleWebSocket(c *gin.Context) {
	//升级连接
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	//记录一个用户连接
	peopleNums += 1
	onLineNums += 1
	this_peopleNum := peopleNums
	clients[peopleNums] = conn
	fmt.Println("用户", peopleNums, "连接")
	//当前连接人数输入管道，前缀1
	updateNums := "1," + fmt.Sprintf("%d", onLineNums)
	updateNums_1 := []byte(updateNums)
	ch <- updateNums_1
	//监测用户退出
	defer func() {
		conn.Close()
		delete(clients, this_peopleNum)
		onLineNums -= 1
		fmt.Println("用户", this_peopleNum, "退出")
		//当前连接人数输入管道，前缀1
		updateNums := "1," + fmt.Sprintf("%d", onLineNums)
		updateNums_1 := []byte(updateNums)
		ch <- updateNums_1
	}()

	//侦听消息
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		//析取消息,原消息：id, type, content
		decode_msg := string(msg)
		msg_UserId := strings.Split(decode_msg, ",")[0]
		msg_type := strings.Split(decode_msg, ",")[1] //消息类型，0正常消息， 1表情消息
		nickName := utils.Db_getNickName(msg_UserId)  //此处根据msg_UserId查询数据库获取对应的nickname
		msg_Content := strings.Split(decode_msg, ",")[2]
		if msg_type == "1" {
			//发送表情，"2, id, Jack, 1.gif"
			decode_msg_1 := "2," + msg_UserId + "," + nickName + "," + msg_Content + ".gif" //此时，字符串值为 "2,x.gif"
			// fmt.Println("lkz:", decode_msg_1)
			new_msg := []byte(decode_msg_1)
			ch <- new_msg

		} else {
			//发送文字消息
			decode_msg_1 := "0," + msg_UserId + "," + nickName + "," + msg_Content
			new_msg := []byte(decode_msg_1)
			ch <- new_msg
		}
		fmt.Println("接收到", msg_UserId, "号", nickName, "的消息：", msg_Content)
	}
}

// 广播消息，在主函数中为该函数开一个协程即可
func broadcastMsg() {
	for {
		msg, err := <-ch
		if !err {
			break
		}

		//这个协程用于处理单个新消息的广播
		go func() {
			for index, client := range clients {
				err1 := client.WriteMessage(websocket.TextMessage, msg)
				if err1 != nil {
					//如果检测到发送失败则认为客户端掉线，踢出
					delete(clients, index)
				}
			}
		}()
	}
}
