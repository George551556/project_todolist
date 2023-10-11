package router_func

import "github.com/gin-gonic/gin"

func Tologinpage(c *gin.Context) {
	c.HTML(200, "loginPage.html", gin.H{})
}
func Tosignuppage(c *gin.Context) {
	c.HTML(200, "signupPage.html", gin.H{})
}
func Tousermainpage(c *gin.Context) {
	c.HTML(200, "userMain.html", gin.H{})
}
func Totestpage(c *gin.Context) {
	c.HTML(200, "test.html", gin.H{})
}
func ToUserPan(c *gin.Context) {
	c.HTML(200, "userPan.html", gin.H{})
}
