package controller

import (
	"GinRest/dto"
	"GinRest/entity"
	"GinRest/helper"
	"GinRest/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type BookController interface {
	All(ctx *gin.Context)
	FindByID(ctx *gin.Context)
	Insert(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type bookController struct {
	bookService service.BookService
	jwtService  service.JWTService
}

func (controller *bookController) All(ctx *gin.Context) {
	var books = controller.bookService.All()
	response := helper.BuildResponse(true, "ok!", books)
	ctx.JSON(http.StatusOK, response)
	return
}

func (controller *bookController) FindByID(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("No param id was found", err.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	var book = controller.bookService.FindByID(id)
	if (book == entity.Book{}) {
		response := helper.BuildErrorResponse("Data not found.", "No data with given id", helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusNotFound, response)
		return
	}

	response := helper.BuildResponse(true, "ok!", book)
	ctx.JSON(http.StatusOK, response)
	return
}

func (controller *bookController) Insert(ctx *gin.Context) {
	var bookCreateDTO dto.BookCreateDTO
	errDTO := ctx.ShouldBind(&bookCreateDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authUser := helper.AuthUser(ctx.GetHeader("Authorization"))
	bookCreateDTO.UserID = authUser.ID

	result := controller.bookService.Insert(bookCreateDTO)
	response := helper.BuildResponse(true, "ok!", result)
	ctx.JSON(http.StatusOK, response)
	return
}

func (controller *bookController) Update(ctx *gin.Context) {
	var bookUpdateDTO dto.BookUpdateDTO
	errDTO := ctx.ShouldBind(&bookUpdateDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authUser := helper.AuthUser(ctx.GetHeader("Authorization"))

	if !controller.bookService.IsAllowedToEdit(fmt.Sprintf("%v", authUser.ID), bookUpdateDTO.ID) {
		response := helper.BuildErrorResponse("You dont have permission", "You are not owner", helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusForbidden, response)
		return
	}

	bookUpdateDTO.UserID = authUser.ID

	result := controller.bookService.Update(bookUpdateDTO)
	response := helper.BuildResponse(true, "ok!", result)
	ctx.JSON(http.StatusOK, response)
	return
}

func (controller *bookController) Delete(ctx *gin.Context) {
	var book entity.Book
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Failed to get id", "No param id were found", helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return
	}

	authUser := helper.AuthUser(ctx.GetHeader("Authorization"))

	book.ID = id
	if !controller.bookService.IsAllowedToEdit(fmt.Sprintf("%v", authUser.ID), book.ID) {
		response := helper.BuildErrorResponse("You dont have permission", "You are not owner", helper.EmptyObject{})
		ctx.AbortWithStatusJSON(http.StatusForbidden, response)
		return
	}

	controller.bookService.Delete(book)
	response := helper.BuildResponse(true, "ok!", helper.EmptyObject{})
	ctx.JSON(http.StatusOK, response)
	return
}

func NewBookController(bookService service.BookService, jwtService service.JWTService) BookController {
	return &bookController{bookService, jwtService}
}
