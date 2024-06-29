package repository

import (
	"net/http"
	"payd/internal/auth/model"
	"payd/pkg/db"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var userInput struct {
			Email    string `json:"email" binding:"required"`
			Name     string `json:"name" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		if err := c.ShouldBindJSON(&userInput); err != nil {
			c.JSON(http.StatusBadRequest, model.UserResponse{Status: http.StatusBadRequest, Message: "Bad request", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		var user model.User
		db.DB.Where("email=?", user.Email).Find(&user)

		if user.UserId != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "email already used"})
			return
		}

		password := model.HashPassword(userInput.Password)
		createUser := model.User{
			Email:    userInput.Email,
			Name:     userInput.Name,
			Password: string(password),
		}
		db.DB.Create(&createUser)
		c.JSON(http.StatusOK, gin.H{"user": createUser})
	}
}

func LoginUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var userInput struct {
			Email    string `json:"email" binding:"required"`
			Password string `json:"password" binding:"required"`
		}
		if err := ctx.ShouldBindJSON(&userInput); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error()})
			return
		}

		var user model.User
		db.DB.Where("email=?", userInput.Email).Find(&user)
		if user.UserId == 0 {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "email not found"})
			return
		}
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalide password"})
			return
		}
		generateToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":  user.UserId,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})
		token, err := generateToken.SignedString([]byte("adaffdla2323"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to generate token"})
		}
		ctx.JSON(200, gin.H{
			"token": token,
		})
	}

}
