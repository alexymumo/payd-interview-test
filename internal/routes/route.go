package routes

import (
	"payd/internal/auth/repository"
	"payd/internal/auth/service"
	repo "payd/internal/payments/repository"

	"github.com/gin-gonic/gin"
)

func UserRoutes(route *gin.Engine) {
	route.POST("v1/auth/register", repository.RegisterUser())
	route.POST("v1/auth/login", repository.LoginUser())
}

func PaymentRoutes(route *gin.Engine) {
	paymentRoutes := route.Group("/payments")
	paymentRoutes.Use(service.AuthMiddleware())
	{
		paymentRoutes.POST("/initiate", repo.MakePayment())
		paymentRoutes.GET("/status/:id", repo.CheckPaymentStatus())
	}

}
