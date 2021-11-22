package service

import (
	"github.com/gin-gonic/gin"
	"github.com/rest-api-market/connection"
	"github.com/rest-api-market/model"
)

type CategoryService struct {
}

func NewCategoryRepository() Repository {
	return &CategoryService{}
}

func (cs *CategoryService) Create(c *gin.Context) (interface{}, error) {
	var category model.Category
	err := c.ShouldBindJSON(&category)
	if err != nil {
		return category, err
	}
	connection.GetConnection().Create(&category)
	return category, nil
}

func (cs *CategoryService) Delete(c *gin.Context) (interface{}, error) {
	var category model.Category
	connection.GetConnection().Find(&category, c.Param("id"))
	err := c.ShouldBindJSON(&category)
	if err != nil {
		return category, err
	}
	connection.GetConnection().Unscoped().Delete(&category, category.ID)
	return category, err
}

func (cs *CategoryService) Update(c *gin.Context) (interface{}, interface{}) {
	//implementing patch method, it could be better.
	var category model.Category

	//recieve any data from the request body
	var dinamicAtributes map[string]interface{}
	c.Bind(&dinamicAtributes)

	if dinamicAtributes == nil {
		//case 1: there's nothing in the request body
		return category, 1
	}
	connection.GetConnection().Find(&category, c.Param("id"))
	if category.ID == 0 {
		//case 2: inexistent ID
		return category, 2
	}

	statusUpdate := connection.GetConnection().Model(&model.Category{}).Where(c.Param("id")).Updates(dinamicAtributes)

	if statusUpdate.RowsAffected == 0 {
		//case 3: there's no matching data from the request body with the database attributes
		return category, 3
	} else {
		connection.GetConnection().Find(&category, c.Param("id"))
	}
	//case 4: Correct operation
	return category, 4
}

func (cs *CategoryService) GetById(c *gin.Context) interface{} {
	var category model.Category
	connection.GetConnection().Find(&category, c.Param("id"))
	return category
}

func (cs *CategoryService) GetAll() interface{} {
	var category []model.Category
	//connection.GetConnection().Model(&model.Category{}).Find(&category)
	//Preloads the relationship
	connection.GetConnection().Preload("Products").Find(&category)
	return category
}
