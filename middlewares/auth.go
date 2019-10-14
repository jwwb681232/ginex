package middlewares

import (
	"ginex/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		var value models.User
		sessionValue := session.Get("ginex_session_key")

		if sessionValue == nil  {
			c.Abort()
			c.Redirect(http.StatusMovedPermanently,"/login")
			return
		}

		value = sessionValue.(models.User)
		if value.ID != 0  {
			c.Next()
			return
		}

		c.Abort()
		c.Redirect(http.StatusMovedPermanently,"/login")
		return
	}
}