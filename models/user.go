package models

import "time"

type User struct{
	UserId int
	Email string
	Password string
	CreatedAt time.Time
	UpdatedAt time.Time
}