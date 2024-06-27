package repository

import (
	"net/http"
	"os"
	"payd/internal/auth/model"
	"payd/pkg/db"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userInput struct {
			Email    string `json:"email"`
			Name     string `json:"name"`
			Password string `json:"password"`
		}
		if err := c.BindJSON(&userInput); err != nil {
			c.JSON(http.StatusBadRequest, model.UserResponse{Status: http.StatusBadRequest, Message: "Bad request", Data: map[string]interface{}{"data": err.Error()}})
		}
		password := model.HashPassword(userInput.Password)
		userInput.Password = password

		user := model.User{
			Email:    userInput.Email,
			Name:     userInput.Name,
			Password: userInput.Password,
		}
		result := db.DB.Create(&user)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to register user",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"user": result})
	}
}

func LoginUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userInput struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := ctx.BindJSON(&userInput); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Login failed",
			})
			return
		}

		var user model.User
		db.DB.First(&user, "email = ?", userInput.Email)
		if user.UserId == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid email or password",
			})
			return
		}
		if user.UserId == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid user"})
			return
		}
		generateToken := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
			"id":  user.UserId,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})
		token, err := generateToken.SignedString([]byte(os.Getenv("SECRET_KEY")))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to generate token"})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}

}
