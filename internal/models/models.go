package models

type Profile struct {
	Name string
}

type Product struct {
	Name   string `json:"name" gorm:"not null"`
	ID     int    `json:"id" gorm:"not null; primary_key"`
	Price  int    `json:"price" gorm:"not null"`
	ShopID uint   `json:"shop_id" gorm:"not null"`
}

type Shop struct {
	ID       int       `json:"id" gorm:"not null; primary_key"`
	Name     string    `json:"name" gorm:"not null"`
	Products []Product `json:"products" gorm:"foreignKey:ShopID; not null"`
}
