package controller

import (
	"GinRest/helper"
	"GinRest/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ProfileController interface {
	Index(ctx *gin.Context)
}

type profileController struct {
	profileService service.ProfileService
	jwtService  service.JWTService
}

func NewProfileController(profileService service.ProfileService, jwtService service.JWTService) ProfileController {
	return &profileController{
		profileService: profileService,
		jwtService:  jwtService,
	}
}

func (c *profileController) Index(ctx *gin.Context) {
	user := helper.AuthUser(ctx.GetHeader("Authorization"))
	response := helper.BuildResponse(true,"Ok!",user)
	ctx.JSON(http.StatusOK, response)
	return
}