package utils

import (
	"fmt"
	"strconv"
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

// 指定网盘文件表
type PanFile struct {
	Id         int
	UserId     int `gorm:"foreignKey:UserId;references:User(Id)"`
	FileName   string
	FileSite   string
	CreateTime time.Time
}

// 指定分享链接存储表，其中usertimes每生成一次链接就+1，每下载一次文件就-1
type FileShare struct {
	Id       int
	Hash     string
	UserId   int
	NickName string
	FileId   int
	FileName string
	UseTimes int
}

// 用于后门查询的结构
type Result struct {
	Username   uint
	Nickname   string
	Content    string
	CreateTime time.Time
}

// 指定表名
func (User) TableName() string {
	return "users"
}
func (Items) TableName() string {
	return "items"
}
func (PanFile) TableName() string {
	return "panfiles"
}
func (FileShare) TableName() string {
	return "fileshares"
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
	//迁移表
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Items{})
	db.AutoMigrate(&PanFile{})
	db.AutoMigrate(&FileShare{})

	//连接数
	sqlDB, err1 := db.DB()
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(80)
	sqlDB.SetConnMaxLifetime(time.Second * 8)
	if err1 != nil {
		panic("连接数据库失败")
	}

	fmt.Println("tables make succs")
}

// 增加一个成员
func Db_createOneUser(username uint, password string, nickname string) bool {
	// db, err = gorm.Open(mysql.Open(dsn))
	// if err != nil {
	// 	panic(err)
	// }
	tempUser := User{
		Username:   username,
		Password:   password,
		Nickname:   nickname,
		SignUpTime: time.Now(),
	}
	//判断该username是否已经存在
	var check_user []User
	db.Where(&User{Username: username}).Find(&check_user)
	/*
		where的查询结果为：
			查不到：err1为空
			查到
	*/
	if len(check_user) < 1 {
		fmt.Println("查询结果数组为0")
		db.Create(&tempUser)
		return true
	} else {
		return false
	}

}

// 查询成员
func Db_findUsers() {
	// db, err = gorm.Open(mysql.Open(dsn))
	// if err != nil {
	// 	panic(err)
	// }
	var userList []User
	db.Find(&userList)
	for i := range userList {
		fmt.Println(i, "行:", userList[i])
	}
}

// 修改成员昵称
func Db_change(id int, new_nickname string) bool {
	// db, err = gorm.Open(mysql.Open(dsn))
	// if err != nil {
	// 	panic(err)
	// }
	tempUser := User{Id: id} //示例，按照id修改对应信息
	err1 := db.Model(&tempUser).Update("Nickname", new_nickname).Error
	fmt.Println("执行了修改昵称，修改为", new_nickname)
	if err1 != nil {
		return false
	} else {
		return true
	}
}

