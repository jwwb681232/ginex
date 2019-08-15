package routes

import (
	"github.com/gin-gonic/gin"
	"ginex/controllers/auth"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/**/*")

	router.GET("/register",auth.ShowRegistrationForm)
	router.POST("/register",auth.Register)

	router.GET("/login",auth.ShowLoginForm)
	router.POST("/login",auth.Login)


	return router
}