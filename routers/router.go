package routers

import (
	"gindemo/controllers"
	"github.com/gin-gonic/gin"
)

func Router() {
	//router := gin.Default()
	////
	////// This handler will match /user/john but will not match /user/ or /user
	////router.GET("/user/:name",)
	////
	////// However, this one will match /user/john/ and also /user/john/send
	////// If no other routers match /user/john, it will redirect to /user/john/
	////router.GET("/user/:name/*action", func(c *gin.Context) {
	////	name := c.Param("name")
	////	action := c.Param("action")
	////	message := name + " is " + action
	////	c.String(http.StatusOK, message)
	////})
	//
	//router.LoadHTMLFiles("./../static/view")
	////加载静态资源，例如网页的css、js
	//router.Static("./../static", "./static")

	r := gin.Default()
	// 加载模板文件
	r.LoadHTMLGlob("templates/*")
	r.Static("static", "./static/")

	//routerG := r.Group("/api")
	//{
	//	routerG.GET("/register", controllers.Register)
	//	routerG.POST("/login", controllers.Login)
	//}
	r.GET("/register", controllers.Register)
	r.POST("/register", controllers.RegisterPostDate)
	r.GET("/login", controllers.Login)
	r.POST("/login", controllers.LoginPostData)

	r.Run(":8080")
}
