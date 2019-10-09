package middlewares

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		sessionKey,_ := c.Request.Cookie("ginex_session")
		fmt.Println(sessionKey)
		session := sessions.Default(c)
		value := session.Get(sessionKey.Value)
		fmt.Println(value)

		/*sessionToken, _ := c.Cookie("ginex_session")
		session := sessions.Default(c)
		value := session.Get(sessionToken)
		fmt.Println(sessionToken)
		fmt.Println(value)*/
		c.Next()
	}
}
