package main

import (
	"payd/internal/routes"
	"payd/pkg/db"
	"payd/pkg/utils"

	"github.com/gin-gonic/gin"
)

func init() {
	db.ConnectDb()
	utils.LoadEnvs()
}

func main() {
	router := gin.New()
	router.Use(gin.Logger())
	routes.UserRoutes(router)
	routes.PaymentRoutes(router)
	router.Run("localhost:8080")
}
