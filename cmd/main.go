package main

import (
	"payd/internal/routes"
	"payd/pkg/db"

	"github.com/gin-gonic/gin"
)

func init() {
	db.ConnectDb()
	//utils.LoadEnvs()
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	routes.PaymentRoutes(router)
	//router.Use()
	router.Run("localhost:8080")
}
