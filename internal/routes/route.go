package routes

import (
	"payd/internal/auth/repository"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	route.POST("v1/auth/register", repository.RegisterUser())
	route.POST("v1/auth/login", repository.LoginUser())
}

func PaymentRoutes(route *gin.Engine) {
	route.POST("v1/payments")
	route.GET("v1/payments/status/:id")
}
