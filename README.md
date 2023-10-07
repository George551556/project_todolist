#### gin框架模板

1. router下面放所有路由，要分组编写

2. router_func下面放对应路由的中间件函数

3. static 存放图片、文件等资源

4. templates 存放html模板文件

5. 必要命令：
   go mod init [模块名] 
   go get -u github.com/mongodb/mongo-go-driver   安装mongodb驱动
   go get -u github.com/gin-gonic/gin
   

   最后使用 go mod tidy 整理包
6. 注意：添加新路由一定要在main.go中注册路由
   第一次部署项目要访问路由 /lkz/init 进行初始化数据表



#### 本项目使用说明：

1.使用MySQL数据库 创建一个todolist数据库，不用创建数据表。

2.在云服务器上部署需要更改main.go端口为80

3.启动项目后，初次先访问路由

```
/lkz/init
```

   来初始化数据表



