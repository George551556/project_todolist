package router

// import (
// 	"mygin/web"

// 	"github.com/gin-gonic/gin"
// )

// func InitCourse(r *gin.Engine) {
// 	course := r.Group("/course")
// 	v1 := course.Group("v1")
// 	v1.GET("/detail/:id", web.GetCourseDetail)
// 	v1.GET("/view/:id", web.GetCourseVideo)

// 	admin := course.Group("/admin")
// 	adminV1 := admin.Group("/v1")
// 	adminV1.POST("/add", web.AddCourse)
// 	adminV1.POST("/publish", web.PublishCourse)
// 	adminV1.POST("/upload", web.UploadCourse)

// }
