package models

type Product struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Nama       string `json:"nama"`
	Harga      int    `json:"harga"`
	Stok       int    `json:"stok"`
	KategoriID uint   `json:"kategori_id"`
}
