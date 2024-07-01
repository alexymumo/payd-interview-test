package routes

import (
	"payd/internal/auth/repository"
	repo "payd/internal/payments/repository"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	route.POST("v1/auth/register", repository.RegisterUser())
	route.POST("v1/auth/login", repository.LoginUser())
}

func PaymentRoutes(route *gin.Engine) {
	route.POST("v1/payments", repo.MakePayment())
	route.GET("v1/payments/status/:id", repo.CheckPaymentStatus())
}
