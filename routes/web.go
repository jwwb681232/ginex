package routes

import (
	"github.com/gin-gonic/gin"
	"ginex/controllers/auth"
)

func Init() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("views/**/*")
	router.GET("/login",auth.ShowLoginForm)
	router.GET("/register",auth.ShowRegistrationForm)
	router.GET("/register-all",auth.ShowRegistrationFormAll)

	return router
}