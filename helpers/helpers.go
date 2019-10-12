package helpers

import (
	"ginex/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	zhCn "github.com/go-playground/locales/zh"
	"github.com/go-playground/universal-translator"
	"gopkg.in/go-playground/validator.v9"
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

func SetUserSession(c *gin.Context,user models.User) {
	//sessionKey := "ginex_session_key"
	//c.SetCookie("ginex_session",sessionKey,60*60*24,"/","localhost",true,true)
	//session := sessions.Default(c)
	//session.Set(sessionKey,ID)
	//_ = session.Save()

	session := sessions.Default(c)
	session.Options(sessions.Options{
		MaxAge:   60*60*2,
	})
	//jsonData,_ := json.Marshal(user)
	session.Set("ginex_session_key",user)
	_ = session.Save()
}