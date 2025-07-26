package models

type Category struct {
	ID     uint      `gorm:"primaryKey" json:"id"`
	Nama   string    `json:"nama"`
	Produk []Product `json:"produk,omitempty" gorm:"foreignKey:KategoriID"`
}
