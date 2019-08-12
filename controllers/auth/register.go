package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ShowRegistrationForm(c *gin.Context) {
	c.HTML(http.StatusOK,"auth/register.html",gin.H{"title":"注册页"})
}