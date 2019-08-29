package middlewares

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("middleware auth")
		c.Next()
	}
}
