package routes

import (
	"ginex/controllers/auth"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession",store))

	router.LoadHTMLGlob("views/**/*")

	router.GET("/register",auth.ShowRegistrationForm)
	router.POST("/register",auth.Register)

	router.GET("/login",auth.ShowLoginForm)
	router.POST("/login",auth.Login)


	return router
}