package middlewares

import (
	"fmt"
	"ginex/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		var value models.User
		value = session.Get("ginex_session_key").(models.User)
		fmt.Println(value)
		if value.ID != 0  {
			c.Next()
			return
		}

		c.Abort()
		c.Redirect(http.StatusMovedPermanently,"/login")
		return
	}
}