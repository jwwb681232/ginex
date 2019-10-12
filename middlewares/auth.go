package middlewares

import (
	"ginex/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/tools/go/ssa/interp/testdata/src/fmt"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {

		session := sessions.Default(c)
		var value models.User
		value = session.Get("ginex_session_key").(models.User)
		fmt.Printf(value)
		if value.ID != 0  {
			c.Next()
			return
		}

		c.Abort()
		c.Redirect(http.StatusMovedPermanently,"/login")
		return
	}
}