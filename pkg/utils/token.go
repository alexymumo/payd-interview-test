package utils

import (
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

var SECRET_KEY = os.Getenv("")

type UserCredential struct {
	UserId uint32
	jwt.StandardClaims
}

func CreateToken(userId uint32) (token string, refreshToken string, err error) {
	claims := &UserCredential{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	refreshClaims := &UserCredential{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(200)).Unix(),
		},
	}
	jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString([]byte(SECRET_KEY))
	if err != nil {
		log.Panic(err)
		return
	}
	jwtRefreshToken, err := jwt.NewWithClaims(jwt.SigningMethodES256, refreshClaims).SignedString([]byte(SECRET_KEY))
	return jwtToken, jwtRefreshToken, err
}

func ParseToken(tokenString string) (*UserCredential, error) {
	claims := &UserCredential{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return t, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, err
	}
	return claims, nil
}

func VerifyToken(signedToken string) (claims *UserCredential, msg string) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&UserCredential{},
		func(t *jwt.Token) (interface{}, error) {
			return []byte(SECRET_KEY), nil
		},
	)
	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*UserCredential)
	if !ok {
		msg = "token is invalid"
		msg = err.Error()
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token is invalid"
		msg = err.Error()
		return
	}
	return claims, msg
}

func GetUserId(ctx *gin.Context) uint32 {
	userID, exists := ctx.Get("userid")
	if !exists {
		return 0
	}
	return userID.(uint32)
}
