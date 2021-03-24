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
	bookRepository    = repository.NewBookRepository(db)
	bookService       = service.NewBookService(bookRepository)
	bookController    = controller.NewBookController(bookService, jwtService)
)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	api := r.Group("api")
	{
		api.POST("/auth/login", authController.Login)
		api.POST("/auth/register", authController.Register)

		api.GET("/user/profile",middleware.AuthorizeJWT(jwtService),userController.Profile)
		api.PUT("/user",middleware.AuthorizeJWT(jwtService),userController.Update)

		api.GET("/books",bookController.All)
		api.POST("/books",middleware.AuthorizeJWT(jwtService),bookController.Insert)
		api.GET("/books/:id",bookController.FindByID)
		api.PUT("/books/:id",middleware.AuthorizeJWT(jwtService),bookController.Update)
		api.DELETE("/books/:id",middleware.AuthorizeJWT(jwtService),bookController.Delete)
	}



	_ = r.Run()
}
