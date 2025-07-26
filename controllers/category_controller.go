package controllers

import (
	"kasir/config"
	"kasir/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCategories(ctx *gin.Context) {
	var kategori []models.Category
	config.DB.Preload("produk").Find(&kategori)
	ctx.JSON(http.StatusOK, gin.H{"data": kategori})
}

func CreateCategory(ctx *gin.Context) {
	var input models.Category
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "input tidak valid"})
		return
	}

	if err := config.DB.Create(&input).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "gagal menyimpan"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": input})
}

func GetCategory(ctx *gin.Context) {
	var kategori models.Category
	if err := config.DB.Preload("Produk").First(&kategori, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "kategori tidak ditemukan"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"data": kategori})
}

func UpdateCategory(ctx *gin.Context) {
	var kategori models.Category
	if err := config.DB.First(&kategori, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "kategori tidak ditemukan"})
		return
	}

	var input models.Category
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "input tidak valid"})
		return
	}

	kategori.Nama = input.Nama
	config.DB.Save(&kategori)

	ctx.JSON(http.StatusOK, gin.H{"data": kategori})
}

func DeleteCategory(ctx *gin.Context) {
	var kategori models.Category
	if err := config.DB.First(&kategori, ctx.Param("id")).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "kategori tidak ditemukan"})
		return
	}

	config.DB.Delete(&kategori)
	ctx.JSON(http.StatusOK, gin.H{"data": true})
}
