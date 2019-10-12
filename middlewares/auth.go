package middlewares

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		value := session.Get("ginex_session_key")

		if value != nil  {
			c.Next()
			return
		}

		c.Abort()
		c.Redirect(http.StatusMovedPermanently,"/login")
		return
	}
}