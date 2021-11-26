package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Name       string     `json:"name" binding:"required" gorm:"type:varchar(50);not null"`
	Price      float64    `json:"price" binding:"required" gorm:"not null"`
	Stock      int16      `json:"stock" binding:"required" gorm:"not null"`
	CategoryID uint       `json:"categoryId" binding:"required"`
	Category   Category   `json:"category" gorm:"foreignkey:CategoryID"`
	Customers  []Customer `gorm:"many2many:customer_products;"`
}
