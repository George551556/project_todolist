gin框架模板
1. router下面放所有路由，要分组编写

2. router_func下面放对应路由的中间件函数

3. static 存放图片、文件等资源

4. templates 存放html模板文件

5. 必要命令：
   go mod init [模块名] 
   go get -u github.com/mongodb/mongo-go-driver   安装mongodb驱动
   go get -u github.com/gin-gonic/gin

   最后使用 go mod tidy 整理包
