package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type CategoryController struct {
}

func (CategoryController) Index(c *gin.Context) {
	c.HTML(http.StatusOK,"category/index.html",nil)
}
