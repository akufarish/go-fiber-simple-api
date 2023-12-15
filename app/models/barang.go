package models

type Barang struct {
	Id        int64  `gorm:"primaryKey" json:"id" params:"id"`
	Nama      string `gorm:"type:varchar(255)" json:"nama" form:"nama"`
	Deskripsi string `gorm:"type:varchar(255)" json:"deskripsi" form:"deskripsi"`
	Harga     int    `json:"harga" form:"harga"`
}