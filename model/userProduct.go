package model

import "github.com/jinzhu/gorm"

type CustomerProduct struct {
	gorm.Model
	CustomerID int `json:"customerId" binding:"required"`
	ProductID  int `json:"ProductId" binding:"required"`
}
