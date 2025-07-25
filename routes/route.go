package routes

import (
	"kasir/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoute(route *gin.Engine) {
	route.GET("/produk", controllers.GetProducts)
	route.GET("/produk/:id", controllers.GetProduct)
	route.POST("/produk", controllers.CreateProduct)
	route.PUT("/produk/:id", controllers.UpdateProduct)
	route.DELETE("/produk/:id", controllers.DeleteProduct)
}
