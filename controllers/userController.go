package controllers

import (
	"fmt"
	"gindemo/models"
	"github.com/gin-gonic/gin"
	"log"
)

func Register(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func RegisterPostDate(c *gin.Context) {
	name := c.PostForm("userName")
	pwd := c.PostForm("password")
	email := c.PostForm("email")
	code := c.PostForm("code")
	if isSucc := models.InserUserData(name, pwd, email, code); isSucc {
		log.Fatalln("插入成功")
	}
}
func Login(c *gin.Context) {
	c.HTML(200, "login.html", nil)

}
func LoginPostData(c *gin.Context) {
	fmt.Println(c.PostForm("userName"))
	fmt.Println(c.PostForm("password"))

}
