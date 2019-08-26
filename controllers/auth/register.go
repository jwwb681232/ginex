package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	userModel "ginex/models/user"
)

type RegisterController struct {}

type registerForm struct {
	Name                 string `form:"name" json:"name" validate:"required"`
	Email                string `form:"email" json:"email" validate:"required,email"`
	Password             string `form:"password" json:"password" validate:"required,min=6,max=20"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" validate:"required,eqfield=password"`
}

func (RegisterController) ShowRegistrationForm(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/register.html", gin.H{"code":http.StatusOK,"message":make(map[string]string)})
}

func (RegisterController) Register(c *gin.Context) {
	var data registerForm
	c.Bind(&data)
	/*validateMessage,err := helpers.Validate(&data)
	if err != nil {
		c.HTML(http.StatusOK,"auth/register.html",gin.H{"code":http.StatusFound,"message":validateMessage,"data":data})
	}*/

	userData,err := userModel.User{}.WhereEmail(&data.Email)
	if err != nil {
		c.JSON(http.StatusOK,gin.H{"error":err})
		return
	}

	c.JSON(http.StatusOK,gin.H{"data":userData})

	return
}