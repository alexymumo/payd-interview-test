package model

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId       int       `json:"userid"`
	Email        string    `json:"email"`
	Name         string    `json:"name"`
	Password     string    `json:"password"`
	Token        string    `json:"token"`
	RefreshToken string    `json:"refresh"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

type UserResponse struct {
	Status  int                    `json:"status"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func VerifyPassword(hashedpassword string, userpassword string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedpassword), []byte(userpassword))
	msg := ""
	if err != nil {
		msg = err.Error()
	}
	check := true
	return check, msg
}
