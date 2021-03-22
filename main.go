package main

import (
	"GinRest/config"
	"GinRest/controller"
	"GinRest/middleware"
	"GinRest/repository"
	"GinRest/service"
	"github.com/gin-gonic/gin"
)

var (
	db                = config.SetupDatabaseConnection()
	userRepository    = repository.NewUserRepository(db)
	jwtService        = service.NewJWTService()
	authService       = service.NewAuthService(userRepository)
	authController    = controller.NewAuthController(authService, jwtService)
	userService       = service.NewUserService(userRepository)
	userController    = controller.NewUserController(userService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()
	/*authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
		authRoutes.POST("/register", authController.Register)
	}
	profileRoutes := r.Group("api/profile",middleware.AuthorizeJWT(jwtService))
	{
		profileRoutes.GET("/a", profileController.Index)
	}*/

	api := r.Group("api")
	{
		api.POST("/auth/login", authController.Login)
		api.POST("/auth/register", authController.Register)

		api.GET("/user/profile",middleware.AuthorizeJWT(jwtService),userController.Profile)
		api.PUT("/user",middleware.AuthorizeJWT(jwtService),userController.Update)
	}

	_ = r.Run()
}
