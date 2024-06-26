package main

import (
	"payd/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	routes.PaymentRoutes(router)
	//router.Use()
	router.Run("localhost:8080")
}
