package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {
}

func (IndexController) Index(c *gin.Context) {
	c.HTML(http.StatusOK,"layout/index.html",nil)
}
