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



#### 改进说明：

-整体：需要改进使用jwt登录验证，目前只是使用浏览器的localStorage来保存用户ID

-网盘功能：界面ui需要改进，改成那种一个个方块的形式，并添加鼠标悬停效果；删除一个文件是只删除数据库中信息，并没有的目录中删除文件本体。

