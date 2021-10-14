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

type categoryService struct {
}

func New() Methods {
	return &categoryService{}
}

func (cs *categoryService) Create(c *gin.Context) (model.Category, error) {
	var category model.Category

	db := connection.GetConnection()
	err := c.ShouldBindJSON(&category)

	if err != nil {
		return category, err
	}

	db.Create(&category)
	defer db.Close()

	return category, nil
}

func (cs *categoryService) Delete(c *gin.Context) {

}

func (cs *categoryService) Update(c *gin.Context) {
	var category model.Category
	c.BindJSON(&category)
	log.Print(category)
	db := connection.GetConnection()
	db.Find(&category, category.ID)
	log.Print(category)

}

func (cs *categoryService) GetbyId(c *gin.Context) model.Category {
	var category model.Category
	db := connection.GetConnection()
	db.Find(&category, c.Param("id"))
	db.Close()
	return category
}

func (cs *categoryService) GetAll() []model.Category {
	var category []model.Category
	db := connection.GetConnection()
	db.Find(&category)
	db.Close()
	return category
}
