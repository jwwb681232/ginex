package auth

import (
	"fmt"
	"ginex/helpers"
	userModel "ginex/models/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
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

func (LoginController) Set(c *gin.Context)  {
	sessionToken,_ := uuid.NewV4()
	//d6db0562-42c1-4bb2-b73e-988357fa0e6d
	session := sessions.Default(c)
	session.Set(sessionToken.String(),map[string]string{"name":"cai xu kun"})
	session.Save()//important

	c.SetCookie("ginex_session",sessionToken.String(),0,"/","localhost",false,true)
}

func (LoginController) Get(c *gin.Context)  {
	sessionToken, _ := c.Cookie("ginex_session")
	fmt.Println(sessionToken)
	session := sessions.Default(c)

	value := session.Get(sessionToken)
	/*for k,v :=range c.Request.Header {
		fmt.Println(k,v)
	}*/
	c.JSON(http.StatusOK,value)
}

