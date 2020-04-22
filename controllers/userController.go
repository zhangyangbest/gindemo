package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	c.HTML(200, "register.html", nil)
}

func Login(c *gin.Context) {

	fmt.Println(c.PostForm("userName"))
	fmt.Println(c.PostForm("password"))

}
