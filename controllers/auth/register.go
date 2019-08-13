package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ginex/models"
)

func ShowRegistrationForm(c *gin.Context) {
	models.Get()
	c.HTML(http.StatusOK,"auth/register.html",gin.H{"title":"注册页"})
}

func ShowRegistrationFormAll(c *gin.Context) {
	data := models.All()
	c.JSON(http.StatusOK,gin.H{
		"data":data,
	})
	//c.HTML(http.StatusOK,"auth/register.html",gin.H{"title":title})
}