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
	CustomerName string      `gorm:"type:varchar(255);not null"`
	TotalAmount  float64     `gorm:"type:decimal(10,2);not null"`
	OrderItems   []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	gorm.Model
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Quantity  uint    `gorm:"not null"`
	Subtotal  float64 `gorm:"type:decimal(10,2);not null"`
	Order     Order   `gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	Product   Product `gorm:"foreignKey:ProductID;constraint:OnDelete:CASCADE"`
}
