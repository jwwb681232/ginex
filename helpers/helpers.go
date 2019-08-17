package helpers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Flash(context *gin.Context,value interface{}) interface{} {
	session := sessions.Default(context)
	session.AddFlash(value)
	flash := session.Flashes()[0]
	session.Save()
	return flash
}
