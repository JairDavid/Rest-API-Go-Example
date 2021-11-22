package model

import "github.com/jinzhu/gorm"

type CustomerProduct struct {
	gorm.Model
	CustomerID uint `json:"customerId" binding:"required"`
	ProductID  uint `json:"ProductId" binding:"required"`
}
