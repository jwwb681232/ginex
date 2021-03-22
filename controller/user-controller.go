package controller

import (
	"GinRest/dto"
	"GinRest/helper"
	"GinRest/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController interface {
	Update(ctx *gin.Context)
	Profile(ctx *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

func (controller userController) Update(ctx *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := ctx.ShouldBind(&userUpdateDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request",errDTO.Error(),helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest,response)
		return
	}

	user := helper.AuthUser(ctx.GetHeader("Authorization"))
	userUpdateDTO.ID = user.ID

	controller.userService.Update(userUpdateDTO)
	response := helper.BuildResponse(true,"success",userUpdateDTO)
	ctx.JSON(http.StatusOK,response)
	return
}

func (controller userController) Profile(ctx *gin.Context) {
	user := helper.AuthUser(ctx.GetHeader("Authorization"))
	response := helper.BuildResponse(true,"success",user)
	ctx.JSON(http.StatusOK,response)
	return
}

func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {
	return &userController{userService, jwtService}
}
