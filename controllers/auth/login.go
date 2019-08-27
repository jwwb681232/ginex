package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginController struct {
}

type loginForm struct {
	Email    string `form:"email" validate:"required"`
	Password string `form:"password" validate:"required"`
}

func (LoginController) ShowLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/login.html", gin.H{"code":http.StatusOK,"message":make(map[string]string)})
}

func (LoginController) Login(c *gin.Context) {
	var postData loginForm
	c.Bind(&postData)
	c.JSON(http.StatusOK, gin.H{"code":http.StatusOK,"message":make(map[string]string),"data":postData})
	return
}