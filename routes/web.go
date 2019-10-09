package routes

import (
	"ginex/controllers"
	"ginex/controllers/auth"
	"ginex/middlewares"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/**/*")

	router.Use(sessions.Sessions("ginex_session",cookie.NewStore([]byte("secret"))))

	router.GET("/register",auth.RegisterController{}.ShowRegistrationForm)
	router.POST("/register",auth.RegisterController{}.Register)

	router.GET("/login",auth.LoginController{}.ShowLoginForm)
	router.POST("/login",auth.LoginController{}.Login)

	router.Use(middlewares.Auth())
	{
		router.GET("/dashboard",controllers.DashboardController{}.Index)
	}

	//router.GET("/set-cookie",auth.LoginController{}.Set)
	//router.GET("/get-cookie",auth.LoginController{}.Get)

	return router
}