package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectionDB() {
	dsn := "golang:golang@tcp(127.0.0.1:3306)/kasir_db?parseTime=true"
	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("gagal koneksi ke database")
	}

	DB = database
	fmt.Println("database terkoneksi")
}
