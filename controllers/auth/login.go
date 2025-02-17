package auth

import (
	"ginex/helpers"
	userModel "ginex/models/user"
	"github.com/gin-gonic/gin"
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
	_ = c.Bind(&postData)
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

	//sessionKey ,_ := uuid.NewV4()
	//c.SetCookie("ginex_session",sessionKey.String(),60*60*24,"/","localhost",true,true)
	//session := sessions.Default(c)
	//session.Set(sessionKey.String(),userData.ID)
	//_ = session.Save()

	helpers.SetUserSession(c,userData.ID)

	//sessionKey := "ginex_session_key"
	//c.SetCookie("ginex_session",sessionKey,60*60*24,"/","localhost",true,true)
	//session := sessions.Default(c)
	//session.Set(sessionKey,userData.ID)
	//_ = session.Save()


	c.Redirect(http.StatusMovedPermanently,"/dashboard")
	return
	/*c.JSON(http.StatusOK, gin.H{"code":http.StatusOK,"message":make(map[string]string),"data":userData})
	return*/
}

/*func (LoginController) Set(c *gin.Context)  {
	sessionToken,_ := uuid.NewV4()

	session := sessions.Default(c)
	//session.Set(sessionToken.String(),map[string]string{"name":"cai xu kun"})
	session.Set("wangxiao",map[string]string{"name":"cai xu kun"})
	_ = session.Save()

	c.SetCookie("ginex_session",sessionToken.String(),0,"/","localhost",false,true)
}

func (LoginController) Get(c *gin.Context)  {
	//sessionToken, _ := c.Cookie("ginex_session")

	session := sessions.Default(c)

	value := session.Get("wangxiao")

	c.JSON(http.StatusOK,gin.H{"data":value})
}*/

