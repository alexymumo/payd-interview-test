package model

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserId    uint32    `gorm:"primary_key;auto_increment" json:"userid"`
	Email     string    `gorm:"size:255;not null;unique" json:"email"`
	Name      string    `gorm:"size:255;not null;unique" json:"name"`
	Password  string    `gorm:"not null" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
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
