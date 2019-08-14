package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ginex/models"
)

func ShowRegistrationForm(c *gin.Context) {
	models.Get()
	c.HTML(http.StatusOK,"auth/register.html",gin.H{"title":"Register Page"})
}

func ShowLoginForm(c *gin.Context) {
	c.HTML(http.StatusOK,"auth/login.html",gin.H{"title":"Login Page"})
}

func ShowRegistrationFormAll(c *gin.Context) {
	data := models.All()
	c.JSON(http.StatusOK,gin.H{
		"data":data,
	})
	//c.HTML(http.StatusOK,"auth/register.html",gin.H{"title":title})
}