package auth

import (
	"database/sql"
	"ginex/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"ginex/helpers"
)

type RegisterController struct {}

type RegisterForm struct {
	Name                 string `form:"name" json:"name" binding:"required"`
	Email                string `form:"email" json:"email" binding:"required"`
	Password             string `form:"password" json:"password" binding:"required"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" binding:"required"`
}

func (RegisterController) ShowRegistrationForm(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/register.html", gin.H{"code":http.StatusOK,"message":make(map[string]string)})
}

func (RegisterController) Register(c *gin.Context) {
	var data RegisterForm
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id,err := models.StoreUser(models.UserModel{Name:data.Name,Email:data.Email,Password:data.Password})

	if err != sql.ErrNoRows {
		flash := helpers.Flash(c, map[string]interface{}{
			"code":http.StatusConflict,
			"message":[]string{"The email address already exists"},
		})
		c.HTML(http.StatusConflict,"auth/register.html",flash)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":id,
	})
}

/*func ShowLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/login.html", gin.H{"title": "Login Page"})
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	models.GetUser(email)

	c.HTML(http.StatusBadRequest, "auth/postTest.html", gin.H{"title": "Login Post Test Page", "email": email, "password": password})
}*/

/*func ShowRegistrationFormAll(c *gin.Context) {
	data := models.All()
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
	//c.HTML(http.StatusOK,"auth/register.html",gin.H{"title":title})
}*/
