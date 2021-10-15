package service

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/rest-api-market/connection"
	"github.com/rest-api-market/model"
)

type Methods interface {
	Create(c *gin.Context) (model.Category, error)
	Delete(c *gin.Context)
	Update(c *gin.Context)
	GetAll() []model.Category
	GetbyId(c *gin.Context) model.Category
}

type CategoryPatch struct {
	ID          int    `json:"ID" binding:"required"`
	Name        string `json:"name,omitempty" binding:"required"`
	Description string `json:"description,omitempty" binding:"required"`
}

type categoryService struct {
}

func New() Methods {
	return &categoryService{}
}

func (cs *categoryService) Create(c *gin.Context) (model.Category, error) {
	var category model.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		return category, err
	}
	connection.GetConnection().Create(&category)
	return category, nil
}

func (cs *categoryService) Delete(c *gin.Context) {

}

func (cs *categoryService) Update(c *gin.Context) {
	var categoryPatch CategoryPatch
	c.BindJSON(&categoryPatch)
	log.Print(categoryPatch)
	connection.GetConnection().Model(&model.Category{}).Select("*").Where("category_id = ?", categoryPatch.ID).Updates(categoryPatch)
	// log.Print(category)
}

func (cs *categoryService) GetbyId(c *gin.Context) model.Category {
	var category model.Category
	connection.GetConnection().Find(&category, c.Param("id"))
	return category
}

func (cs *categoryService) GetAll() []model.Category {
	var category []model.Category
	connection.GetConnection().Find(&category)

	//Preloads the relationship
	//db.Preload("Products").Find(&category)
	return category
}
