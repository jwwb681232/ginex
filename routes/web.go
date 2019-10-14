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

	authorized := router.Group("/")
	authorized.Use(middlewares.Auth())
	{
		authorized.GET("/index",controllers.IndexController{}.Index)
		authorized.GET("/dashboard",controllers.DashboardController{}.Index)
		authorized.GET("/home-cook/categories",controllers.CategoryController{}.Index)
	}


	return router
}