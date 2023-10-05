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

// 根据id查密码，用于登录验证，返回验证状态以及用户ID及昵称
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

// 事项item相关数据库操作函数====================================================================
// 增加一条事项记录
func Db_createOneItem(userid int, content string) bool {
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	tempItem := Items{UserId: userid, Content: content, CreateTime: time.Now(), State: false}
	err := db.Create(&tempItem).Error
	fmt.Println("添加事项执行完.....")
	return err == nil
}

// 根据userid查找某个用户所有对应事项记录
func Db_findItems(userid int) []Items {
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	var tempItems []Items
	db.Where(&Items{UserId: userid}).Find(&tempItems) // 也可以使用 db.Where("UserId = ?", userId).Find(&items)
	fmt.Println("查找执行完毕")
	// for i := range tempItems {
	// 	fmt.Println(i, "行：", tempItems[i])
	// }
	return tempItems
}

// 修改某事项的工作内容
func Db_changeItemContent(id int, newContent string) bool {
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	tempItem := Items{Id: id}
	this_err := db.Model(&tempItem).Update("content", newContent).Error
	return this_err == nil
}

// 根据事项id删除某事项
func Db_deleteItem(id int) bool {
	db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	var tempItem Items
	db.Where(&Items{Id: id}).Find(&tempItem)
	this_err := db.Delete(&tempItem).Error
	return this_err == nil
}
