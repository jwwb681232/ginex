package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ginex/helpers"
	userModel "ginex/models/user"
	"golang.org/x/crypto/bcrypt"
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
	validateMessage, err := helpers.Validate(&postData)
	if err != nil {
		c.HTML(http.StatusBadRequest, "auth/login.html", gin.H{"code": http.StatusFound, "message": validateMessage, "data": postData})
		return
	}

	userData, notFound := userModel.User{}.WhereEmail(&postData.Email)
	if notFound {
		c.HTML(http.StatusBadRequest, "auth/login.html", gin.H{"code": http.StatusFound, "message": map[string]string{"loginForm.Email": "该账号不存在"}, "data": postData} )
		return
	}

	if err = bcrypt.CompareHashAndPassword([]byte(userData.Password),[]byte(postData.Password)); err != nil {
		c.HTML(http.StatusBadRequest, "auth/login.html", gin.H{"code": http.StatusFound, "message": map[string]string{"loginForm.Password": "密码错误"}, "data": postData} )
		return
	}

	c.JSON(http.StatusOK, gin.H{"code":http.StatusOK,"message":make(map[string]string),"data":userData})
	return
}