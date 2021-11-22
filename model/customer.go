package model

import "github.com/jinzhu/gorm"

type Customer struct {
	gorm.Model
	Name     string `json:"name" binding:"required" gorm:"type:varchar(50);not null"`
	Lastname string `json:"lastname" binding:"required" gorm:"type:varchar(50);not null"`
	Order    []CustomerProduct
}
