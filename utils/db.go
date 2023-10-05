package utils

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//使用gorm方式，连接mysql数据库
/*
各个数据表如下：
user : {Id int, Username string(账号), Password string, Nickname string, SignUpTime time.time}
items : {Id int, UserId int, Content string, CreateTime time.time, State int} //State意为该事件的状态，比如紧急、普通

*/
//指定用户表
type User struct {
	Id         int  `gorm:"primaryKey"`
	Username   uint `gorm:"primaryKey;type:integer"`
	Password   string
	Nickname   string
	SignUpTime time.Time
}

// 指定事项表Items
type Items struct {
	Id         int
	UserId     int `gorm:"foreignKey:UserId;references:User(Id)"`
	Content    string
	CreateTime time.Time
	State      bool
}

// 指定表名
func (User) TableName() string {
	return "users"
}
func (Items) TableName() string {
	return "items"
}

var db *gorm.DB
var err error
var dsn = "root:liu314314@tcp(127.0.0.1:3306)/todolist?charsetutf8mb4&parseTime=true&loc=Local"

// 创建表，只调用一次
func Db_makeTable() {
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&Items{})
	fmt.Println("tables make succs")
}

// 增加一个成员
func Db_createOneUser(username uint, password string, nickname string) {
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	tempUser := User{
		Username:   username,
		Password:   password,
		Nickname:   nickname,
		SignUpTime: time.Now(),
	}
	db.Create(&tempUser)
}

// 查询成员
func Db_findUsers() {
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	var userList []User
	db.Find(&userList)
	for i := range userList {
		fmt.Println(i, "行:", userList[i])
	}
}

// 修改成员信息
func Db_change() {
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	tempUser := User{Id: 3} //示例，按照id修改对应信息
	db.Model(&tempUser).Update("Nickname", "完美世界")
}

// 删除(注销)用户
func db_delete(id int, password string) bool {
	//需要核对密码是否正确
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	tempUser := User{Id: id}
	db.Where(&User{Id: id}).Find(tempUser)
	// tempUser := User{Id: 4}
	if tempUser.Password == password {
		db.Delete(&tempUser)
		return true
	} else {
		return false
	}
}

// 根据id查密码，用于登录验证，返回验证状态以及用户昵称
func Db_Login_middle(username int, password string) (bool, int, string) {
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	var tempUser User
	db.Where(&User{Username: uint(username)}).Find(&tempUser)
	if tempUser.Password == password {
		fmt.Println("用户", tempUser.Nickname, "密码验证通过")
		return true, tempUser.Id, tempUser.Nickname
	} else {
		fmt.Println("用户", tempUser.Nickname, "验证失败。。。")
		return false, -1, ""
	}
}
