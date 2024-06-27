package repository

import (
	"github.com/gin-gonic/gin"
)

func RegisterUser() gin.HandlerFunc {
	return func(c *gin.Context) {

		/*
			var ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
			var user model.User
			if err := c.BindJSON(&user); err != nil {
				c.JSON(http.StatusBadRequest, model.UserResponse{Status: http.StatusBadRequest, Message: "Bad Request", Data: map[string]interface{}{"data": err.Error()}})
			}
			password := model.HashPassword(*&user.Password)
			user.Password = password
			user.CreatedAt, _ = time.Parse(time.RFC1123, time.Now().Format(time.RFC1123))
			user.UpdatedAt, _ = time.Parse(time.RFC1123, time.Now().Format(time.RFC1123))
			token, refreshToken, _ := utils.CreateToken(*&user.Name, *&user.Email, user.UserId)
			user.Token = token
			user.RefreshToken = refreshToken
			stmt,err := db.Prepare("insert into users(name,password,email,token) values ($1,$2,$3,$4)")
			if err !=nil {
				log.Fatal(err)
			}
			defer stmt.Close()
			if _,err := stmt.Exec() {
				log.Fatal(err)
			}
			c.JSON(http.StatusCreated, user)
		*/
	}
}

func LoginUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
