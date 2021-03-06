package model

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name        string `json:"name" binding:"required" gorm:"type:varchar(20);not null"`
	Description string `json:"description" binding:"required" gorm:"type:varchar(50);not null"`

	//add foreing key at the migration, reference(category_id)
	Products []Product
}
