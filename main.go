package main

import (
	"GinRest/config"
	"GinRest/controller"
	"GinRest/repository"
	"GinRest/service"
	"github.com/gin-gonic/gin"
)

var (
	db             = config.SetupDatabaseConnection()
	userRepository = repository.NewUserRepository(db)
	jwtService     = service.NewJWTService()
	authService    = service.NewAuthService(userRepository)
	authController = controller.NewAuthController(authService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}

	_ = r.Run()
}
