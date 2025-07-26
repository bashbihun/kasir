package main

import (
	"kasir/config"
	"kasir/models"
	"kasir/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	config.ConnectionDB()
	config.DB.AutoMigrate(
		&models.Product{},
		&models.Category{},
	)
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"message": "kasir api berjalan"})
	})
	routes.SetupRoute(router)

	router.Run()
}
