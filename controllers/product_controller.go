package controllers

import (
	"kasir/config"
	"kasir/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateProduct(ctx *gin.Context) {
	var product models.Product
	if err := ctx.ShouldBindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "input tidak valid"})
		return
	}

	if err := config.DB.Create(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menyimpan produk"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": product})
}

func GetProducts(ctx *gin.Context) {
	var products []models.Product
	config.DB.Find(&products)
	ctx.JSON(http.StatusOK, gin.H{"data": products})
}

func GetProduct(ctx *gin.Context) {
	var product models.Product

	if err := config.DB.First(&product, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "produk tidak ditemukan"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": product})
}

func UpdateProduct(ctx *gin.Context) {
	var product models.Product

	if err := config.DB.First(&product, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "produk tidak ditemukan"})
		return
	}

	var input models.Product
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "input tidak valid"})
		return
	}

	product.Nama = input.Nama
	product.Harga = input.Harga
	product.Stok = input.Stok

	config.DB.Save(&product)
	ctx.JSON(http.StatusOK, gin.H{"data": product})
}

func DeleteProduct(ctx *gin.Context) {
	var product models.Product

	if err := config.DB.First(&product, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "produk tidak ditemukan"})
		return
	}

	config.DB.Delete(&product)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
