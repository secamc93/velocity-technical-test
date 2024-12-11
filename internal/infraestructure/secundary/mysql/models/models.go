package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name  string  `gorm:"type:varchar(255);not null"`
	Price float64 `gorm:"type:decimal(10,2);not null"`
	Stock int     `gorm:"not null"`
}

type Order struct {
	gorm.Model
	CustomerName string  `gorm:"type:varchar(255);not null"`
	TotalAmount  float64 `gorm:"type:decimal(10,2);not null"`
}

type OrderItem struct {
	gorm.Model
	OrderID   int     `gorm:"not null"`
	ProductID int     `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Subtotal  float64 `gorm:"type:decimal(10,2);not null"`
	Order     Order   `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}
