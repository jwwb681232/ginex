package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ginex/models"
)

type RegisterForm struct {
	Name                 string `form:"name" json:"name" binding:"required"`
	Email                string `form:"email" json:"email" binding:"required"`
	Password             string `form:"password" json:"password" binding:"required"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" binding:"required"`
}

func ShowRegistrationForm(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/register.html", gin.H{"title": "Register Page"})
}

func Register(c *gin.Context) {
	var data RegisterForm
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func ShowLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/login.html", gin.H{"title": "Login Page"})
}

func Login(c *gin.Context) {
	email := c.PostForm("email")
	password := c.PostForm("password")

	models.GetUser(email)

	c.HTML(http.StatusBadRequest, "auth/postTest.html", gin.H{"title": "Login Post Test Page", "email": email, "password": password})
}

/*func ShowRegistrationFormAll(c *gin.Context) {
	data := models.All()
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
	//c.HTML(http.StatusOK,"auth/register.html",gin.H{"title":title})
}*/
