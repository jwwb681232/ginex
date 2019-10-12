package middlewares

import (
	"ginex/models/user"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		var value user.User
		value = session.Get("ginex_session_key").(user.User)

		if value.ID != 0  {
			c.Next()
			return
		}

		c.Abort()
		c.Redirect(http.StatusMovedPermanently,"/login")
		return
	}
}