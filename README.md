#### 目录

[TOC]



#### gin框架模板

1. router下面放所有路由，要分组编写

2. router_func下面放对应路由的中间件函数

3. static 存放图片、文件等资源

4. templates 存放html模板文件

5. 必要命令：
   go mod init [模块名] 
   go get -u github.com/mongodb/mongo-go-driver   安装mongodb驱动（本项目未使用）
   go get -u github.com/gin-gonic/gin（如果安装gin失败：[修改代理](https://blog.csdn.net/asd1358355022/article/details/128397188)） 
   
   最后使用 go mod tidy 整理包
   
   
   
6. 注意：添加新路由一定要在main.go中注册路由
   第一次部署项目要访问路由 /lkz/init 进行初始化数据表



#### 本项目使用说明：

1.使用MySQL数据库 创建一个todolist数据库，不用创建数据表。（注意对照db.go文件中MySQL的连接语句dsn的连接用户和密码以及对应的服务器地址或端口）

2.在云服务器上部署需要更改main.go文件中端口为80 (已经添加了linuxGen分支，可以直接切换分支)

​	进行socket连接时的IP地址改为服务器的IP地址（main文件和chatRoom.html文件中的Ip_port的值）

3.后门路由：

```
查询所有用户的所有items
/lkz/find
```



#### 改进说明：

-整体：需要改进使用jwt登录验证，目前只是使用浏览器的localStorage来保存用户ID

-网盘功能：界面ui需要改进，改成那种一个个方块的形式，并添加鼠标悬停效果；删除一个文件是只删除数据库中信息，并没有的目录中删除文件本体。

#### linux安装go教程

[ubuntu go卸载,安装以及版本升级_删除apt安装的go-CSDN博客](https://blog.csdn.net/qq_36389107/article/details/107972274#:~:text=ubuntu go卸载%2C安装以及版本升级 1 1.检查是否安装GO 2 2.卸载旧版本,(不存在可以忽略这一步) 3 3.下载新的安装包 4 4.更新 5 5.设置GOSUMDB) 

注意要看好自己的是x86还是arm，我在这里出问题一起没发现。

使用系统指令查看为x86-64，但是使用apt install golang-go 安装上的竟然是arm版本，所以决定选择arm版本的安装包就好了
