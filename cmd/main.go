package main

import (
	"payd/internal/routes"
	"payd/pkg/db"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()
	db.ConnectDb()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	routes.PaymentRoutes(router)
	//router.Use()
	router.Run("localhost:8080")
}
