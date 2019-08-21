package auth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"ginex/helpers"
)

type RegisterController struct {}

type registerForm struct {
	Name                 string `form:"name" json:"name" validate:"required"`
	Email                string `form:"email" json:"email" validate:"required,email"`
	Password             string `form:"password" json:"password" validate:"required"`
	PasswordConfirmation string `form:"password_confirmation" json:"password_confirmation" validate:"required"`
}

func (RegisterController) ShowRegistrationForm(c *gin.Context) {
	c.HTML(http.StatusOK, "auth/register.html", gin.H{"code":http.StatusOK,"message":make(map[string]string)})
}

func (RegisterController) Register(c *gin.Context) {
	var data registerForm
	if err := c.Bind(&data); err != nil {
		/*c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return*/
	}

	message,err := helpers.Validate(&data)
	if err != nil {
		c.JSON(http.StatusOK,message)
	}

	/*var uni *ut.UniversalTranslator
	var validate *validator.Validate

	zh := zhCn.New()
	uni = ut.New(zh,zh)

	trans,_ := uni.GetTranslator("zh")

	validate = validator.New()
	zhTranslations.RegisterDefaultTranslations(validate,trans)

	err := validate.Struct(&data)
	if err != nil{
		errs := err.(validator.ValidationErrors)

		c.JSON(http.StatusOK,errs.Translate(trans))

		//fmt.Println(errs.Translate(trans))
	}*/

}