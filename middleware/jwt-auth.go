package middleware

import (
	"GinRest/helper"
	"GinRest/service"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func AuthorizeJWT(service service.JWTService) gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			context.AbortWithStatusJSON(http.StatusBadRequest, response)
			return
		}

		token, err := service.ValidateToken(authHeader)
		if err != nil {
			response := helper.BuildErrorResponse("Token is not valid",err.Error(),nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized,response)
			return
		}

		if !token.Valid {
			response := helper.BuildErrorResponse("Token is not valid","Token is not valid",nil)
			context.AbortWithStatusJSON(http.StatusUnauthorized,response)
			return
		}else{
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[user_id]", claims["user_id"])
			log.Println("Claim[issuer]", claims["issuer"])
		}
	}
}
