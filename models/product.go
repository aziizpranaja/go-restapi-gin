package models

type Product struct {
	Id          int32  `gorm:"primaryKey" json:"id"`
	ProductName string `gorm:"type:varchar(150)" binding:"required" json:"product_name"`
	Deskripsi   string `gorm:"type:text" binding:"required" json:"deskripsi"`
}