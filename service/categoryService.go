package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rest-api-market/connection"
	"github.com/rest-api-market/model"
)

type CategoryRepository interface {
	Create(c *gin.Context) (model.Category, error)
	Delete(c *gin.Context) (model.Category, error)
	Update(c *gin.Context) (model.Category, interface{})
	GetAll() []model.Category
	GetbyId(c *gin.Context) model.Category
}

type categoryService struct {
}

func NewCategoryRepository() CategoryRepository {
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

func (cs *categoryService) Delete(c *gin.Context) (model.Category, error) {
	var category model.Category
	connection.GetConnection().Find(&category, c.Param("id"))
	err := c.ShouldBindJSON(&category)
	if err != nil {
		return category, err
	}
	connection.GetConnection().Delete(&category, category.ID)
	return category, err
}

func (cs *categoryService) Update(c *gin.Context) (model.Category, interface{}) {
	//implementing patch method, it could be better.
	var category model.Category

	//recieve any data from the request body
	var dinamicAtributes map[string]interface{}
	c.Bind(&dinamicAtributes)

	if dinamicAtributes == nil {
		//case 1: there's nothing in the request body
		return category, 1
	}
	statusFind := connection.GetConnection().Find(&category, c.Param("id"))
	if statusFind != nil {
		//case 2: inexistent ID
		return category, 2
	}

	statusUpdate := connection.GetConnection().Model(&model.Category{}).Where(c.Param("id")).Updates(dinamicAtributes)
	if statusUpdate != nil {
		//case 3: there's no matching data from the request body with the database attributes
		return category, 3
	} else {
		connection.GetConnection().Find(&category, c.Param("id"))
	}
	//case 4: Correct operation
	return category, 4
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
