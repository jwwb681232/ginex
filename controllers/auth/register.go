package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	userModel "ginex/models/user"
	"ginex/helpers"
	"golang.org/x/crypto/bcrypt"
)

type RegisterController struct {}

type registerForm struct {
	Name                 string `form:"name" json:"name" validate:"required"`
	Email                string `form:"email" json:"email" validate:"required,email"`
	Password             string `form:"password" json:"password" validate:"required,min=6,max=20"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" validate:"required,eqfield=Password"`
}

func (RegisterController) ShowRegistrationForm(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/register.html", gin.H{"code":http.StatusOK,"message":make(map[string]string)})
}

func (RegisterController) Register(c *gin.Context) {
	var data registerForm
	_ = c.Bind(&data)
	validateMessage, err := helpers.Validate(&data)
	if err != nil {
		c.HTML(http.StatusBadRequest, "auth/register.html", gin.H{"code": http.StatusFound, "message": validateMessage, "data": data})
		return
	}

	userData, notFound := userModel.User{}.WhereEmail(&data.Email)
	if !notFound {
		c.HTML(http.StatusBadRequest, "auth/register.html", gin.H{"code": http.StatusFound, "message": map[string]string{"registerForm.Email": "邮箱已经存在"}, "data": data,}, )
		return
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(data.Password), bcrypt.DefaultCost)
	createUser := userModel.User{
		Name:data.Name,
		Email:data.Email,
		Password:string(hash),
	}

	result := userModel.User{}.CreateUser(createUser)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest,result.Error)
		return
	}

	c.JSON(http.StatusOK, userData)

	return
}