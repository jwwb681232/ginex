package helpers

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/sessions"
)

func Flash(context *gin.Context,value interface{}) []interface{} {
	session := sessions.Default(context)
	session.AddFlash(value)
	flash := session.Flashes()
	session.Save()
	return flash
}
