// 测试使用mongodb数据库，暂时不用这个了，改为使用mysql XXXXXXXXXXXXXXXXXXXXXXXXXX
package utils

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type user struct {
	id       string `bson:"id"`
	password string `bson:"password"`
	nickname string `bson:"nickname"`
}

// func initDB() {
// 	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")
// 	//连接到mongodb
// 	client, err := mongo.Connect(context.TODO(), clientOptions)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//检查连接
// 	err = client.Ping(context.TODO(), nil)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println("conn is ok")
// }

func TestDB(c *gin.Context) {
	clientOptions := options.Client().ApplyURI("mongodb://127.0.0.1:27017/")
	//连接到mongodb
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	//检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("conn is ok")

	//开始查询
	collection := client.Database("todolist").Collection("events")
	var users []*user
	result, err := collection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	//遍历查询结果
	for result.Next(context.TODO()) {
		var tmp_user user
		err := result.Decode(&tmp_user)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, &tmp_user)
	}

	//检查游标错误
	// if err := result.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	//关闭游标
	result.Close(context.TODO())

	//关闭连接
	err = client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("close connection")

	//打印查询结果
	for _, user := range users {
		fmt.Println("ok:", user.id, user.password, user.nickname)
		fmt.Println("ooo")
	}

}