// 删除(注销)用户
func db_delete(id int, password string) bool {
	//需要核对密码是否正确
	// db, err = gorm.Open(mysql.Open(dsn))
	// if err != nil {
	// 	panic(err)
	// }
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
	// db, err = gorm.Open(mysql.Open(dsn))
	// if err != nil {
	// 	panic(err)
	// }
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

// 根据ID查昵称，返回昵称字符串
func Db_getNickName(userid string) string {
	// db, err = gorm.Open(mysql.Open(dsn))
	// if err != nil {
	// 	panic(err)
	// }
	userId, err1 := strconv.Atoi(userid)
	if err != nil {
		fmt.Println("转换失败", err1)
		return "错昵称"
	}
	var resultUser User
	db.Where(&User{Id: userId}).Find(&resultUser)
	return resultUser.Nickname
}

// ===============================================================================================================
// 事项item相关数据库操作函数====================================================================
// 增加一条事项记录
func Db_createOneItem(userid int, content string) bool {
	// db, err = gorm.Open(mysql.Open(dsn))
	// if err != nil {
	// 	panic(err)
	// }

	//查询该用户拥有的todo条数
	var todoCount int64
	var useless []Items
	// db.Model(&Items{UserId: userid, deleted_at}).Count(&todoCount)
	db.Where(&Items{UserId: userid}).Find(&useless).Count(&todoCount)
	if todoCount >= 100 {
		fmt.Println("用户", userid, "超限...", todoCount)
		return false
	}

	tempItem := Items{UserId: userid, Content: content, CreateTime: time.Now(), State: false}
	err := db.Create(&tempItem).Error
	fmt.Println("添加事项执行完.....")
	return err == nil
}

// 根据userid查找某个用户所有对应事项记录
func Db_findItems(userid int) []Items {
	// db, err = gorm.Open(mysql.Open(dsn))
	// if err != nil {
	// 	panic(err)
	// }
	var tempItems []Items
	db.Where(&Items{UserId: userid}).Order("create_time desc").Find(&tempItems) // 也可以使用db.Where("UserId = ?", userId).Find(&items)
	fmt.Println("查找执行完毕")
	// for i := range tempItems {
	// 	fmt.Println(i, "行：", tempItems[i])
	// }
	return tempItems
}

// 修改某事项的工作内容，同时更新其修改（创建）时间
func Db_changeItemContent(id int, newContent string) bool {
	// db, err = gorm.Open(mysql.Open(dsn))
	// if err != nil {
	// 	panic(err)
	// }
	tempItem := Items{Id: id}
	myUpdates := map[string]interface{}{
		"Content":    newContent,
		"CreateTime": time.Now(),
	}
	// this_err := db.Model(&tempItem).Update("content", newContent).Error
	this_err := db.Model(&tempItem).Updates(myUpdates).Error
	fmt.Println("执行修改事项内容函数")
	return this_err == nil
}

// 根据事项id删除某事项
func Db_deleteItem(id int) bool {
	// db, err = gorm.Open(mysql.Open(dsn))
	// if err != nil {
	// 	panic(err)
	// }
	var tempItem Items
	db.Where(&Items{Id: id}).Find(&tempItem)
	this_err := db.Delete(&tempItem).Error
	return this_err == nil
}

// 根据事项id完成某事项
func Db_changeItemState(id int) bool {
	var resItem Items
	tempItem := Items{Id: id}
	//获取原本完成状态
	db.Where(&tempItem).Find(&resItem)
	tag := resItem.State
	tag = !tag
	// if tag{
	// 	resItem.State = false
	// }else{
	// 	resItem.State = true
	// }
	this_err := db.Model(&tempItem).Update("State", tag).Error
	fmt.Println("执行修改事项状态函数")
	return this_err == nil
}

// 根据关键字查询事项包含的事项
func Db_findContentInclude(userid int, fil string) []Items {
	var result []Items
	filter := "%" + fil + "%"
	db.Where("user_id = ? and Content like ?", userid, filter).Order("create_time desc").Find(&result)
	return result
}

//==================================================================================================

// 函数：用于连接查询两个数据表，显示用户及其对应的事项记录
func Db_findUserItems() []Result {
	// db, err = gorm.Open(mysql.Open(dsn))
	// if err != nil {
	// 	panic(err)
	// }
	var results []Result
	fmt.Println("调用lkz查询...")
	db.Table("users").
		Select("users.username, users.nickname, items.content, items.create_time").
		Joins("JOIN items ON users.id = items.user_id").
		Find(&results)

	return results
}

// ===============================================================================================================
// 网盘相关函数====================================================================
// 根据user_id和文件名将文件信息存放在PanFile表中
func Db_storeFiles(user_id int, file_name string, file_site string) string {
	// db, err = gorm.Open(mysql.Open(dsn))
	// if err != nil {
	// 	panic(err)
	// }
	//先查询数据库中有无同名文件，若有则直接修改其创建时间
	var check []PanFile
	db.Where(&PanFile{FileName: file_name}).Find(&check)
	if len(check) < 1 {
		//无同名文件
		thisFile := PanFile{UserId: user_id, FileName: file_name, FileSite: file_site, CreateTime: time.Now()}
		db.Create(&thisFile)
		fmt.Println("文件保存成功")
		return "文件保存成功"
	} else {
		db.Model(&PanFile{Id: check[0].Id}).Update("CreateTime", time.Now())
		fmt.Println("成功覆盖已有文件")
		return "成功覆盖已有文件"
	}
}

// 根据user_id返回用户的文件信息
func Db_findFiles(userid int) []PanFile {
	// db, err = gorm.Open(mysql.Open(dsn))
	// if err != nil {
	// 	panic(err)
	// }
	var results []PanFile
	db.Where(&PanFile{UserId: userid}).Find(&results)
	return results
}

// 根据userid和file_id返回对应文件的存放目录及文件名
func Db_getFileInfo(userid int, file_id int) PanFile {
	var result PanFile
	db.Where(&PanFile{UserId: userid, Id: file_id}).Find(&result)
	return result
}

// 根据userid和file_id删除一条文件信息
func Db_deleteFileInfo(userid int, file_id int) (bool, PanFile) {
	// db, err = gorm.Open(mysql.Open(dsn))
	// if err != nil {
	// 	panic(err)
	// }

	var temp PanFile
	db.Where(&PanFile{Id: file_id, UserId: userid}).Find(&temp)
	err1 := db.Delete(&temp).Error
	return err1 == nil, temp
}

//================================================================================================
//网盘文件的分享链接表相关函数

// 增加一个链接
func Db_fileShareAdd(hash string, userid int, fileid int) {
	str_userid := strconv.Itoa(userid)
	nickname := Db_getNickName(str_userid)
	filename := Db_getFileInfo(userid, fileid).FileName
	//若原本存在该链接则次数加1，否则创建一个数据项
	var temp []FileShare
	db.Where(&FileShare{Hash: hash}).Find(&temp)
	if len(temp) == 1 {
		db.Model(&temp[0]).Update("UseTimes", temp[0].UseTimes+1)
		fmt.Println(filename, "文件增加一个次数")
		return
	}
	//添加一个项
	new_link := FileShare{Hash: hash, UserId: userid, NickName: nickname, FileId: fileid, FileName: filename, UseTimes: 1}
	db.Create(&new_link)
	fmt.Println("新增加一个文件分享链接:", filename)
}

// 通过hash后缀返回该条分享链接对应的nickname和filename
func Db_getLinkInfo(hash string) FileShare {
	var temp FileShare
	db.Where(&FileShare{Hash: hash}).Find(&temp)
	return temp
}

// 通过hash减少一次分享文件的下载次数
func Db_reduceDownloadTime(hash string) {
	var temp FileShare
	db.Where(&FileShare{Hash: hash}).Find(&temp)
	newTimes := temp.UseTimes - 1
	err := db.Model(&FileShare{}).Where(&FileShare{Hash: hash}).Update("UseTimes", newTimes).Error
	if err != nil {
		fmt.Println("lkz error")
		fmt.Println(err)
	}
}
