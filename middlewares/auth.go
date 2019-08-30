package middlewares

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionToken, _ := c.Cookie("ginex_session")
		session := sessions.Default(c)
		value := session.Get(sessionToken)
		fmt.Println(value)
	}
}
