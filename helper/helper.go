package helper

import (
	"GinRest/config"
	"GinRest/entity"
	"GinRest/repository"
	"GinRest/service"
	"github.com/dgrijalva/jwt-go"
)

func AuthUser(text string) *entity.User {
	token, _ := service.NewJWTService().ValidateToken(text)
	claims := token.Claims.(jwt.MapClaims)
	userRepository := repository.NewUserRepository(config.SetupDatabaseConnection())
	user := userRepository.ProfileUser(claims["user_id"].(string))
	return &user
}