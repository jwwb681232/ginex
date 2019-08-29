package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"ginex/controllers/auth"
	"ginex/middlewares"
)

func Init() *gin.Engine {
	router := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession",store))
	router.Use(middlewares.Auth())

	router.LoadHTMLGlob("views/**/*")

	router.GET("/register",auth.RegisterController{}.ShowRegistrationForm)
	router.POST("/register",auth.RegisterController{}.Register)

	router.GET("/login",auth.LoginController{}.ShowLoginForm)
	router.POST("/login",auth.LoginController{}.Login)

	router.GET("/set-cookie",auth.LoginController{}.Set)
	router.GET("/get-cookie",auth.LoginController{}.Get)

	return router
}