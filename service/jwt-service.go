package service

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

type JWTService interface {
	GenerateToken(userID string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "ydhnwb",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "ginrest"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(UserID string) string {
	claims := &jwtCustomClaim{
		UserID: UserID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}

	jsonWebToken := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	token,err := jsonWebToken.SignedString([]byte(j.secretKey))
	if err != nil {
		panic(err)
	}

	return token
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token,error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC);!ok {
			return nil,fmt.Errorf("unexpected signing method %v",t.Header["alg"])
		}

		return []byte(j.secretKey),nil
	})
}