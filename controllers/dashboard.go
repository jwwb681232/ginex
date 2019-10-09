package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type DashboardController struct {
}

func (DashboardController) Index(c *gin.Context) {
	c.HTML(http.StatusOK,"dashboard/index.html",nil)
}
