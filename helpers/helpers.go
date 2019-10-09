package helpers

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
	zhCn "github.com/go-playground/locales/zh"
	zhTranslations "gopkg.in/go-playground/validator.v9/translations/zh"
)

func Flash(context *gin.Context,value interface{}) interface{} {
	session := sessions.Default(context)
	session.AddFlash(value)
	flash := session.Flashes()[0]
	_ = session.Save()
	return flash
}

func Validate(s interface{}) (map[string]string,error)  {
	var uni *ut.UniversalTranslator
	var validate *validator.Validate

	zh := zhCn.New()
	uni = ut.New(zh,zh)

	trans,_ := uni.GetTranslator("zh")

	validate = validator.New()
	_ = zhTranslations.RegisterDefaultTranslations(validate, trans)

	err := validate.Struct(s)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		return errs.Translate(trans),err
	}
	return make(map[string]string),err
}